package httprule

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateHttpResponseRuleWriter struct {
	ParentName    string
	ParentType    string
	TransactionID string
	ResponseRule  models.HTTPResponseRule
	Context       context.Context
}

func NewCreateHttpResponseRuleWriter() *CreateHttpResponseRuleWriter {
	return &CreateHttpResponseRuleWriter{}
}

func (w *CreateHttpResponseRuleWriter) WithParentName(parentName string) *CreateHttpResponseRuleWriter {
	w.ParentName = parentName
	return w
}

func (w *CreateHttpResponseRuleWriter) WithParentType(parentType string) *CreateHttpResponseRuleWriter {
	w.ParentType = parentType
	return w
}

func (w *CreateHttpResponseRuleWriter) WithTransactionID(transactionID string) *CreateHttpResponseRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *CreateHttpResponseRuleWriter) WithResponseRule(responseRule models.HTTPResponseRule) *CreateHttpResponseRuleWriter {
	w.ResponseRule = responseRule
	return w
}

func (w *CreateHttpResponseRuleWriter) WithContext(context context.Context) *CreateHttpResponseRuleWriter {
	w.Context = context
	return w
}

func (w *CreateHttpResponseRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.ParentName == "" {
		return nil, errors.New("parent name should be set")
	}
	if w.ParentType == "" {
		return nil, errors.New("parent type should be set")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction should be set")
	}
	if w.ParentType != "frontend" && w.ParentType != "backend" {
		return nil, errors.New("parent type should be set either frontend or backend")
	}
	request.SetQueryParam("parent_name", w.ParentName)
	request.SetQueryParam("parent_type", w.ParentType)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetBody(w.ResponseRule)
	return request.Send()
}
