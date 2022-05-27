package global

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditGlobalWriter struct {
	TransactionID string
	Globals       models.Global
	Context       context.Context
}

func NewEditGlobalWriter() *EditGlobalWriter {
	return &EditGlobalWriter{}
}

func (w *EditGlobalWriter) WithTransactionID(transactionID string) *EditGlobalWriter {
	w.TransactionID = transactionID
	return w
}

func (w *EditGlobalWriter) WithGlobals(globals models.Global) *EditGlobalWriter {
	w.Globals = globals
	return w
}

func (w *EditGlobalWriter) WithContext(context context.Context) *EditGlobalWriter {
	w.Context = context
	return w
}

func (w *EditGlobalWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID == "" {
		return nil, errors.New("transactions must be set")
	}
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetBody(w.Globals)
	return request.Send()
}
