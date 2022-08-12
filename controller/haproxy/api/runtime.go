package api

import (
	"context"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/maps"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/server"
	"github.com/haproxytech/kubernetes-ingress/controller/store"
	"github.com/haproxytech/kubernetes-ingress/controller/utils"
	"github.com/haproxytech/models"
)

func (c *haProxyClient) ExecuteRaw(command string) (result []string, err error) {
	return []string{}, nil
}

func (c *haProxyClient) SetServerAddr(backendName string, serverName string, ip string, port int) error {
	runtimeServer := models.RuntimeServer{Name: serverName, Address: ip, Port: utils.PtrInt64(int64(port))}
	runtimeWriter := server.NewUpdateRuntimeWriter()
	runtimeWriter.WithContext(context.Background()).WithName(serverName).WithBackend(backendName).
		WithRuntimeServer(runtimeServer)
	_, err := c.client.Server.UpdateRuntimeServer(runtimeWriter)
	return err
}

func (c *haProxyClient) SetServerState(backendName string, serverName string, state string) error {
	runtimeServer := models.RuntimeServer{Name: serverName, AdminState: state}
	runtimeWriter := server.NewUpdateRuntimeWriter()
	runtimeWriter.WithContext(context.Background()).WithName(serverName).WithBackend(backendName).
		WithRuntimeServer(runtimeServer)
	_, err := c.client.Server.UpdateRuntimeServer(runtimeWriter)
	return err
}

func (c *haProxyClient) SetMapContent(mapFile string, key string, value string) error {
	mapsWriter := maps.NewCreateMapFileWriter()
	mapsWriter.WithContext(context.Background()).WithFileName(mapFile).WithForceSync(true).
		WithData(models.MapEntry{Key: key, Value: value})
	_, err := c.client.Maps.CreateMapFile(mapsWriter)
	return err
}

func (c *haProxyClient) GetMap(mapFile string) (*models.Map, error) {
	mapsWriter := maps.NewGetMapFileWriter()
	mapsWriter.WithContext(context.Background()).WithName(mapFile)
	maps, err := c.client.Maps.GetMapFile(mapsWriter)
	if err != nil {
		return nil, err
	}
	return maps.Payload, nil
}

func (c *haProxyClient) DeleteMap(mapFile string) error {
	mapsWriter := maps.NewDeleteMapFileWriter()
	mapsWriter.WithName(mapFile).WithContext(context.Background()).WithForceDelete(false).WithForeSync(true)
	_, err := c.client.Maps.DeleteMapFile(mapsWriter)
	return err
}

// SyncBackendSrvs syncs states and addresses of a backend servers with corresponding endpoints.
func (c *haProxyClient) SyncBackendSrvs(backend *store.RuntimeBackend, portUpdated bool) error {
	if backend.Name == "" {
		return nil
	}
	haproxySrvs := backend.HAProxySrvs
	addresses := backend.Endpoints.Addresses
	// Disable stale entries from HAProxySrvs
	// and provide list of Disabled Srvs
	var disabled []*store.HAProxySrv
	var errors utils.Errors
	for i, srv := range haproxySrvs {
		srv.Modified = srv.Modified || portUpdated
		if _, ok := addresses[srv.Address]; ok {
			delete(addresses, srv.Address)
		} else {
			haproxySrvs[i].Address = ""
			haproxySrvs[i].Modified = true
			disabled = append(disabled, srv)
		}
	}

	// Configure new Addresses in available HAProxySrvs
	for newAddr := range addresses {
		if len(disabled) == 0 {
			break
		}
		disabled[0].Address = newAddr
		disabled[0].Modified = true
		disabled = disabled[1:]
		delete(addresses, newAddr)
	}
	// Dynamically updates HAProxy backend servers  with HAProxySrvs content
	var addrErr, stateErr error
	for _, srv := range haproxySrvs {
		if !srv.Modified {
			continue
		}
		if srv.Address == "" {
			// logger.Tracef("server '%s/%s' changed status to %v", newEndpoints.BackendName, srv.Name, "maint")
			addrErr = c.SetServerAddr(backend.Name, srv.Name, "127.0.0.1", 0)
			stateErr = c.SetServerState(backend.Name, srv.Name, "maint")
		} else {
			// logger.Tracef("server '%s/%s' changed status to %v", newEndpoints.BackendName, srv.Name, "ready")
			addrErr = c.SetServerAddr(backend.Name, srv.Name, srv.Address, int(backend.Endpoints.Port))
			stateErr = c.SetServerState(backend.Name, srv.Name, "ready")
		}
		if addrErr != nil || stateErr != nil {
			backend.DynUpdateFailed = true
			errors.Add(addrErr)
			errors.Add(stateErr)
		}
	}
	return errors.Result()
}
