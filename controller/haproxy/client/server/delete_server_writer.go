package server

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type DeleteServerWriter struct {
	Name          string
	Backend       string
	TransactionID string
	Context       context.Context
}

func NewDeleteServerWriter() *DeleteServerWriter {
	return &DeleteServerWriter{}
}

func (w *DeleteServerWriter) WithName(name string) *DeleteServerWriter {
	w.Name = name
	return w
}

func (w *DeleteServerWriter) WithBackend(backend string) *DeleteServerWriter {
	w.Backend = backend
	return w
}

func (w *DeleteServerWriter) WithTransactionID(transactionID string) *DeleteServerWriter {
	w.TransactionID = transactionID
	return w
}

func (w *DeleteServerWriter) WithContext(context context.Context) *DeleteServerWriter {
	w.Context = context
	return w
}

func (w *DeleteServerWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Name == "" {
		return nil, errors.New("name should be provided")
	}
	if w.Backend == "" {
		return nil, errors.New("backend should be provided")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction should be set")
	}
	request.SetPathParam("name", w.Name)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetPathParam("backend", w.Backend)
	return request.Send()
}
