package api

import (
	"context"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/backend"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/backendrule"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/httprule"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/server"
	"github.com/haproxytech/models"
)

func (c *haProxyClient) BackendsGet() (models.Backends, error) {
	logger.Infof("Getting backends")
	backendsWriter := backend.NewGetBackendsWriter()
	backendsWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction)
	backends, err := c.client.Backend.GetBackends(backendsWriter)
	if err != nil {
		return models.Backends{}, err
	}
	return backends.Payload.Data, nil
}

func (c *haProxyClient) BackendGet(backendName string) (*models.Backend, error) {
	logger.Infof("Getting backend %s", backendName)
	backendWriter := backend.NewGetBackendWriter()
	backendWriter.WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithName(backendName)
	backend, err := c.client.Backend.GetBackend(backendWriter)
	if err != nil {
		logger.Errorf("Error in getting backend %s ", err)
		return nil, err
	}
	return &backend.Payload.Data, nil
}

func (c *haProxyClient) BackendCreate(b models.Backend) error {
	logger.Infof("Creating backend %s", b.Name)
	backendWriter := backend.NewCreateBackendWriter()
	backendWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction).WithData(b)
	_, _, err := c.client.Backend.CreateBackend(backendWriter)
	if err != nil {
		logger.Errorf("Error in creating backend %s", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) BackendEdit(b models.Backend) error {
	logger.Infof("Editing backend %s", b.Name)
	backendWriter := backend.NewEditBackendWriter()
	backendWriter.WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithName(b.Name).WithBackend(b)
	_, _, err := c.client.Backend.EditBackend(backendWriter)
	if err != nil {
		logger.Errorf("Error in creating editing backend %s ", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) BackendDelete(backendName string) error {
	logger.Infof("Deleting backend %s ", backendName)
	backendWriter := backend.NewDeleteBackendWriter()
	backendWriter.WithName(backendName).WithContext(context.Background()).WithTransactionID(c.activeTransaction)
	_, _, err := c.client.Backend.DeleteBackend(backendWriter)
	if err != nil {
		logger.Errorf("Error in deleting backend %s ", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) BackendCfgSnippetSet(backendName string, value []string) error {
	return nil
}

func (c *haProxyClient) BackendHTTPRequestRuleCreate(backend string, rule models.HTTPRequestRule) error {
	logger.Infof("Creating request rule for backend %s", backend)
	httpRequestRuleWriter := httprule.NewCreateHttpRequestRuleWriter()
	httpRequestRuleWriter.WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithParentType("backend").WithParentName(backend)
	_, _, err := c.client.HttpRule.CreateHttpRequestRule(httpRequestRuleWriter)
	if err != nil {
		logger.Errorf("Error in creating request rule %s ", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) BackendServerDeleteAll(backendName string) bool {
	logger.Infof("Deleting backend server for %s backend ", backendName)
	isDeleted := false
	backedServerWriter := server.NewGetServersWriter()
	backedServerWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction).WithBackend(backendName)
	servers, err := c.client.Server.GetServers(backedServerWriter)
	if err != nil {
		isDeleted = false
		return isDeleted
	}
	for _, s := range *servers.Payload.Data {
		serverDeleteWriter := server.NewDeleteServerWriter()
		serverDeleteWriter.WithBackend(backendName).WithName(s.Name).WithTransactionID(c.activeTransaction).WithContext(context.Background())
		_, _, err := c.client.Server.DeleteServer(serverDeleteWriter)
		if err != nil {
			isDeleted = false
			return isDeleted
		}
	}
	c.activeTransactionHasChanges = true
	logger.Infof("All backends are deleted ")
	return isDeleted
}

func (c *haProxyClient) BackendRuleDeleteAll(backend string) {
	logger.Infof("Deleting all rules for backend %s ", backend)
	httpRequestRuleWriter := httprule.NewGetHttpRequestRulesWriter()
	httpRequestRuleWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction).WithParentName(backend).WithParentType("backend")
	rules, err := c.client.HttpRule.GetHttpRequestRules(httpRequestRuleWriter)
	if err != nil {
		return
	}
	for _, rule := range *rules.Payload.Data {
		deleteRuleWriter := httprule.NewDeleteHttpRequestRuleWriter()
		deleteRuleWriter.WithParentType("backend").WithParentName(backend).WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithIndex(*rule.Index)
		_, _, err := c.client.HttpRule.DeleteHttpRequestRule(deleteRuleWriter)
		if err != nil {
			return
		}
	}
	c.activeTransactionHasChanges = true
	logger.Infof("All backends are deleted")
}

func (c *haProxyClient) BackendServerCreate(backendName string, data models.Server) error {
	logger.Infof("Creating server for backend %s server name %s", backendName, data.Name)
	serverWriter := server.NewCreateServerWriter()
	serverWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction).WithBackend(backendName).WithServer(data)
	_, _, err := c.client.Server.CreateServer(serverWriter)
	if err != nil {
		logger.Errorf("Error in creating backend server %s ", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) BackendServerEdit(backendName string, data models.Server) error {
	logger.Infof("Editing backend server for backend %s server name %s", backendName, data.Name)
	serverWriter := server.NewEditServerWriter()
	serverWriter.WithServer(data).WithBackend(backendName).WithName(data.Name).WithTransactionID(c.activeTransaction).WithContext(context.Background())
	_, _, err := c.client.Server.EditServer(serverWriter)
	if err != nil {
		logger.Infof("Error in editing backend server %s ", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) BackendServerDelete(backendName string, serverName string) error {
	logger.Infof("Deleting server for backend %s server %s ", backendName, serverName)
	serverWriter := server.NewDeleteServerWriter()
	serverWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction).WithName(serverName).WithBackend(backendName)
	_, _, err := c.client.Server.DeleteServer(serverWriter)
	if err != nil {
		logger.Errorf("Error in deleting server %s", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) BackendSwitchingRuleCreate(frontend string, rule models.BackendSwitchingRule) error {
	logger.Infof("Creating backend switching rule for frontend %s and rule %s", frontend, rule.Name)
	backendSwitchingRule := backendrule.NewCreateBackendSwitchingRuleWriter()
	backendSwitchingRule.WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithBackendSwitchingRule(rule).WithFrontend(frontend)
	_, _, err := c.client.BackendSwitchingRule.CreateBackendSwitchingRule(backendSwitchingRule)
	if err != nil {
		logger.Errorf("Error in creating switching rule %s ", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) BackendSwitchingRuleDeleteAll(frontend string) {
	logger.Infof("Delete backend switching rule")
	backendSwitchingRules := backendrule.NewGetBackendSwitchingRulesWriter()
	backendSwitchingRules.WithFrontend(frontend).WithContext(context.Background()).WithTransactionID(c.activeTransaction)
	rules, err := c.client.BackendSwitchingRule.GetBackendSwitchingRules(backendSwitchingRules)
	if err != nil {
		return
	}
	for _, rule := range rules.Payload.Data {
		deleteRule := backendrule.NewDeleteBackendSwitchingWriter()
		deleteRule.WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithFrontend(frontend).WithIndex(*rule.Index)
		_, _, err := c.client.BackendSwitchingRule.DeleteBackendSwitchingRule(deleteRule)
		if err != nil {
			return
		}
	}
	c.activeTransactionHasChanges = true
	logger.Infof("Backend switching rule deleted")
}

func (c *haProxyClient) ServerGet(serverName, backendName string) (models.Server, error) {
	logger.Infof("Getting server for backend %s server %s ", backendName, serverName)
	serverWriter := server.NewGetServerWriter()
	serverWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction).WithBackend(backendName).WithName(serverName)
	server, err := c.client.Server.GetServer(serverWriter)
	if err != nil {
		logger.Errorf("Error in getting server %s ", err)
		return models.Server{}, err
	}
	return *server.Data, nil
}

func (c *haProxyClient) BackendServersGet(backendName string) (models.Servers, error) {
	logger.Infof("Getting servers for backend %s ", backendName)
	serverWriter := server.NewGetServersWriter()
	serverWriter.WithBackend(backendName).WithTransactionID(c.activeTransaction).WithContext(context.Background())
	servers, err := c.client.Server.GetServers(serverWriter)
	if err != nil {
		logger.Errorf("Error in getting backed servers %s ", err)
		return models.Servers{}, err
	}
	return *servers.Payload.Data, nil
}
