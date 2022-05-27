package backend

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditBackendWriter struct {
	TransactionID string
	Name          string
	Backend       models.Backend
	Context       context.Context
}

func NewEditBackendWriter() *EditBackendWriter {
	return &EditBackendWriter{}
}

func (w *EditBackendWriter) WithContext(context context.Context) *EditBackendWriter {
	w.Context = context
	return w
}

func (w *EditBackendWriter) WithTransactionID(transactionID string) *EditBackendWriter {
	w.TransactionID = transactionID
	return w
}

func (w *EditBackendWriter) WithName(name string) *EditBackendWriter {
	w.Name = name
	return w
}

func (w *EditBackendWriter) WithBackend(backend models.Backend) *EditBackendWriter {
	w.Backend = backend
	return w
}

func (w *EditBackendWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Name == "" {
		return nil, errors.New("backend name should be provided")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction should be provided")
	}
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetPathParam("name", w.Name)
	request.SetBody(w.Backend)
	return request.Send()
}
