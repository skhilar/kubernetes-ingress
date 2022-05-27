package service

import (
	"github.com/haproxytech/models"

	"github.com/haproxytech/kubernetes-ingress/controller/annotations/common"
	"github.com/haproxytech/kubernetes-ingress/controller/store"
	"github.com/haproxytech/kubernetes-ingress/controller/utils"
)

type TimeoutCheck struct {
	name    string
	backend *models.Backend
}

func NewTimeoutCheck(n string, b *models.Backend) *TimeoutCheck {
	return &TimeoutCheck{name: n, backend: b}
}

func (a *TimeoutCheck) GetName() string {
	return a.name
}

func (a *TimeoutCheck) Process(k store.K8s, annotations ...map[string]string) error {
	input := common.GetValue(a.GetName(), annotations...)
	if input == "" {
		a.backend.CheckTimeout = nil
		return nil
	}
	timeout, err := utils.ParseTime(input)
	if err != nil {
		return err
	}
	a.backend.CheckTimeout = timeout
	return nil
}
