package service

import (
	"strings"

	"github.com/haproxytech/models"

	"github.com/haproxytech/kubernetes-ingress/controller/annotations/common"
	"github.com/haproxytech/kubernetes-ingress/controller/store"
)

type Cookie struct {
	name    string
	backend *models.Backend
}

func NewCookie(n string, b *models.Backend) *Cookie {
	return &Cookie{name: n, backend: b}
}

func (a *Cookie) GetName() string {
	return a.name
}

func (a *Cookie) Process(k store.K8s, annotations ...map[string]string) error {
	input := common.GetValue(a.GetName(), annotations...)
	params := strings.Fields(input)
	if len(params) == 0 {
		a.backend.Cookie = nil
		return nil
	}
	cookieName := params[0]
	a.backend.Cookie = &models.Cookie{
		Name:     &cookieName,
		Type:     "insert",
		Nocache:  true,
		Indirect: true,
		Dynamic:  true,
	}
	return nil
}
