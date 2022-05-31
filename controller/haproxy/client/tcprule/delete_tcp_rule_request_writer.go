package tcprule

import (
	"context"
	"github.com/go-resty/resty/v2"
	"strconv"
)

type DeleteTCPRequestRuleWriter struct {
	Index         int64
	ParentName    string
	ParentType    string
	TransactionID string
	Context       context.Context
}

func NewDeleteTCPRequestRuleWriter() *DeleteTCPRequestRuleWriter {
	return &DeleteTCPRequestRuleWriter{}
}

func (w *DeleteTCPRequestRuleWriter) WithIndex(index int64) *DeleteTCPRequestRuleWriter {
	w.Index = index
	return w
}

func (w *DeleteTCPRequestRuleWriter) WithParentName(parentName string) *DeleteTCPRequestRuleWriter {
	w.ParentName = parentName
	return w
}

func (w *DeleteTCPRequestRuleWriter) WithParentType(parentType string) *DeleteTCPRequestRuleWriter {
	w.ParentType = parentType
	return w
}

func (w *DeleteTCPRequestRuleWriter) WithTransactionID(transactionID string) *DeleteTCPRequestRuleWriter {
	w.TransactionID = transactionID
	return w
}

func (w *DeleteTCPRequestRuleWriter) WithContext(context context.Context) *DeleteTCPRequestRuleWriter {
	w.Context = context
	return w
}

func (w *DeleteTCPRequestRuleWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	request.SetPathParam("index", strconv.FormatInt(w.Index, 10))
	request.SetQueryParam("parent_name", w.ParentName)
	request.SetQueryParam("parent_type", w.ParentType)
	request.SetQueryParam("transaction_id", w.TransactionID)
	return request.Send()
}
