package maps

import (
	"context"
	"github.com/go-resty/resty/v2"
	"strconv"
)

type DeleteMapFileWriter struct {
	Name        string
	ForceDelete bool
	ForeSync    bool
	Context     context.Context
}

func NewDeleteMapFileWriter() *DeleteMapFileWriter {
	return &DeleteMapFileWriter{}
}

func (w *DeleteMapFileWriter) WithName(name string) *DeleteMapFileWriter {
	w.Name = name
	return w
}
func (w *DeleteMapFileWriter) WithForceDelete(forceDelete bool) *DeleteMapFileWriter {
	w.ForceDelete = forceDelete
	return w
}

func (w *DeleteMapFileWriter) WithForeSync(forceSync bool) *DeleteMapFileWriter {
	w.ForeSync = forceSync
	return w
}

func (w *DeleteMapFileWriter) WithContext(context context.Context) *DeleteMapFileWriter {
	w.Context = context
	return w
}

func (w *DeleteMapFileWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	request.SetPathParam("name", w.Name)
	request.SetQueryParam("forceDelete", strconv.FormatBool(w.ForceDelete))
	request.SetQueryParam("force_sync", strconv.FormatBool(w.ForeSync))
	return request.Send()
}
