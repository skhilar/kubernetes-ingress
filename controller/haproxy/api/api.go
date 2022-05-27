package api

import (
	"context"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/configuration"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/dataplane"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/transactions"
	"github.com/haproxytech/kubernetes-ingress/controller/store"
	"github.com/haproxytech/kubernetes-ingress/controller/utils"
	"github.com/haproxytech/models"
)

var logger = utils.GetLogger()

type HAProxyClient interface {
	GetConfigVersion() (int64, error)
	APIStartTransaction() error
	APICommitTransaction() error
	APIDisposeTransaction()
	BackendsGet() (models.Backends, error)
	BackendGet(backendName string) (*models.Backend, error)
	BackendCreate(backend models.Backend) error
	BackendEdit(backend models.Backend) error
	BackendDelete(backendName string) error
	BackendCfgSnippetSet(backendName string, value []string) error
	BackendHTTPRequestRuleCreate(backend string, rule models.HTTPRequestRule) error
	BackendRuleDeleteAll(backend string)
	BackendServerDeleteAll(backendName string) (deleteServers bool)
	BackendServerCreate(backendName string, data models.Server) error
	BackendServerEdit(backendName string, data models.Server) error
	BackendServerDelete(backendName string, serverName string) error
	BackendServersGet(backendName string) (models.Servers, error)
	BackendSwitchingRuleCreate(frontend string, rule models.BackendSwitchingRule) error
	BackendSwitchingRuleDeleteAll(frontend string)
	DefaultsGetConfiguration() (*models.Defaults, error)
	DefaultsPushConfiguration(models.Defaults) error
	ExecuteRaw(command string) (result []string, err error)
	FrontendCfgSnippetSet(frontendName string, value []string) error
	FrontendCreate(frontend models.Frontend) error
	FrontendDelete(frontendName string) error
	FrontendsGet() (models.Frontends, error)
	FrontendGet(frontendName string) (models.Frontend, error)
	FrontendEdit(frontend models.Frontend) error
	FrontendEnableSSLOffload(frontendName string, certDir string, alpn string, strictSNI bool) (err error)
	FrontendDisableSSLOffload(frontendName string) (err error)
	FrontendBindsGet(frontend string) (models.Binds, error)
	FrontendBindCreate(frontend string, bind models.Bind) error
	FrontendBindEdit(frontend string, bind models.Bind) error
	FrontendHTTPRequestRuleCreate(frontend string, rule models.HTTPRequestRule, ingressACL string) error
	FrontendHTTPResponseRuleCreate(frontend string, rule models.HTTPResponseRule, ingressACL string) error
	FrontendTCPRequestRuleCreate(frontend string, rule models.TCPRequestRule, ingressACL string) error
	FrontendRuleDeleteAll(frontend string)
	GlobalGetLogTargets() (models.LogTargets, error)
	GlobalPushLogTargets(models.LogTargets) error
	GlobalGetConfiguration() (*models.Global, error)
	GlobalPushConfiguration(models.Global) error
	GlobalCfgSnippet(snippet []string) error
	GetMap(mapFile string) (*models.Map, error)
	DeleteMap(mapFile string) error
	SetMapContent(mapFile string, key string, value string) error
	SetServerAddr(backendName string, serverName string, ip string, port int) error
	SetServerState(backendName string, serverName string, state string) error
	ServerGet(serverName, backendName string) (models.Server, error)
	SetAuxCfgFile(auxCfgFile string)
	SyncBackendSrvs(backend *store.RuntimeBackend, portUpdated bool) error
	UserListDeleteAll() error
	UserListExistsByGroup(group string) (bool, error)
	UserListCreateByGroup(group string, userPasswordMap map[string][]byte) error
}

type haProxyClient struct {
	client                      *dataplane.ApiClient
	activeTransaction           string
	activeTransactionHasChanges bool
}

func NewHAProxyClient(haProxyHost, userName, password string, haProxyPort int) HAProxyClient {
	return &haProxyClient{client: dataplane.NewApiClient(haProxyHost, userName, password, haProxyPort), activeTransaction: "",
		activeTransactionHasChanges: false,
	}
}

func (c *haProxyClient) GetConfigVersion() (int64, error) {
	configVersion, err := c.client.Configuration.GetConfigurationVersion(configuration.NewGetConfigurationVersionWriter())
	if err != nil {
		return 0, err
	}
	return configVersion.Version, nil
}
func (c *haProxyClient) APIStartTransaction() error {
	logger.Infof("Transaction started")
	configVersion, err := c.client.Configuration.GetConfigurationVersion(configuration.NewGetConfigurationVersionWriter())
	if err != nil {
		return err
	}
	transactionWriter := transactions.NewCreateTransactionWriter()
	transactionWriter.WithVersion(configVersion.Version).WithContext(context.Background())
	activeTransaction, err := c.client.Transaction.CreateTransaction(transactionWriter)
	if err != nil {
		return err
	}
	c.activeTransaction = activeTransaction.Payload.ID
	return nil
}

func (c *haProxyClient) APICommitTransaction() error {
	logger.Infof("committing transaction")
	if !c.activeTransactionHasChanges {
		logger.Infof("Deleting transaction as there is no change")
		deleteTransactionWriter := transactions.NewDeleteTransactionWriter()
		deleteTransactionWriter.WithContext(context.Background()).WithTransactionID(c.activeTransaction)
		_, err := c.client.Transaction.DeleteTransaction(deleteTransactionWriter)
		return err
	}
	commitWriter := transactions.NewCommitTransactionWriter()
	commitWriter.WithTransactionID(c.activeTransaction).WithContext(context.Background())
	_, _, err := c.client.Transaction.CommitTransaction(commitWriter)
	logger.Infof("Transaction committed")
	return err
}

func (c *haProxyClient) APIDisposeTransaction() {
	c.activeTransaction = ""
	c.activeTransactionHasChanges = false
}

func (c *haProxyClient) SetAuxCfgFile(auxCfgFile string) {
	//NA
}
