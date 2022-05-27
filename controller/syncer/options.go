package syncer

import (
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/api"
	"k8s.io/client-go/kubernetes"
)

type OptionFunc func(s *Syncer) error

func WithBackends(backends []Backend) OptionFunc {
	return func(s *Syncer) error {
		s.backends = backends
		return nil
	}
}

func WithLabelSelector(selector string) OptionFunc {
	return func(s *Syncer) error {
		s.labelSelector = selector
		return nil
	}
}

func WithAddressType(addressType string) OptionFunc {
	return func(s *Syncer) error {
		s.addressType = addressType
		return nil
	}
}

func WithK8sClient(client *kubernetes.Clientset) OptionFunc {
	return func(s *Syncer) error {
		s.client = client
		return nil
	}
}

func WithApiClient(apiClient api.HAProxyClient) OptionFunc {
	return func(s *Syncer) error {
		s.apiClient = apiClient
		return nil
	}
}
