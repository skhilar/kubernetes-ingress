package storage

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
	"strconv"
)

type CreateCertificateWriter struct {
	ForceReload    bool
	SSLCertificate models.SslCertificate
	Context        context.Context
}

func NewCreateCertificateWriter() *CreateCertificateWriter {
	return &CreateCertificateWriter{}
}

func (w *CreateCertificateWriter) WithForceReload(foreReload bool) *CreateCertificateWriter {
	w.ForceReload = foreReload
	return w
}

func (w *CreateCertificateWriter) WithSSLCertificate(sslCertificate models.SslCertificate) *CreateCertificateWriter {
	w.SSLCertificate = sslCertificate
	return w
}

func (w *CreateCertificateWriter) WithContext(context context.Context) *CreateCertificateWriter {
	w.Context = context
	return w
}

func (w *CreateCertificateWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	request.SetQueryParam("force_reload", strconv.FormatBool(w.ForceReload))
	request.SetFile("file_upload", w.SSLCertificate.File)
	return request.Send()
}
