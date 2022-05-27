package maps

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type GetMapFileWriter struct {
	Name    string
	Context context.Context
}

func NewGetMapFileWriter() *GetMapFileWriter {
	return &GetMapFileWriter{}
}

func (w *GetMapFileWriter) WithName(name string) *GetMapFileWriter {
	w.Name = name
	return w
}

func (w *GetMapFileWriter) WithContext(context context.Context) *GetMapFileWriter {
	w.Context = context
	return w
}

func (w *GetMapFileWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	request.SetPathParam("name", w.Name)
	return request.Send()
}
