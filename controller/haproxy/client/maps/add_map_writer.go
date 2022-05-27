package maps

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
	"strconv"
)

type CreateMapFileWriter struct {
	MapFileName string
	ForeSync    bool
	Data        models.MapEntry
	Context     context.Context
}

func NewCreateMapFileWriter() *CreateMapFileWriter {
	return &CreateMapFileWriter{}
}

func (w *CreateMapFileWriter) WithFileName(mapFileName string) *CreateMapFileWriter {
	w.MapFileName = mapFileName
	return w
}

func (w *CreateMapFileWriter) WithForceSync(foreSync bool) *CreateMapFileWriter {
	w.ForeSync = foreSync
	return w
}

func (w *CreateMapFileWriter) WithData(data models.MapEntry) *CreateMapFileWriter {
	w.Data = data
	return w
}

func (w *CreateMapFileWriter) WithContext(context context.Context) *CreateMapFileWriter {
	w.Context = context
	return w
}

func (w *CreateMapFileWriter) WriteToRequest(request *resty.Request) (*resty.Response, error) {
	request.SetQueryParam("map", w.MapFileName)
	request.SetQueryParam("force_sync", strconv.FormatBool(w.ForeSync))
	request.SetBody(w.Data)
	return request.Send()
}
