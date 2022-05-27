package global

import (
	"strings"

	"github.com/haproxytech/models"

	"github.com/haproxytech/kubernetes-ingress/controller/annotations/common"
	"github.com/haproxytech/kubernetes-ingress/controller/store"
)

type LogFormat struct {
	name     string
	defaults *models.Defaults
}

func NewLogFormat(n string, d *models.Defaults) *LogFormat {
	return &LogFormat{name: n, defaults: d}
}

func (a *LogFormat) GetName() string {
	return a.name
}

func (a *LogFormat) Process(k store.K8s, annotations ...map[string]string) error {
	input := common.GetValue(a.GetName(), annotations...)
	if input != "" {
		input = "'" + strings.TrimSpace(input) + "'"
	}
	a.defaults.LogFormat = input
	return nil
}
