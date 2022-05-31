package backend

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type GetBackendWriter struct {
	TransactionID string
	Name          string
	Context       context.Context
}

func NewGetBackendWriter() *GetBackendWriter {
	return &GetBackendWriter{}
}

func (w *GetBackendWriter) WithTransactionID(transactionID string) *GetBackendWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetBackendWriter) WithName(name string) *GetBackendWriter {
	w.Name = name
	return w
}

func (w *GetBackendWriter) WithContext(context context.Context) *GetBackendWriter {
	w.Context = context
	return w
}

func (w *GetBackendWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Name == "" {
		return nil, errors.New("backend should be provided")
	}
	request.SetPathParam("name", w.Name)
	if w.TransactionID != "" {
		request.SetQueryParam("transaction_id", w.TransactionID)
	}
	return request.Send()
}
