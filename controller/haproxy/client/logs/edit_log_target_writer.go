package logs

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditLogTargetWriter struct {
	Index         int
	ParentName    string
	ParentType    string
	TransactionID string
	LogTarget     models.LogTarget
	Context       context.Context
}

func NewEditLogTargetWriter() *EditLogTargetWriter {
	return &EditLogTargetWriter{}
}

func (w *EditLogTargetWriter) WithParentName(parentName string) *EditLogTargetWriter {
	w.ParentName = parentName
	return w
}

func (w *EditLogTargetWriter) WithIndex(index int) *EditLogTargetWriter {
	w.Index = index
	return w
}

func (w *EditLogTargetWriter) WithLogTarget(logTarget models.LogTarget) *EditLogTargetWriter {
	w.LogTarget = logTarget
	return w
}

func (w *EditLogTargetWriter) WithParentType(parentType string) *EditLogTargetWriter {
	w.ParentType = parentType
	return w
}

func (w *EditLogTargetWriter) WithTransactionID(transactionID string) *EditLogTargetWriter {
	w.TransactionID = transactionID
	return w
}

func (w *EditLogTargetWriter) WithContext(context context.Context) *EditLogTargetWriter {
	w.Context = context
	return w
}

func (w *EditLogTargetWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
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
	request.SetBody(w.LogTarget)
	return request.Send()
}
