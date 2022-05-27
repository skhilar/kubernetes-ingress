package frontend

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type DeleteFrontendWriter struct {
	Name          string
	TransactionID string
	Context       context.Context
}

func NewDeleteFrontendWriter() *DeleteFrontendWriter {
	return &DeleteFrontendWriter{}
}

func (w *DeleteFrontendWriter) WithName(name string) *DeleteFrontendWriter {
	w.Name = name
	return w
}

func (w *DeleteFrontendWriter) WithTransactionID(transactionID string) *DeleteFrontendWriter {
	w.TransactionID = transactionID
	return w
}

func (w *DeleteFrontendWriter) WithContext(context context.Context) *DeleteFrontendWriter {
	w.Context = context
	return w
}

func (w *DeleteFrontendWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Name == "" {
		return nil, errors.New("name must be provided")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction must be set")
	}
	request.SetPathParam("name", w.Name)
	request.SetQueryParam("transaction_id", w.TransactionID)
	return request.Send()
}
