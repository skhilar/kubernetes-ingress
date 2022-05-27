package transactions

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type CommitTransactionWriter struct {
	TransactionID string
	Context       context.Context
}

func NewCommitTransactionWriter() *CommitTransactionWriter {
	return &CommitTransactionWriter{}
}

func (w *CommitTransactionWriter) WithTransactionID(transactionID string) *CommitTransactionWriter {
	w.TransactionID = transactionID
	return w
}

func (w *CommitTransactionWriter) WithContext(context context.Context) *CommitTransactionWriter {
	w.Context = context
	return w
}

func (w *CommitTransactionWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID != "" {
		request.SetPathParam("transaction_id", w.TransactionID)
	}
	return request.Send()
}
