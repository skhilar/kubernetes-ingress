package backendrule

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type GetBackendSwitchingRulesWriter struct {
	Frontend      string
	TransactionID string
	Context       context.Context
}

func NewGetBackendSwitchingRulesWriter() *GetBackendSwitchingRulesWriter {
	return &GetBackendSwitchingRulesWriter{}
}

func (w *GetBackendSwitchingRulesWriter) WithFrontend(frontend string) *GetBackendSwitchingRulesWriter {
	w.Frontend = frontend
	return w
}

func (w *GetBackendSwitchingRulesWriter) WithTransactionID(transactionID string) *GetBackendSwitchingRulesWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetBackendSwitchingRulesWriter) WithContext(context context.Context) *GetBackendSwitchingRulesWriter {
	w.Context = context
	return w
}

func (w *GetBackendSwitchingRulesWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Frontend == "" {
		return nil, errors.New("frontend should be set")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction should be set")
	}
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetQueryParam("frontend", w.Frontend)
	return request.Send()
}
