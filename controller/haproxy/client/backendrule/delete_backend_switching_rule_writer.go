package backendrule

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type DeleteBackendSwitchingWriter struct {
	Frontend      string
	TransactionID string
	Index         int64
	Context       context.Context
}

func NewDeleteBackendSwitchingWriter() *DeleteBackendSwitchingWriter {
	return &DeleteBackendSwitchingWriter{}
}

func (w *DeleteBackendSwitchingWriter) WithContext(context context.Context) *DeleteBackendSwitchingWriter {
	w.Context = context
	return w
}

func (w *DeleteBackendSwitchingWriter) WithFrontend(frontend string) *DeleteBackendSwitchingWriter {
	w.Frontend = frontend
	return w
}

func (w *DeleteBackendSwitchingWriter) WithTransactionID(transactionID string) *DeleteBackendSwitchingWriter {
	w.TransactionID = transactionID
	return w
}

func (w *DeleteBackendSwitchingWriter) WithIndex(index int64) *DeleteBackendSwitchingWriter {
	w.Index = index
	return w
}

func (w *DeleteBackendSwitchingWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	if w.TransactionID == "" {
		return nil, errors.New("transactions should be set")
	}
	if w.Frontend == "" {
		return nil, errors.New("frontend should be set")
	}
	request.SetQueryParam("frontend", w.Frontend)
	request.SetQueryParam("transaction_id", w.TransactionID)
	request.SetPathParam("index", fmt.Sprintf("%d", w.Index))
	return request.Send()
}
