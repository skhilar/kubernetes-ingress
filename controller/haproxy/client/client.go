package client

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type HAProxyRequestWriter interface {
	WriteToRequest(request *resty.Request) (*resty.Response, error)
}

type HAProxyResponseReader interface {
	ReadResponse(response *resty.Response) (interface{}, error)
}

type HAProxyClientOperation struct {
	ID                 string
	Method             string
	PathPattern        string
	ProducesMediaTypes string
	ConsumesMediaTypes string
	Writer             HAProxyRequestWriter
	Reader             HAProxyResponseReader
	Context            context.Context
}

type HAProxyTransport interface {
	Execute(*HAProxyClientOperation) (interface{}, error)
}
