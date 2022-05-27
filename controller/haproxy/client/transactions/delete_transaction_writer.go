package transactions

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type DeleteTransactionWriter struct {
	TransactionID string
	Context       context.Context
}

func NewDeleteTransactionWriter() *DeleteTransactionWriter {
	return &DeleteTransactionWriter{}
}

func (w *DeleteTransactionWriter) WithTransactionID(transactionID string) *DeleteTransactionWriter {
	w.TransactionID = transactionID
	return w
}

func (w *DeleteTransactionWriter) WithContext(context context.Context) *DeleteTransactionWriter {
	w.Context = context
	return w
}

func (w *DeleteTransactionWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID != "" {
		request.SetPathParam("transaction_id", w.TransactionID)
	}
	return request.Send()
}
