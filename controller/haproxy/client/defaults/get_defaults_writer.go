package defaults

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type GetDefaultsWriter struct {
	TransactionID string
	Context       context.Context
}

func NewGetDefaultsWriter() *GetDefaultsWriter {
	return &GetDefaultsWriter{}
}

func (w *GetDefaultsWriter) WithTransactionID(transactionID string) *GetDefaultsWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetDefaultsWriter) WithContext(context context.Context) *GetDefaultsWriter {
	w.Context = context
	return w
}

func (w *GetDefaultsWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID != "" {
		request.SetQueryParam("transaction_id", w.TransactionID)
	}
	return request.Send()
}
