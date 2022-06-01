package httprule

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type GetHttpRequestRuleWriter struct {
	Index         int64
	ParentName    string
	ParentType    string
	TransactionID string
	Context       context.Context
}

func NewGetHttpRequestRuleWriter() *GetHttpRequestRuleWriter {
	return &GetHttpRequestRuleWriter{}
}

func (w *GetHttpRequestRuleWriter) WithIndex(index int64) *GetHttpRequestRuleWriter {
	w.Index = index
	return w
}

func (w *GetHttpRequestRuleWriter) WithParentName(parentName string) *GetHttpRequestRuleWriter {
	w.ParentName = parentName
	return w
}

func (w *GetHttpRequestRuleWriter) WithParentType(parentType string) *GetHttpRequestRuleWriter {
	w.ParentType = parentType
	return w
}

func (w *GetHttpRequestRuleWriter) WithTransactionID(transactionID string) *GetHttpRequestRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetHttpRequestRuleWriter) WithContext(context context.Context) *GetHttpRequestRuleWriter {
	w.Context = context
	return w
}

func (w *GetHttpRequestRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID != "" {
		request.SetQueryParam("transaction_id", w.TransactionID)
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
	request.SetPathParam("index", fmt.Sprintf("%d", w.Index))
	return request.Send()
}
