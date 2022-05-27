package server

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type GetServerWriter struct {
	Name          string
	Backend       string
	TransactionID string
	Context       context.Context
}

func NewGetServerWriter() *GetServerWriter {
	return &GetServerWriter{}
}

func (w *GetServerWriter) WithName(name string) *GetServerWriter {
	w.Name = name
	return w
}

func (w *GetServerWriter) WithBackend(backend string) *GetServerWriter {
	w.Backend = backend
	return w
}

func (w *GetServerWriter) WithTransactionID(transactionID string) *GetServerWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetServerWriter) WithContext(context context.Context) *GetServerWriter {
	w.Context = context
	return w
}

func (w *GetServerWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
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
	request.SetPathParam("backend", w.Backend)
	return request.Send()
}
