package backendrule

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type GetBackendSwitchingRuleWriter struct {
	TransactionID string
	Frontend      string
	Index         int
	Context       context.Context
}

func NewGetBackendSwitchingRuleWriter() *GetBackendSwitchingRuleWriter {
	return &GetBackendSwitchingRuleWriter{}
}

func (w *GetBackendSwitchingRuleWriter) WithTransactionID(transactionID string) *GetBackendSwitchingRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetBackendSwitchingRuleWriter) WithFrontend(frontend string) *GetBackendSwitchingRuleWriter {
	w.Frontend = frontend
	return w
}

func (w *GetBackendSwitchingRuleWriter) WithIndex(index int) *GetBackendSwitchingRuleWriter {
	w.Index = index
	return w
}

func (w *GetBackendSwitchingRuleWriter) WithContext(context context.Context) *GetBackendSwitchingRuleWriter {
	w.Context = context
	return w
}

func (w *GetBackendSwitchingRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID == "" {
		return nil, errors.New("transaction is not set")
	}
	if w.Frontend == "" {
		return nil, errors.New("frontend is not set")
	}
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetQueryParam("frontend", w.Frontend)
	request.SetPathParam("index", fmt.Sprintf("%d", w.Index))
	return request.Send()
}
