package handler

import (
	config "github.com/haproxytech/kubernetes-ingress/controller/configuration"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/api"
	"github.com/haproxytech/kubernetes-ingress/controller/store"
)

type CleanUp struct {
}

func (h CleanUp) Update(k store.K8s, cfg *config.ControllerCfg, api api.HAProxyClient) (bool, error) {
	logger.Infof("Cleaning all pending transactions")
	err := api.DeleteAllTransactions()
	return true, err
}
