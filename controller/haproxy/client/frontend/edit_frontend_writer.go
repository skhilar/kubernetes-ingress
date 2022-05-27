package frontend

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditFrontendWriter struct {
	Name          string
	TransactionID string
	Frontend      models.Frontend
	Context       context.Context
}

func NewEditFrontendWriter() *EditFrontendWriter {
	return &EditFrontendWriter{}
}

func (w *EditFrontendWriter) WithName(name string) *EditFrontendWriter {
	w.Name = name
	return w
}

func (w *EditFrontendWriter) WithTransactionID(transactionID string) *EditFrontendWriter {
	w.TransactionID = transactionID
	return w
}

func (w *EditFrontendWriter) WithFrontend(frontend models.Frontend) *EditFrontendWriter {
	w.Frontend = frontend
	return w
}

func (w *EditFrontendWriter) WithContext(context context.Context) *EditFrontendWriter {
	w.Context = context
	return w
}

func (w *EditFrontendWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Name == "" {
		return nil, errors.New("name must be provided")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction must be set")
	}
	request.SetPathParam("name", w.Name)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetBody(w.Frontend)
	return request.Send()
}
