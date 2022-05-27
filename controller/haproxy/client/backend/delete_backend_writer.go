package backend

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type DeleteBackendWriter struct {
	TransactionID string
	Name          string
	Context       context.Context
}

func NewDeleteBackendWriter() *DeleteBackendWriter {
	return &DeleteBackendWriter{}
}

func (w *DeleteBackendWriter) WithTransactionID(transactionID string) *DeleteBackendWriter {
	w.TransactionID = transactionID
	return w
}

func (w *DeleteBackendWriter) WithName(name string) *DeleteBackendWriter {
	w.Name = name
	return w
}

func (w *DeleteBackendWriter) WithContext(context context.Context) *DeleteBackendWriter {
	w.Context = context
	return w
}

func (w *DeleteBackendWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID == "" {
		return nil, errors.New("transaction should be set")
	}
	if w.Name == "" {
		return nil, errors.New("backend name should be provided")
	}
	request.SetPathParam("name", w.Name)
	request.SetQueryParam("transaction_id", w.TransactionID)
	return request.Send()
}
