package dataplane

import (
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/backend"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/backendrule"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/bind"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/configuration"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/defaults"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/frontend"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/global"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/httprule"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/logs"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/maps"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/server"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/transactions"
)

func NewApiClient(haProxyHost, user, password string, haProxyPort int) *ApiClient {
	runtime := client.NewTransport(haProxyHost, user, password, haProxyPort)
	apiClient := &ApiClient{}
	apiClient.Configuration = configuration.NewClient(runtime)
	apiClient.Transaction = transactions.NewClient(runtime)
	apiClient.Backend = backend.NewClient(runtime)
	apiClient.BackendSwitchingRule = backendrule.NewClient(runtime)
	apiClient.Bind = bind.NewClient(runtime)
	apiClient.Defaults = defaults.NewClient(runtime)
	apiClient.Frontend = frontend.NewClient(runtime)
	apiClient.Global = global.NewClient(runtime)
	apiClient.HttpRule = httprule.NewClient(runtime)
	apiClient.LogTarget = logs.NewClient(runtime)
	apiClient.Server = server.NewClient(runtime)
	apiClient.Maps = maps.NewClient(runtime)
	return apiClient
}

type ApiClient struct {
	Configuration        configuration.ClientService
	Transaction          transactions.ClientService
	Backend              backend.ClientService
	BackendSwitchingRule backendrule.ClientService
	Bind                 bind.ClientService
	Defaults             defaults.ClientService
	Frontend             frontend.ClientService
	Global               global.ClientService
	HttpRule             httprule.ClientService
	LogTarget            logs.ClientService
	Server               server.ClientService
	Maps                 maps.ClientService
}
