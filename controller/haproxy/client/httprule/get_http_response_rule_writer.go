package httprule

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type GetHttpResponseRuleWriter struct {
	Index         int
	ParentName    string
	ParentType    string
	TransactionID string
	Context       context.Context
}

func NewGetHttpResponseRuleWriter() *GetHttpResponseRuleWriter {
	return &GetHttpResponseRuleWriter{}
}

func (w *GetHttpResponseRuleWriter) WithIndex(index int) *GetHttpResponseRuleWriter {
	w.Index = index
	return w
}

func (w *GetHttpResponseRuleWriter) WithParentName(parentName string) *GetHttpResponseRuleWriter {
	w.ParentName = parentName
	return w
}

func (w *GetHttpResponseRuleWriter) WithParentType(parentType string) *GetHttpResponseRuleWriter {
	w.ParentType = parentType
	return w
}

func (w *GetHttpResponseRuleWriter) WithTransactionID(transactionID string) *GetHttpResponseRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetHttpResponseRuleWriter) WithContext(context context.Context) *GetHttpResponseRuleWriter {
	w.Context = context
	return w
}

func (w *GetHttpResponseRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
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
	return request.Send()
}
