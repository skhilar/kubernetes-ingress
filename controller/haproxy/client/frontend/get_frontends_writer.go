package frontend

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type GetFrontendsWriter struct {
	TransactionID string
	Context       context.Context
}

func NewGetFrontendsWriter() *GetFrontendsWriter {
	return &GetFrontendsWriter{}
}

func (w *GetFrontendsWriter) WithTransactionID(transactionID string) *GetFrontendsWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetFrontendsWriter) WithContext(context context.Context) *GetFrontendsWriter {
	w.Context = context
	return w
}

func (w *GetFrontendsWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID == "" {
		return nil, errors.New("transactions must be set")
	}
	request.SetQueryParam("transaction_id", w.TransactionID)
	return request.Send()
}
