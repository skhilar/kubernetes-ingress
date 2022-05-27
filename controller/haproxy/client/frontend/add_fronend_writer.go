package frontend

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateFrontendWriter struct {
	TransactionID string
	Frontend      models.Frontend
	Context       context.Context
}

func NewCreateFrontendWriter() *CreateFrontendWriter {
	return &CreateFrontendWriter{}
}

func (w *CreateFrontendWriter) WithTransactionID(transactionID string) *CreateFrontendWriter {
	w.TransactionID = transactionID
	return w
}

func (w *CreateFrontendWriter) WithFrontend(frontend models.Frontend) *CreateFrontendWriter {
	w.Frontend = frontend
	return w
}

func (w *CreateFrontendWriter) WithContext(context context.Context) *CreateFrontendWriter {
	w.Context = context
	return w
}

func (w *CreateFrontendWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID == "" {
		return nil, errors.New("transaction must be set")
	}
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetBody(w.Frontend)
	return request.Send()
}
