package api

import (
	"context"
	"fmt"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/bind"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/frontend"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/httprule"
	"github.com/haproxytech/models"
)

func (c *haProxyClient) FrontendCfgSnippetSet(frontendName string, value []string) error {
	return nil
}

func (c *haProxyClient) FrontendCreate(f models.Frontend) error {
	logger.Infof("Creating front end %s ", f.Name)
	frontendWriter := frontend.NewCreateFrontendWriter()
	frontendWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction).WithFrontend(f)
	_, _, err := c.client.Frontend.CreateFrontend(frontendWriter)
	if err != nil {
		logger.Errorf("Error in creating front end %s ", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) FrontendDelete(frontendName string) error {
	logger.Infof("Deleting frontend %s ", frontendName)
	frontendWriter := frontend.NewDeleteFrontendWriter()
	frontendWriter.WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithName(frontendName)
	_, _, err := c.client.Frontend.DeleteFrontend(frontendWriter)
	if err != nil {
		logger.Infof("Error in deleting frontend %s ", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) FrontendsGet() (models.Frontends, error) {
	logger.Infof("Getting all frontends")
	frontendWriter := frontend.NewGetFrontendsWriter()
	frontendWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction)
	frontends, err := c.client.Frontend.GetFrontends(frontendWriter)
	if err != nil {
		logger.Errorf("Error in getting all frontends %s ", err)
		return models.Frontends{}, err
	}
	return *frontends.Payload.Data, nil
}

func (c *haProxyClient) FrontendGet(frontendName string) (models.Frontend, error) {
	logger.Infof("Getting fronened %s ", frontendName)
	frontendWriter := frontend.NewGetFrontendWriter()
	frontendWriter.WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithName(frontendName)
	frontend, err := c.client.Frontend.GetFrontend(frontendWriter)
	if err != nil {
		logger.Errorf("Error in getting frontend %s ", err)
		return models.Frontend{}, err
	}
	return *frontend.Payload.Data, nil
}

func (c *haProxyClient) FrontendEdit(f models.Frontend) error {
	logger.Infof("Editing frontend %s ", f.Name)
	frontendWriter := frontend.NewEditFrontendWriter()
	frontendWriter.WithName(f.Name).WithContext(context.Background()).WithTransactionID(c.activeTransaction).WithFrontend(f)
	_, _, err := c.client.Frontend.EditFrontend(frontendWriter)
	if err != nil {
		logger.Errorf("Error in editing frontend %s ", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) FrontendEnableSSLOffload(frontendName string, certDir string, alpn string, strictSNI bool) error {
	logger.Infof("Enabling ssl offload for frnotend %s ", frontendName)
	bindWriter := bind.NewGetBindsWriter()
	bindWriter.WithFrontend(frontendName).WithTransactionID(c.activeTransaction).WithContext(context.Background())
	binds, err := c.client.Bind.GetBinds(bindWriter)
	if err != nil {
		return err
	}
	for _, b := range *binds.Payload.Data {
		b.Ssl = true
		b.SslCertificate = certDir
		if alpn != "" {
			b.Alpn = alpn
			b.StrictSni = strictSNI
		}
		editBindWriter := bind.NewEditBindWriter()
		editBindWriter.WithContext(context.Background()).
			WithTransactionID(c.activeTransaction).WithFrontend(frontendName).
			WithName(b.Name).WithBind(*b)
		_, _, err = c.client.Bind.EditBind(editBindWriter)
		if err != nil {
			return err
		}
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) FrontendDisableSSLOffload(frontendName string) error {
	logger.Infof("Disabling ssl offload for frnotend %s ", frontendName)
	bindWriter := bind.NewGetBindsWriter()
	bindWriter.WithFrontend(frontendName).WithTransactionID(c.activeTransaction).WithContext(context.Background())
	binds, err := c.client.Bind.GetBinds(bindWriter)
	if err != nil {
		return err
	}
	for _, b := range *binds.Payload.Data {
		b.Ssl = false
		b.SslCafile = ""
		b.Verify = ""
		b.SslCertificate = ""
		b.Alpn = ""
		b.StrictSni = false
		editBindWriter := bind.NewEditBindWriter()
		editBindWriter.WithContext(context.Background()).
			WithTransactionID(c.activeTransaction).WithFrontend(frontendName).
			WithName(b.Name).WithBind(*b)
		_, _, err = c.client.Bind.EditBind(editBindWriter)
		if err != nil {
			return err
		}
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) FrontendBindsGet(frontend string) (models.Binds, error) {
	logger.Infof("Getting bind for frontend %s ", frontend)
	bindsWriter := bind.NewGetBindsWriter()
	bindsWriter.WithFrontend(frontend).WithTransactionID(c.activeTransaction).WithContext(context.Background())
	binds, err := c.client.Bind.GetBinds(bindsWriter)
	if err != nil {
		logger.Errorf("Error in getting bind %s ", err)
		return models.Binds{}, err
	}
	return *binds.Payload.Data, nil
}

func (c *haProxyClient) FrontendBindCreate(frontend string, b models.Bind) error {
	logger.Infof("Creating bind for frontend %s bind %s ", frontend, b.Name)
	bindWriter := bind.NewCreateBindWriter()
	bindWriter.WithContext(context.Background()).
		WithTransactionID(c.activeTransaction).WithFrontend(frontend).WithBind(b)
	_, _, err := c.client.Bind.CreateBind(bindWriter)
	if err != nil {
		logger.Infof("Error in creating bind %s ", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) FrontendBindEdit(frontend string, b models.Bind) error {
	logger.Infof("Editing bind for frontend %s bind %s ", frontend, b.Name)
	bindWriter := bind.NewEditBindWriter()
	bindWriter.WithBind(b).WithFrontend(frontend).
		WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithName(b.Name)
	_, _, err := c.client.Bind.EditBind(bindWriter)
	if err != nil {
		logger.Errorf("Error in editing bind %s ", err)
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) FrontendHTTPRequestRuleCreate(frontend string, rule models.HTTPRequestRule, ingressACL string) error {
	logger.Infof("Creating request rule for frontend %s ", frontend)
	requestRuleWriter := httprule.NewCreateHttpRequestRuleWriter()
	requestRuleWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction).
		WithParentName(frontend).WithParentType("frontend").WithRequestRule(rule)
	if ingressACL != "" {
		rule.Cond = "if"
		rule.CondTest = fmt.Sprintf("%s %s", ingressACL, rule.CondTest)
	}
	_, _, err := c.client.HttpRule.CreateHttpRequestRule(requestRuleWriter)
	if err != nil {
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) FrontendHTTPResponseRuleCreate(frontend string, rule models.HTTPResponseRule, ingressACL string) error {
	responseRuleWriter := httprule.NewCreateHttpResponseRuleWriter()
	responseRuleWriter.WithParentType("frontend").WithParentName(frontend).
		WithTransactionID(c.activeTransaction).WithContext(context.Background()).
		WithResponseRule(rule)
	if ingressACL != "" {
		rule.Cond = "if"
		rule.CondTest = fmt.Sprintf("%s %s", ingressACL, rule.CondTest)
	}
	_, _, err := c.client.HttpRule.CreateHttpResponseRule(responseRuleWriter)
	if err != nil {
		return err
	}
	c.activeTransactionHasChanges = true
	return nil
}

func (c *haProxyClient) FrontendTCPRequestRuleCreate(frontend string, rule models.TCPRequestRule, ingressACL string) error {
	return nil
}

func (c *haProxyClient) FrontendRuleDeleteAll(frontend string) {
	c.activeTransactionHasChanges = true
	requestRuleWriter := httprule.NewGetHttpRequestRulesWriter()
	requestRuleWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction).
		WithParentName(frontend).WithParentType("frontend")
	requestRules, err := c.client.HttpRule.GetHttpRequestRules(requestRuleWriter)
	if err != nil {
		return
	}
	for _, rule := range *requestRules.Payload.Data {
		deleteRuleWriter := httprule.NewDeleteHttpRequestRuleWriter()
		deleteRuleWriter.WithParentType("frontend").WithParentName(frontend).
			WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithIndex(*rule.Index)
		_, _, err := c.client.HttpRule.DeleteHttpRequestRule(deleteRuleWriter)
		if err != nil {
			return
		}
	}

	responseRuleWriter := httprule.NewGetHttpResponseRulesWriter()
	responseRuleWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction).
		WithParentName(frontend).WithParentType("frontend")
	responseRules, err := c.client.HttpRule.GetHttpResponseRules(responseRuleWriter)
	if err != nil {
		return
	}
	for _, rule := range *responseRules.Payload.Data {
		deleteRuleWriter := httprule.NewDeleteHttpResponseRuleWriter()
		deleteRuleWriter.WithParentType("frontend").WithParentName(frontend).
			WithTransactionID(c.activeTransaction).WithContext(context.Background()).WithIndex(*rule.Index)
		_, _, err := c.client.HttpRule.DeleteHttpResponseRule(deleteRuleWriter)
		if err != nil {
			return
		}
	}
	c.activeTransactionHasChanges = true
}
