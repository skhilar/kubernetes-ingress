package global

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type GetGlobalWriter struct {
	TransactionID string
	Context       context.Context
}

func NewGetGlobalWriter() *GetGlobalWriter {
	return &GetGlobalWriter{}
}

func (w *GetGlobalWriter) WithTransactionID(transactionID string) *GetGlobalWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetGlobalWriter) WithContext(context context.Context) *GetGlobalWriter {
	w.Context = context
	return w
}

func (w *GetGlobalWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID != "" {
		request.SetQueryParam("transaction_id", w.TransactionID)
	}
	return request.Send()
}
