package storage

import (
	"context"
	"github.com/go-resty/resty/v2"
	"strconv"
)

type DeleteCertificateWriter struct {
	Name        string
	ForceReload bool
	SkipReload  bool
	Context     context.Context
}

func NewDeleteCertificateWriter() *DeleteCertificateWriter {
	return &DeleteCertificateWriter{}
}

func (w *DeleteCertificateWriter) WithName(name string) *DeleteCertificateWriter {
	w.Name = name
	return w
}

func (w *DeleteCertificateWriter) WithForceReload(foreReload bool) *DeleteCertificateWriter {
	w.ForceReload = foreReload
	return w
}

func (w *DeleteCertificateWriter) WithSkipReload(skipReload bool) *DeleteCertificateWriter {
	w.SkipReload = skipReload
	return w
}

func (w *DeleteCertificateWriter) WithContext(context context.Context) *DeleteCertificateWriter {
	w.Context = context
	return w
}

func (w *DeleteCertificateWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	request.SetPathParam("name", w.Name)
	request.SetQueryParam("skip_reload", strconv.FormatBool(w.SkipReload))
	request.SetQueryParam("force_reload", strconv.FormatBool(w.ForceReload))
	return request.Send()
}
