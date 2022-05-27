package httprule

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type DeleteHttpRequestRuleWriter struct {
	Index         int64
	ParentName    string
	ParentType    string
	TransactionID string
	Context       context.Context
}

func NewDeleteHttpRequestRuleWriter() *DeleteHttpRequestRuleWriter {
	return &DeleteHttpRequestRuleWriter{}
}

func (w *DeleteHttpRequestRuleWriter) WithIndex(index int64) *DeleteHttpRequestRuleWriter {
	w.Index = index
	return w
}

func (w *DeleteHttpRequestRuleWriter) WithParentName(parentName string) *DeleteHttpRequestRuleWriter {
	w.ParentName = parentName
	return w
}

func (w *DeleteHttpRequestRuleWriter) WithParentType(parentType string) *DeleteHttpRequestRuleWriter {
	w.ParentType = parentType
	return w
}

func (w *DeleteHttpRequestRuleWriter) WithTransactionID(transactionID string) *DeleteHttpRequestRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *DeleteHttpRequestRuleWriter) WithContext(context context.Context) *DeleteHttpRequestRuleWriter {
	w.Context = context
	return w
}

func (w *DeleteHttpRequestRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID == "" {
		return nil, errors.New("transaction must be set")
	}
	if w.ParentType == "" {
		return nil, errors.New("parent type should be provided")
	}
	if w.ParentName == "" {
		return nil, errors.New("parent name should be set")
	}
	request.SetQueryParam("parent_name", w.ParentName)
	request.SetQueryParam("parent_type", w.ParentType)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetPathParam("index", fmt.Sprintf("%d", w.Index))
	return request.Send()
}
