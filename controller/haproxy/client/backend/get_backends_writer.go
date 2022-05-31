package backend

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type GetBackendsWriter struct {
	TransactionID string
	Context       context.Context
}

func NewGetBackendsWriter() *GetBackendsWriter {
	return &GetBackendsWriter{}
}

func (w *GetBackendsWriter) WithContext(context context.Context) *GetBackendsWriter {
	w.Context = context
	return w
}

func (w *GetBackendsWriter) WithTransactionID(transactionID string) *GetBackendsWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetBackendsWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID != "" {
		request.SetQueryParam("transaction_id", w.TransactionID)
	}
	return request.Send()
}
