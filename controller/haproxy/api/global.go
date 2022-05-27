package api

import (
	"context"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/defaults"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/global"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/logs"
	"github.com/haproxytech/models"
)

func (c *haProxyClient) DefaultsGetConfiguration() (*models.Defaults, error) {
	defaultWriter := defaults.NewGetDefaultsWriter()
	defaultWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction)
	defaults, err := c.client.Defaults.GetDefaults(defaultWriter)
	if err != nil {
		return nil, err
	}
	return defaults.Payload.Data, nil
}

func (c *haProxyClient) DefaultsPushConfiguration(d models.Defaults) error {
	defaultWriter := defaults.NewEditDefaultsWriter()
	defaultWriter.WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithDefaults(d)
	_, _, err := c.client.Defaults.EditDefaults(defaultWriter)
	if err != nil {
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) GlobalCfgSnippet(value []string) error {
	return nil
}

func (c *haProxyClient) GlobalGetLogTargets() (models.LogTargets, error) {
	logTargetWriter := logs.NewGetLogTargetsWriter()
	logTargetWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction).WithParentName("data").WithParentType("global")
	logTargets, err := c.client.LogTarget.GetLogTargets(logTargetWriter)
	if err != nil {
		return models.LogTargets{}, err
	}
	return *logTargets.Payload.Data, nil
}

func (c *haProxyClient) GlobalPushLogTargets(lg models.LogTargets) error {
	return nil
}

func (c *haProxyClient) GlobalGetConfiguration() (*models.Global, error) {
	globalWriter := global.NewGetGlobalWriter()
	globalWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction)
	global, err := c.client.Global.GetGlobal(globalWriter)
	if err != nil {
		return nil, err
	}
	return global.Payload.Data, nil
}

func (c *haProxyClient) GlobalPushConfiguration(g models.Global) error {
	globalWriter := global.NewEditGlobalWriter()
	globalWriter.WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithGlobals(g)
	_, _, err := c.client.Global.EditGlobal(globalWriter)
	if err != nil {
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}
