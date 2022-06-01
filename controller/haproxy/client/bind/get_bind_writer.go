package bind

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type GetBindWriter struct {
	Name          string
	Frontend      string
	TransactionID string
	Context       context.Context
}

func NewGetBindWriter() *GetBindWriter {
	return &GetBindWriter{}
}

func (w *GetBindWriter) WithName(name string) *GetBindWriter {
	w.Name = name
	return w
}

func (w *GetBindWriter) WithFrontend(frontend string) *GetBindWriter {
	w.Frontend = frontend
	return w
}

func (w *GetBindWriter) WithTransactionID(transactionID string) *GetBindWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetBindWriter) WithContext(context context.Context) *GetBindWriter {
	w.Context = context
	return w
}

func (w *GetBindWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Name == "" {
		return nil, errors.New("name must be provided")
	}
	if w.TransactionID != "" {
		request.SetQueryParam("transaction_id", w.TransactionID)
	}
	if w.Frontend == "" {
		return nil, errors.New("frontend must be provided")
	}
	request.SetQueryParam("frontend", w.Frontend)
	request.SetPathParam("name", w.Name)
	return request.Send()
}
