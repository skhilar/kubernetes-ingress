package api

import (
	"context"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client/storage"
	"github.com/haproxytech/models"
)

func (c *haProxyClient) CreateCertificate(certificate string) error {
	certWriter := storage.NewCreateCertificateWriter()
	certWriter.WithContext(context.Background()).WithForceReload(true).
		WithSSLCertificate(models.SslCertificate{File: certificate})
	_, err := c.client.Storage.CreateCertificate(certWriter)
	return err
}

func (c *haProxyClient) DeleteCertificate(certName string) error {
	certWriter := storage.NewDeleteCertificateWriter()
	certWriter.WithForceReload(true).WithContext(context.Background()).WithName(certName).WithSkipReload(false)
	_, _, err := c.client.Storage.DeleteCertificate(certWriter)
	return err
}
