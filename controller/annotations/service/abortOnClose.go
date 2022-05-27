package service

import (
	"github.com/haproxytech/models"

	"github.com/haproxytech/kubernetes-ingress/controller/annotations/common"
	"github.com/haproxytech/kubernetes-ingress/controller/store"
	"github.com/haproxytech/kubernetes-ingress/controller/utils"
)

type AbortOnClose struct {
	name    string
	backend *models.Backend
}

func NewAbortOnClose(n string, b *models.Backend) *AbortOnClose {
	return &AbortOnClose{name: n, backend: b}
}

func (a *AbortOnClose) GetName() string {
	return a.name
}

func (a *AbortOnClose) Process(k store.K8s, annotations ...map[string]string) error {
	input := common.GetValue(a.GetName(), annotations...)
	var enabled bool
	var err error
	if input != "" {
		enabled, err = utils.GetBoolValue(input, "abortonclose")
		if err != nil {
			return err
		}
	}
	if enabled {
		a.backend.Abortonclose = "enabled"
	} else {
		a.backend.Abortonclose = "disabled"
	}
	return nil
}
