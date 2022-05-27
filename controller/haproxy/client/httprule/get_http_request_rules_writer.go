package httprule

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type GetHttpRequestRulesWriter struct {
	ParentName    string
	ParentType    string
	TransactionID string
	Context       context.Context
}

func NewGetHttpRequestRulesWriter() *GetHttpRequestRulesWriter {
	return &GetHttpRequestRulesWriter{}
}

func (w *GetHttpRequestRulesWriter) WithParentName(parentName string) *GetHttpRequestRulesWriter {
	w.ParentName = parentName
	return w
}

func (w *GetHttpRequestRulesWriter) WithParentType(parentType string) *GetHttpRequestRulesWriter {
	w.ParentType = parentType
	return w
}

func (w *GetHttpRequestRulesWriter) WithTransactionID(transactionID string) *GetHttpRequestRulesWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetHttpRequestRulesWriter) WithContext(context context.Context) *GetHttpRequestRulesWriter {
	w.Context = context
	return w
}

func (w *GetHttpRequestRulesWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
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
	return request.Send()
}
