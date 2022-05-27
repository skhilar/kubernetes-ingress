package server

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateServerWriter struct {
	Backend       string
	TransactionID string
	Server        models.Server
	Context       context.Context
}

func NewCreateServerWriter() *CreateServerWriter {
	return &CreateServerWriter{}
}

func (w *CreateServerWriter) WithBackend(backend string) *CreateServerWriter {
	w.Backend = backend
	return w
}

func (w *CreateServerWriter) WithTransactionID(transactionID string) *CreateServerWriter {
	w.TransactionID = transactionID
	return w
}

func (w *CreateServerWriter) WithServer(server models.Server) *CreateServerWriter {
	w.Server = server
	return w
}

func (w *CreateServerWriter) WithContext(context context.Context) *CreateServerWriter {
	w.Context = context
	return w
}

func (w *CreateServerWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Backend == "" {
		return nil, errors.New("backend must be set")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction must be set")
	}
	request.SetQueryParam("backend", w.Backend)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetBody(w.Server)
	return request.Send()
}
