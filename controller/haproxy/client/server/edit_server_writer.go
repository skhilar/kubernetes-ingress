package server

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditServerWriter struct {
	Name          string
	Backend       string
	TransactionID string
	Server        models.Server
	Context       context.Context
}

func NewEditServerWriter() *EditServerWriter {
	return &EditServerWriter{}
}

func (w *EditServerWriter) WithName(name string) *EditServerWriter {
	w.Name = name
	return w
}

func (w *EditServerWriter) WithBackend(backend string) *EditServerWriter {
	w.Backend = backend
	return w
}

func (w *EditServerWriter) WithTransactionID(transactionID string) *EditServerWriter {
	w.TransactionID = transactionID
	return w
}

func (w *EditServerWriter) WithServer(server models.Server) *EditServerWriter {
	w.Server = server
	return w
}

func (w *EditServerWriter) WithContext(context context.Context) *EditServerWriter {
	w.Context = context
	return w
}

func (w *EditServerWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Name == "" {
		return nil, errors.New("name should be provided")
	}
	if w.Backend == "" {
		return nil, errors.New("backend must be set")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction must be set")
	}
	request.SetPathParam("name", w.Name)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetQueryParam("backend", w.Backend)
	request.SetBody(w.Server)
	return request.Send()
}
