package transactions

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type GetTransactionWriter struct {
	TransactionID string
	Context       context.Context
}

func NewGetTransactionWriter(transactionID string) *GetTransactionWriter {
	return &GetTransactionWriter{TransactionID: transactionID}
}

func NewGetTransactionWriterWithContext(transactionID string, context context.Context) *GetTransactionWriter {
	return &GetTransactionWriter{TransactionID: transactionID, Context: context}
}

func (w *GetTransactionWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID != "" {
		request.SetPathParam("transaction_id", w.TransactionID)
	}
	return request.Send()
}
