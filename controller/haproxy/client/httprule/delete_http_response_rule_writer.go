package httprule

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type DeleteHttpResponseRuleWriter struct {
	Index         int64
	ParentName    string
	ParentType    string
	TransactionID string
	Context       context.Context
}

func NewDeleteHttpResponseRuleWriter() *DeleteHttpResponseRuleWriter {
	return &DeleteHttpResponseRuleWriter{}
}

func (w *DeleteHttpResponseRuleWriter) WithIndex(index int64) *DeleteHttpResponseRuleWriter {
	w.Index = index
	return w
}

func (w *DeleteHttpResponseRuleWriter) WithParentName(parentName string) *DeleteHttpResponseRuleWriter {
	w.ParentName = parentName
	return w
}

func (w *DeleteHttpResponseRuleWriter) WithParentType(parentType string) *DeleteHttpResponseRuleWriter {
	w.ParentType = parentType
	return w
}

func (w *DeleteHttpResponseRuleWriter) WithTransactionID(transactionID string) *DeleteHttpResponseRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *DeleteHttpResponseRuleWriter) WithContext(context context.Context) *DeleteHttpResponseRuleWriter {
	w.Context = context
	return w
}

func (w *DeleteHttpResponseRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID == "" {
		return nil, errors.New("transaction should be set")
	}
	if w.ParentName == "" {
		return nil, errors.New("parent name should be set")
	}
	if w.ParentType == "" {
		return nil, errors.New("parent type should be set")
	}
	if w.ParentType != "frontend" && w.ParentType != "backend" {
		return nil, errors.New("parent type should be set either frontend or backend")
	}
	request.SetQueryParam("parent_name", w.ParentName)
	request.SetQueryParam("parent_type", w.ParentType)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetPathParam("index", fmt.Sprintf("%d", w.Index))
	return request.Send()
}
