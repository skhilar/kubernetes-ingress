package tcprule

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateTCPRequestRuleWriter struct {
	ParentName     string
	ParentType     string
	TransactionID  string
	TCPRequestRule models.TCPRequestRule
	Context        context.Context
}

func NewCreateTCPRequestRuleWriter() *CreateTCPRequestRuleWriter {
	return &CreateTCPRequestRuleWriter{}
}

func (w *CreateTCPRequestRuleWriter) WithParentName(parentName string) *CreateTCPRequestRuleWriter {
	w.ParentName = parentName
	return w
}

func (w *CreateTCPRequestRuleWriter) WithParentType(parentType string) *CreateTCPRequestRuleWriter {
	w.ParentType = parentType
	return w
}

func (w *CreateTCPRequestRuleWriter) WithTransactionID(transactionID string) *CreateTCPRequestRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *CreateTCPRequestRuleWriter) WithTCPRequestRule(tcpRequestRule models.TCPRequestRule) *CreateTCPRequestRuleWriter {
	w.TCPRequestRule = tcpRequestRule
	return w
}

func (w *CreateTCPRequestRuleWriter) WithContext(context context.Context) *CreateTCPRequestRuleWriter {
	w.Context = context
	return w
}

func (w *CreateTCPRequestRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	request.SetQueryParam("parent_name", w.ParentName)
	request.SetQueryParam("parent_type", w.ParentType)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetBody(w.TCPRequestRule)
	return request.Send()
}
