package service

import (
	"strconv"

	"github.com/haproxytech/models"

	"github.com/haproxytech/kubernetes-ingress/controller/annotations/common"
	"github.com/haproxytech/kubernetes-ingress/controller/store"
)

type Maxconn struct {
	name    string
	backend *models.Backend
}

func NewMaxconn(n string, b *models.Backend) *Maxconn {
	return &Maxconn{name: n, backend: b}
}

func (a *Maxconn) GetName() string {
	return a.name
}

func (a *Maxconn) Process(k store.K8s, annotations ...map[string]string) error {
	input := common.GetValue(a.GetName(), annotations...)
	if input == "" {
		if a.backend.DefaultServer != nil {
			a.backend.DefaultServer.Maxconn = nil
		}
		return nil
	}
	v, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return err
	}
	// adjust backend maxconn when using multiple HAProxy Instances
	if k.NbrHAProxyInst != 0 {
		v /= k.NbrHAProxyInst
	}
	if a.backend.DefaultServer == nil {
		a.backend.DefaultServer = &models.DefaultServer{}
	}
	a.backend.DefaultServer.Maxconn = &v
	return nil
}
