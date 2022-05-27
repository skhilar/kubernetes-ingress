package transactions

import (
	"context"
	"github.com/go-resty/resty/v2"
)

const (
	status = "status"
)

type GetTransactionsWriter struct {
	Status  string
	Context context.Context
}

func NewGetTransactionsWriter(status string) *GetTransactionsWriter {
	return &GetTransactionsWriter{Status: status}
}

func NewGetTransactionsWriterWithContext(status string, context context.Context) *GetTransactionsWriter {
	return &GetTransactionsWriter{Status: status, Context: context}
}

func (w *GetTransactionsWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Status != "" {
		request.SetQueryParam(status, w.Status)
	}
	return request.Send()
}
