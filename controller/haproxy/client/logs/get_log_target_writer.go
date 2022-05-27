package logs

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type GetLogTargetWriter struct {
	Index         int
	ParentName    string
	ParentType    string
	TransactionID string
	Context       context.Context
}

func NewGetLogTargetWriter() *GetLogTargetWriter {
	return &GetLogTargetWriter{}
}

func (w *GetLogTargetWriter) WithParentName(parentName string) *GetLogTargetWriter {
	w.ParentName = parentName
	return w
}

func (w *GetLogTargetWriter) WithIndex(index int) *GetLogTargetWriter {
	w.Index = index
	return w
}

func (w *GetLogTargetWriter) WithParentType(parentType string) *GetLogTargetWriter {
	w.ParentType = parentType
	return w
}

func (w *GetLogTargetWriter) WithTransactionID(transactionID string) *GetLogTargetWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetLogTargetWriter) WithContext(context context.Context) *GetLogTargetWriter {
	w.Context = context
	return w
}

func (w *GetLogTargetWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
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
