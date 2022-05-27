package bind

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateBindWriter struct {
	TransactionID string
	Frontend      string
	Bind          models.Bind
	Context       context.Context
}

func NewCreateBindWriter() *CreateBindWriter {
	return &CreateBindWriter{}
}

func (w *CreateBindWriter) WithTransactionID(transactionID string) *CreateBindWriter {
	w.TransactionID = transactionID
	return w
}

func (w *CreateBindWriter) WithFrontend(frontend string) *CreateBindWriter {
	w.Frontend = frontend
	return w
}

func (w *CreateBindWriter) WithBind(bind models.Bind) *CreateBindWriter {
	w.Bind = bind
	return w
}

func (w *CreateBindWriter) WithContext(context context.Context) *CreateBindWriter {
	w.Context = context
	return w
}

func (w *CreateBindWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID == "" {
		return nil, errors.New("transaction should be set")
	}
	if w.Frontend == "" {
		return nil, errors.New("frontend should be set")
	}
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetQueryParam("frontend", w.Frontend)
	request.SetBody(w.Bind)
	return request.Send()
}
