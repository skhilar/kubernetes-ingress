package backendrule

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateBackendSwitchingRuleWriter struct {
	Frontend             string
	TransactionID        string
	BackendSwitchingRule models.BackendSwitchingRule
	Context              context.Context
}

func NewCreateBackendSwitchingRuleWriter() *CreateBackendSwitchingRuleWriter {
	return &CreateBackendSwitchingRuleWriter{}
}

func (w *CreateBackendSwitchingRuleWriter) WithContext(context context.Context) *CreateBackendSwitchingRuleWriter {
	w.Context = context
	return w
}

func (w *CreateBackendSwitchingRuleWriter) WithFrontend(frontend string) *CreateBackendSwitchingRuleWriter {
	w.Frontend = frontend
	return w
}

func (w *CreateBackendSwitchingRuleWriter) WithTransactionID(transactionID string) *CreateBackendSwitchingRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *CreateBackendSwitchingRuleWriter) WithBackendSwitchingRule(backendSwitchingRule models.BackendSwitchingRule) *CreateBackendSwitchingRuleWriter {
	w.BackendSwitchingRule = backendSwitchingRule
	return w
}

func (w *CreateBackendSwitchingRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.Frontend == "" {
		return nil, errors.New("frontend should be provided")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transactions should be provided")
	}
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetQueryParam("frontend", w.Frontend)
	request.SetBody(w.BackendSwitchingRule)
	return request.Send()
}
