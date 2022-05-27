package httprule

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditHttpResponseRuleWriter struct {
	Index         int
	ParentName    string
	ParentType    string
	TransactionID string
	ResponseRule  models.HTTPResponseRule
	Context       context.Context
}

func NewEditHttpResponseRuleWriter() *EditHttpResponseRuleWriter {
	return &EditHttpResponseRuleWriter{}
}

func (w *EditHttpResponseRuleWriter) WithIndex(index int) *EditHttpResponseRuleWriter {
	w.Index = index
	return w
}

func (w *EditHttpResponseRuleWriter) WithParentName(parentName string) *EditHttpResponseRuleWriter {
	w.ParentName = parentName
	return w
}

func (w *EditHttpResponseRuleWriter) WithParentType(parentType string) *EditHttpResponseRuleWriter {
	w.ParentType = parentType
	return w
}

func (w *EditHttpResponseRuleWriter) WithTransactionID(transactionID string) *EditHttpResponseRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *EditHttpResponseRuleWriter) WithResponseRule(responseRule models.HTTPResponseRule) *EditHttpResponseRuleWriter {
	w.ResponseRule = responseRule
	return w
}

func (w *EditHttpResponseRuleWriter) WithContext(context context.Context) *EditHttpResponseRuleWriter {
	w.Context = context
	return w
}

func (w *EditHttpResponseRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID == "" {
		return nil, errors.New("transaction should be set")
	}
	if w.ParentType == "" {
		return nil, errors.New("parent type should be provided")
	}
	if w.ParentName == "" {
		return nil, errors.New("parent name should be set")
	}
	if w.ParentType != "frontend" && w.ParentType != "backend" {
		return nil, errors.New("parent type should be frontend or backend")
	}
	request.SetQueryParam("parent_name", w.ParentName)
	request.SetQueryParam("parent_type", w.ParentType)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetPathParam("index", fmt.Sprintf("%d", w.Index))
	request.SetBody(w.ResponseRule)
	return request.Send()
}
