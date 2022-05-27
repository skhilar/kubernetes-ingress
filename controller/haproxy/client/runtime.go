package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Runtime struct {
	client      *resty.Client
	haProxyHost string
	haProxyPort int
}

func NewTransport(haProxyHost, user, password string, haProxyPort int) HAProxyTransport {
	client := resty.New()
	client.SetBasicAuth(user, password)
	client.SetDisableWarn(true)
	return &Runtime{client: client, haProxyHost: haProxyHost, haProxyPort: haProxyPort}
}

func (r *Runtime) Execute(o *HAProxyClientOperation) (interface{}, error) {
	request := r.client.R()
	request.Method = o.Method
	request.URL = fmt.Sprintf("http://%s:%d/v2/%s", r.haProxyHost, r.haProxyPort, o.PathPattern)
	request.SetContext(o.Context)
	request.SetHeader("Content-Type", o.ProducesMediaTypes)
	request.SetHeader("Accept", o.ConsumesMediaTypes)
	resp, err := o.Writer.WriteToRequest(request)
	if err != nil {
		return nil, err
	}
	return o.Reader.ReadResponse(resp)
}
