package httprule

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditHttpRequestRuleWriter struct {
	Index         int
	ParentName    string
	ParentType    string
	TransactionID string
	RequestRule   models.HTTPRequestRule
	Context       context.Context
}

func NewEditHttpRequestRuleWriter() *EditHttpRequestRuleWriter {
	return &EditHttpRequestRuleWriter{}
}

func (w *EditHttpRequestRuleWriter) WithIndex(index int) *EditHttpRequestRuleWriter {
	w.Index = index
	return w
}

func (w *EditHttpRequestRuleWriter) WithParentName(parentName string) *EditHttpRequestRuleWriter {
	w.ParentName = parentName
	return w
}

func (w *EditHttpRequestRuleWriter) WithParentType(parentType string) *EditHttpRequestRuleWriter {
	w.ParentType = parentType
	return w
}

func (w *EditHttpRequestRuleWriter) WithTransactionID(transactionID string) *EditHttpRequestRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *EditHttpRequestRuleWriter) WithRequestRule(requestRule models.HTTPRequestRule) *EditHttpRequestRuleWriter {
	w.RequestRule = requestRule
	return w
}

func (w *EditHttpRequestRuleWriter) WithContext(context context.Context) *EditHttpRequestRuleWriter {
	w.Context = context
	return w
}

func (w *EditHttpRequestRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
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
	request.SetBody(w.RequestRule)
	return request.Send()
}
