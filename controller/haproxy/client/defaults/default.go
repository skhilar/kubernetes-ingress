package defaults

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
	GetDefaults(writer *GetDefaultsWriter) (*GetDefaultsOk, error)
	EditDefaults(writer *EditDefaultsWriter) (*EditDefaultsOk, *EditDefaultsAccepted, error)
}

func (c *Client) GetDefaults(writer *GetDefaultsWriter) (*GetDefaultsOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetDefaults",
		Reader:             NewGetDefaultsReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/defaults",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetDefaultsOk:
		return value, nil
	case *GetDefaultsDefault:
		return nil, errors.New(fmt.Sprintf("error while getting defaults code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting defaults")
}

func (c *Client) EditDefaults(writer *EditDefaultsWriter) (*EditDefaultsOk, *EditDefaultsAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "EditDefaults",
		Reader:             NewEditDefaultsReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "PUT",
		PathPattern:        "services/haproxy/configuration/defaults",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *EditDefaultsOk:
		return value, nil, nil
	case *EditDefaultsAccepted:
		return nil, value, nil
	case *EditDefaultsBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while editing defaults code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditDefaultsDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while editing defaults code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while editing defaults")
}
