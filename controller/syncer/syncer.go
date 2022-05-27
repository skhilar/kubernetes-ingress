package syncer

import (
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/api"
	"github.com/haproxytech/kubernetes-ingress/controller/utils"
	"github.com/haproxytech/models"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/utils/strings/slices"
	"time"
)

const (
	labelSelector = "node-role.kubernetes.io/node="
	addressType   = "InternalIP"
)

var logger = utils.GetLogger()

type Syncer struct {
	labelSelector string
	addressType   string
	client        *kubernetes.Clientset
	apiClient     api.HAProxyClient
	factory       informers.SharedInformerFactory
	nodeSelector  labels.Selector
	lister        listers.NodeLister
	backends      []Backend
}

type Backend struct {
	Name string
	Port int
}

func New(opts ...OptionFunc) (*Syncer, error) {
	s := &Syncer{
		labelSelector: labelSelector,
		addressType:   addressType,
	}
	for _, o := range opts {
		if err := o(s); err != nil {
			return nil, err
		}
	}
	s.factory = informers.NewSharedInformerFactory(s.client, 15*time.Second)
	s.lister = s.factory.Core().V1().Nodes().Lister()
	s.nodeSelector = labels.NewSelector()
	if s.labelSelector != "" {
		selector, err := labels.Parse(s.labelSelector)
		if err != nil {
			return nil, err
		} else {
			s.nodeSelector = selector
		}
	}
	return s, nil
}

func (s *Syncer) addressByAddressType(node *core.Node) string {
	for _, address := range node.Status.Addresses {
		if string(address.Type) == s.addressType {
			return address.Address
		}
	}
	return ""
}

// Start does the main work
func (s *Syncer) Start() (chan struct{}, chan bool, error) {
	stop := make(chan struct{})
	done := make(chan bool)
	go func() {
		informer := s.factory.Core().V1().Nodes().Informer()
		informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
			AddFunc: func(_ interface{}) {
				_ = s.resync()
			},
			UpdateFunc: func(_, _ interface{}) {
				_ = s.resync()
			},
			DeleteFunc: func(_ interface{}) {
				_ = s.resync()
			},
		})
		informer.Run(stop)

		select {
		case <-done:
		}
	}()
	return stop, done, nil
}

func (s *Syncer) resync() error {
	// Kubernetes Nodes
	nodes, err := s.lister.List(s.nodeSelector)
	if err != nil {
		return err
	}
	logger.Infof("%s --> Found %d nodes in the cluster (filtered)", time.Now().Format(time.RFC3339), len(nodes))
	for _, node := range nodes {
		logger.Infof("%v %s: %v", node.Name, s.addressType, s.addressByAddressType(node))
	}
	// HAProxy Backends
	for _, be := range s.backends {
		_, err := s.apiClient.BackendGet(be.Name)
		if err != nil {
			logger.Infof("Error fetching backend %s. Skipping", be.Name)
			continue
		}
		_ = s.syncBackend(be, nodes)
	}
	return nil
}

func (s *Syncer) syncBackend(be Backend, nodes []*core.Node) error {
	// Read server info
	servers, err := s.apiClient.BackendServersGet(be.Name)
	if err != nil {
		logger.Infof("Error fetching servers from backend %s. Skipping", be.Name)
		return err
	}
	var current []string
	for _, s := range servers {
		current = append(current, s.Name)
	}

	var toAdd []*core.Node
	var toRemove []string
	var nodeList []string

	for _, node := range nodes {
		nodeList = append(nodeList, node.Name)
		if !slices.Contains(current, node.Name) { // Generics FTW!
			toAdd = append(toAdd, node)
		}
	}
	for _, c := range current {
		if !slices.Contains(nodeList, c) { // Generics FTW!
			toRemove = append(toRemove, c)
		}
	}
	if len(toAdd) == 0 && len(toRemove) == 0 { // Nothing to do
		logger.Infof("Cluster and HAProxy backend %s are in sync!", be.Name)
		return nil
	}

	s.apiClient.APIStartTransaction()
	// Add
	for _, a := range toAdd {
		logger.Infof("adding %s to %s", a.Name, be.Name)
		server := models.Server{
			Name:    a.Name,
			Address: a.Status.Addresses[0].Address,
			Port:    func() *int64 { i := int64(be.Port); return &i }(),
			Check:   "enabled",
		}
		err := s.apiClient.BackendServerCreate(be.Name, server)
		if err != nil {
			logger.Infof("failed to add server: %v", err)
		}
	}
	// Remove
	for _, r := range toRemove {
		logger.Infof("removing %s from %s", r, be.Name)
		err := s.apiClient.BackendServerDelete(be.Name, r)
		if err != nil {
			logger.Infof("failed to remove server: %v", err)
		}
	}
	s.apiClient.APICommitTransaction()
	return err
}
