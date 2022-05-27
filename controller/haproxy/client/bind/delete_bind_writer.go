package bind

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type DeleteBindWriter struct {
	Name          string
	Frontend      string
	TransactionID string
	Context       context.Context
}

func NewDeleteBindWriter() *DeleteBindWriter {
	return &DeleteBindWriter{}
}

func (w *DeleteBindWriter) WithName(name string) *DeleteBindWriter {
	w.Name = name
	return w
}

func (w *DeleteBindWriter) WithFrontend(frontend string) *DeleteBindWriter {
	w.Frontend = frontend
	return w
}

func (w *DeleteBindWriter) WithTransactionID(transactionID string) *DeleteBindWriter {
	w.TransactionID = transactionID
	return w
}

func (w *DeleteBindWriter) WithContext(context context.Context) *DeleteBindWriter {
	w.Context = context
	return w
}

func (w *DeleteBindWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Name == "" {
		return nil, errors.New("name must be provided")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction must be provided")
	}
	if w.Frontend == "" {
		return nil, errors.New("frontend must be provided")
	}
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetQueryParam("frontend", w.Frontend)
	request.SetPathParam("name", w.Name)
	return request.Send()
}
