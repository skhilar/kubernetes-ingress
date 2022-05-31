package backend

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateBackendWriter struct {
	TransactionID string
	Backend       models.Backend
	Context       context.Context
}

func NewCreateBackendWriter() *CreateBackendWriter {
	return &CreateBackendWriter{}
}

func (w *CreateBackendWriter) WithTransactionID(transactionID string) *CreateBackendWriter {
	w.TransactionID = transactionID
	return w
}

func (w *CreateBackendWriter) WithContext(context context.Context) *CreateBackendWriter {
	w.Context = context
	return w
}

func (w *CreateBackendWriter) WithData(backend models.Backend) *CreateBackendWriter {
	w.Backend = backend
	return w
}

func (w *CreateBackendWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID != "" {
		request.SetQueryParam("transaction_id", w.TransactionID)
	}
	request.SetBody(w.Backend)
	return request.Send()
}
