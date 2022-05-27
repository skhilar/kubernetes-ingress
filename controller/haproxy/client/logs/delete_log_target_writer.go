package logs

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type DeleteLogTargetWriter struct {
	Index         int
	ParentName    string
	ParentType    string
	TransactionID string
	Context       context.Context
}

func NewDeleteLogTargetWriter() *DeleteLogTargetWriter {
	return &DeleteLogTargetWriter{}
}

func (w *DeleteLogTargetWriter) WithParentName(parentName string) *DeleteLogTargetWriter {
	w.ParentName = parentName
	return w
}

func (w *DeleteLogTargetWriter) WithIndex(index int) *DeleteLogTargetWriter {
	w.Index = index
	return w
}

func (w *DeleteLogTargetWriter) WithParentType(parentType string) *DeleteLogTargetWriter {
	w.ParentType = parentType
	return w
}

func (w *DeleteLogTargetWriter) WithTransactionID(transactionID string) *DeleteLogTargetWriter {
	w.TransactionID = transactionID
	return w
}

func (w *DeleteLogTargetWriter) WithContext(context context.Context) *DeleteLogTargetWriter {
	w.Context = context
	return w
}

func (w *DeleteLogTargetWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.ParentName == "" {
		return nil, errors.New("parent name should be set")
	}
	if w.ParentType == "" {
		return nil, errors.New("parent type should be set")
	}
	if w.ParentType != "frontend" && w.ParentType != "backend" && w.ParentType != "defaults" && w.ParentType != "global" {
		return nil, errors.New("parent type should be frontend or backend or defaults or global")
	}
	if w.TransactionID == "" {
		return nil, errors.New("transaction should be set")
	}
	request.SetQueryParam("parent_name", w.ParentName)
	request.SetQueryParam("parent_type", w.ParentType)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetPathParam("index", fmt.Sprintf("%d", w.Index))
	return request.Send()
}
