package tcprule

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type GetTCPRulesWriter struct {
	ParentName    string
	ParentType    string
	TransactionID string
	Context       context.Context
}

func NewGetTCPRulesWriter() *GetTCPRulesWriter {
	return &GetTCPRulesWriter{}
}

func (w *GetTCPRulesWriter) WithParentName(parentName string) *GetTCPRulesWriter {
	w.ParentName = parentName
	return w
}

func (w *GetTCPRulesWriter) WithParentType(parentType string) *GetTCPRulesWriter {
	w.ParentType = parentType
	return w
}

func (w *GetTCPRulesWriter) WithTransactionID(transactionID string) *GetTCPRulesWriter {
	w.TransactionID = transactionID
	return w
}

func (w *GetTCPRulesWriter) WithContext(context context.Context) *GetTCPRulesWriter {
	w.Context = context
	return w
}

func (w *GetTCPRulesWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	request.SetQueryParam("parent_name", w.ParentName)
	request.SetQueryParam("parent_type", w.ParentType)
	if w.TransactionID != "" {
		request.SetQueryParam("transaction_id", w.TransactionID)
	}
	return request.Send()
}
