package configuration

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type GetConfigurationVersionWriter struct {
	TransactionID string
	Context       context.Context
}

func NewGetConfigurationVersionWriter() *GetConfigurationVersionWriter {
	return &GetConfigurationVersionWriter{}
}

func (w *GetConfigurationVersionWriter) WithTransactionID(transactionID string) *GetConfigurationVersionWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetConfigurationVersionWriter) WithContext(context context.Context) *GetConfigurationVersionWriter {
	w.Context = context
	return w
}

func (w *GetConfigurationVersionWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID != "" {
		request.SetQueryParam(TransactionID, w.TransactionID)
	}
	return request.Send()
}
