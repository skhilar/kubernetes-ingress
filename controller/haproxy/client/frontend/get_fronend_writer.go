package frontend

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type GetFrontendWriter struct {
	Name          string
	TransactionID string
	Context       context.Context
}

func NewGetFrontendWriter() *GetFrontendWriter {
	return &GetFrontendWriter{}
}

func (w *GetFrontendWriter) WithName(name string) *GetFrontendWriter {
	w.Name = name
	return w
}

func (w *GetFrontendWriter) WithTransactionID(transactionID string) *GetFrontendWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetFrontendWriter) WithContext(context context.Context) *GetFrontendWriter {
	w.Context = context
	return w
}

func (w *GetFrontendWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Name == "" {
		return nil, errors.New("name must be set")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction must be set")
	}
	request.SetPathParam("name", w.Name)
	request.SetQueryParam("transaction_id", w.TransactionID)
	return request.Send()
}
