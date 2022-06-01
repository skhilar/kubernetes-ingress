package bind

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type GetBindsWriter struct {
	Frontend      string
	TransactionID string
	Context       context.Context
}

func NewGetBindsWriter() *GetBindsWriter {
	return &GetBindsWriter{}
}

func (w *GetBindsWriter) WithFrontend(frontend string) *GetBindsWriter {
	w.Frontend = frontend
	return w
}

func (w *GetBindsWriter) WithTransactionID(transactionID string) *GetBindsWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetBindsWriter) WithContext(context context.Context) *GetBindsWriter {
	w.Context = context
	return w
}

func (w *GetBindsWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID != "" {
		request.SetQueryParam("transaction_id", w.TransactionID)
	}
	if w.Frontend == "" {
		return nil, errors.New("frontend must be set")
	}
	request.SetQueryParam("frontend", w.Frontend)
	return request.Send()
}
