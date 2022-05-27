package logs

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateLogTargetWriter struct {
	ParentName    string
	ParentType    string
	TransactionID string
	LogTarget     models.LogTarget
	Context       context.Context
}

func NewCreateLogTargetWriter() *CreateLogTargetWriter {
	return &CreateLogTargetWriter{}
}

func (w *CreateLogTargetWriter) WithParentName(parentName string) *CreateLogTargetWriter {
	w.ParentName = parentName
	return w
}

func (w *CreateLogTargetWriter) WithParentType(parentType string) *CreateLogTargetWriter {
	w.ParentType = parentType
	return w
}

func (w *CreateLogTargetWriter) WithTransactionID(transactionID string) *CreateLogTargetWriter {
	w.TransactionID = transactionID
	return w
}

func (w *CreateLogTargetWriter) WithLogTarget(logTarget models.LogTarget) *CreateLogTargetWriter {
	w.LogTarget = logTarget
	return w
}

func (w *CreateLogTargetWriter) WithContext(context context.Context) *CreateLogTargetWriter {
	w.Context = context
	return w
}

func (w *CreateLogTargetWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
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
	request.SetBody(w.LogTarget)
	return request.Send()
}
