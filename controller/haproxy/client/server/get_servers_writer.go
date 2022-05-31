package server

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type GetServersWriter struct {
	Backend       string
	TransactionID string
	Context       context.Context
}

func NewGetServersWriter() *GetServersWriter {
	return &GetServersWriter{}
}

func (w *GetServersWriter) WithBackend(backend string) *GetServersWriter {
	w.Backend = backend
	return w
}

func (w *GetServersWriter) WithTransactionID(transactionID string) *GetServersWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetServersWriter) WithContext(context context.Context) *GetServersWriter {
	w.Context = context
	return w
}

func (w *GetServersWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Backend == "" {
		return nil, errors.New("backend must be set")
	}
	if w.TransactionID != "" {
		request.SetQueryParam("transaction_id", w.TransactionID)
	}
	request.SetPathParam("backend", w.Backend)
	return request.Send()
}
