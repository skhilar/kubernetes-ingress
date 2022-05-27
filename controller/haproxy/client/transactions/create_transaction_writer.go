package transactions

import (
	"context"
	"github.com/go-resty/resty/v2"
	"strconv"
)

const (
	version = "version"
)

type CreateTransactionWriter struct {
	Version int64
	Context context.Context
}

func NewCreateTransactionWriter() *CreateTransactionWriter {
	return &CreateTransactionWriter{}
}

func (w *CreateTransactionWriter) WithVersion(version int64) *CreateTransactionWriter {
	w.Version = version
	return w
}

func (w *CreateTransactionWriter) WithContext(context context.Context) *CreateTransactionWriter {
	w.Context = context
	return w
}

func (w *CreateTransactionWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	request.SetQueryParam(version, strconv.FormatInt(w.Version, 10))
	return request.Send()
}
