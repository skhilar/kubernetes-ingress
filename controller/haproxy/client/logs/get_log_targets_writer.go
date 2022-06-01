package logs

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
)

type GetLogTargetsWriter struct {
	ParentName    string
	ParentType    string
	TransactionID string
	Context       context.Context
}

func NewGetLogTargetsWriter() *GetLogTargetsWriter {
	return &GetLogTargetsWriter{}
}

func (w *GetLogTargetsWriter) WithParentName(parentName string) *GetLogTargetsWriter {
	w.ParentName = parentName
	return w
}

func (w *GetLogTargetsWriter) WithParentType(parentType string) *GetLogTargetsWriter {
	w.ParentType = parentType
	return w
}

func (w *GetLogTargetsWriter) WithTransactionID(transactionID string) *GetLogTargetsWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetLogTargetsWriter) WithContext(context context.Context) *GetLogTargetsWriter {
	w.Context = context
	return w
}

func (w *GetLogTargetsWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.ParentName == "" {
		return nil, errors.New("parent name should be set")
	}
	if w.ParentType == "" {
		return nil, errors.New("parent type should be set")
	}
	if w.ParentType != "frontend" && w.ParentType != "backend" && w.ParentType != "defaults" && w.ParentType != "global" {
		return nil, errors.New("parent type should be frontend or backend or defaults or global")
	}
	if w.TransactionID != "" {
		request.SetQueryParam("transaction_id", w.TransactionID)
	}
	request.SetQueryParam("parent_name", w.ParentName)
	request.SetQueryParam("parent_type", w.ParentType)
	return request.Send()
}
