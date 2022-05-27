package service

import (
	"github.com/haproxytech/models"

	"github.com/haproxytech/kubernetes-ingress/controller/annotations/common"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/certs"
	"github.com/haproxytech/kubernetes-ingress/controller/store"
)

type Crt struct {
	name         string
	haproxyCerts *certs.Certificates
	backend      *models.Backend
}

func NewCrt(n string, c *certs.Certificates, b *models.Backend) *Crt {
	return &Crt{
		name:         n,
		haproxyCerts: c,
		backend:      b,
	}
}

func (a *Crt) GetName() string {
	return a.name
}

func (a *Crt) Process(k store.K8s, annotations ...map[string]string) error {
	var secret *store.Secret
	var crtFile string
	ns, name, err := common.GetK8sPath(a.name, annotations...)
	if err != nil {
		return err
	}
	secret, _ = k.GetSecret(ns, name)
	if secret == nil {
		if a.backend.DefaultServer != nil {
			a.backend.DefaultServer.SslCertificate = ""
			// Other values from serverSSL annotation are kept
		}
		return nil
	}
	crtFile, err = a.haproxyCerts.HandleTLSSecret(secret, certs.BD_CERT)
	if err != nil {
		return err
	}
	if a.backend.DefaultServer == nil {
		a.backend.DefaultServer = &models.DefaultServer{}
	}
	a.backend.DefaultServer.Ssl = "enabled"
	a.backend.DefaultServer.Alpn = "h2,http/1.1"
	a.backend.DefaultServer.Verify = "none"
	a.backend.DefaultServer.SslCertificate = crtFile
	return nil
}
