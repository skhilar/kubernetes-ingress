package server

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type UpdateRuntimeWriter struct {
	Name          string
	Backend       string
	RuntimeServer models.RuntimeServer
	Context       context.Context
}

func NewUpdateRuntimeWriter() *UpdateRuntimeWriter {
	return &UpdateRuntimeWriter{}
}

func (w *UpdateRuntimeWriter) WithName(name string) *UpdateRuntimeWriter {
	w.Name = name
	return w
}

func (w *UpdateRuntimeWriter) WithBackend(backend string) *UpdateRuntimeWriter {
	w.Backend = backend
	return w
}

func (w *UpdateRuntimeWriter) WithRuntimeServer(server models.RuntimeServer) *UpdateRuntimeWriter {
	w.RuntimeServer = server
	return w
}

func (w *UpdateRuntimeWriter) WithContext(context context.Context) *UpdateRuntimeWriter {
	w.Context = context
	return w
}

func (w *UpdateRuntimeWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	request.SetPathParam("name", w.Name)
	request.SetQueryParam("backend", w.Backend)
	request.SetBody(w.RuntimeServer)
	return request.Send()
}
