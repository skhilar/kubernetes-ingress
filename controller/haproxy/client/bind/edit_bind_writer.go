package bind

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditBindWriter struct {
	Name          string
	Frontend      string
	TransactionID string
	Bind          models.Bind
	Context       context.Context
}

func NewEditBindWriter() *EditBindWriter {
	return &EditBindWriter{}
}

func (w *EditBindWriter) WithName(name string) *EditBindWriter {
	w.Name = name
	return w
}

func (w *EditBindWriter) WithFrontend(frontend string) *EditBindWriter {
	w.Frontend = frontend
	return w
}

func (w *EditBindWriter) WithTransactionID(transactionID string) *EditBindWriter {
	w.TransactionID = transactionID
	return w
}

func (w *EditBindWriter) WithBind(bind models.Bind) *EditBindWriter {
	w.Bind = bind
	return w
}

func (w *EditBindWriter) WithContext(context context.Context) *EditBindWriter {
	w.Context = context
	return w
}

func (w *EditBindWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID == "" {
		return nil, errors.New("transaction must be provided")
	}
	if w.Frontend == "" {
		return nil, errors.New("frontend must be provided")
	}
	if w.Name == "" {
		return nil, errors.New("name must be provided")
	}
	request.SetPathParam("name", w.Name)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetQueryParam("frontend", w.Frontend)
	request.SetBody(w.Bind)
	return request.Send()
}
