package httprule

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateHttpRequestRuleWriter struct {
	ParentName    string
	ParentType    string
	TransactionID string
	RequestRule   models.HTTPRequestRule
	Context       context.Context
}

func NewCreateHttpRequestRuleWriter() *CreateHttpRequestRuleWriter {
	return &CreateHttpRequestRuleWriter{}
}

func (w *CreateHttpRequestRuleWriter) WithParentName(parentName string) *CreateHttpRequestRuleWriter {
	w.ParentName = parentName
	return w
}

func (w *CreateHttpRequestRuleWriter) WithParentType(parentType string) *CreateHttpRequestRuleWriter {
	w.ParentType = parentType
	return w
}

func (w *CreateHttpRequestRuleWriter) WithTransactionID(transactionID string) *CreateHttpRequestRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *CreateHttpRequestRuleWriter) WithRequestRule(requestRule models.HTTPRequestRule) *CreateHttpRequestRuleWriter {
	w.RequestRule = requestRule
	return w
}

func (w *CreateHttpRequestRuleWriter) WithContext(context context.Context) *CreateHttpRequestRuleWriter {
	w.Context = context
	return w
}

func (w *CreateHttpRequestRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.ParentName == "" {
		return nil, errors.New("parent name should be set")
	}
	if w.ParentType == "" {
		return nil, errors.New("parent type should be set")
	}
	if w.ParentType != "frontend" && w.ParentType != "backend" {
		return nil, errors.New("parent type should be frontend or backend")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction should be set")
	}
	request.SetQueryParam("parent_name", w.ParentName)
	request.SetQueryParam("parent_type", w.ParentType)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetBody(w.RequestRule)
	return request.Send()
}
