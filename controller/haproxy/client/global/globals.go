package global

import (
	"errors"
	"fmt"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client"
)

type Client struct {
	client.HAProxyTransport
}

func NewClient(client client.HAProxyTransport) ClientService {
	return &Client{client}
}

type ClientService interface {
	GetGlobal(writer *GetGlobalWriter) (*GetGlobalOk, error)
	EditGlobal(writer *EditGlobalWriter) (*EditGlobalOk, *EditGlobalAccepted, error)
}

func (c *Client) GetGlobal(writer *GetGlobalWriter) (*GetGlobalOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetGlobal",
		Reader:             NewGetGlobalReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/global",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetGlobalOk:
		return value, nil
	case *GetGlobalDefault:
		return nil, errors.New(fmt.Sprintf("error while getting global code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting global")
}

func (c *Client) EditGlobal(writer *EditGlobalWriter) (*EditGlobalOk, *EditGlobalAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "EditGlobal",
		Reader:             NewEditGlobalReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "PUT",
		PathPattern:        "services/haproxy/configuration/global",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *EditGlobalOk:
		return value, nil, nil
	case *EditGlobalAccepted:
		return nil, value, nil
	case *EditGlobalBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while editing global code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditGlobalDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while editing global code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while editing global")
}
