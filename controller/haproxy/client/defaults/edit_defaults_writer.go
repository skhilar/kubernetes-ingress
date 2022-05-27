package defaults

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditDefaultsWriter struct {
	TransactionID string
	Defaults      models.Defaults
	Context       context.Context
}

func NewEditDefaultsWriter() *EditDefaultsWriter {
	return &EditDefaultsWriter{}
}

func (w *EditDefaultsWriter) WithTransactionID(transactionID string) *EditDefaultsWriter {
	w.TransactionID = transactionID
	return w
}

func (w *EditDefaultsWriter) WithDefaults(defaults models.Defaults) *EditDefaultsWriter {
	w.Defaults = defaults
	return w
}

func (w *EditDefaultsWriter) WithContext(context context.Context) *EditDefaultsWriter {
	w.Context = context
	return w
}
func (w *EditDefaultsWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID == "" {
		return nil, errors.New("transaction must be set")
	}
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetBody(w.Defaults)
	return request.Send()
}
