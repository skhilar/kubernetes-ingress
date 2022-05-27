package backendrule

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditBackendSwitchingRuleWriter struct {
	Frontend             string
	TransactionID        string
	Index                int
	BackendSwitchingRule models.BackendSwitchingRule
	Context              context.Context
}

func NewEditBackendSwitchingRuleWriter() *EditBackendSwitchingRuleWriter {
	return &EditBackendSwitchingRuleWriter{}
}

func (w *EditBackendSwitchingRuleWriter) WithFrontend(frontend string) *EditBackendSwitchingRuleWriter {
	w.Frontend = frontend
	return w
}

func (w *EditBackendSwitchingRuleWriter) WithTransactionID(transactionID string) *EditBackendSwitchingRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *EditBackendSwitchingRuleWriter) WithIndex(index int) *EditBackendSwitchingRuleWriter {
	w.Index = index
	return w
}

func (w *EditBackendSwitchingRuleWriter) WithBackendSwitchingRule(backendSwitchingRule models.BackendSwitchingRule) *EditBackendSwitchingRuleWriter {
	w.BackendSwitchingRule = backendSwitchingRule
	return w
}

func (w *EditBackendSwitchingRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Frontend == "" {
		return nil, errors.New("frontend should be provided")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction should be provided")
	}
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetQueryParam("frontend", w.Frontend)
	request.SetBody(w.BackendSwitchingRule)
	return request.Send()
}
