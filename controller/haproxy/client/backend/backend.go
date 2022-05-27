package backend

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
	CreateBackend(writer *CreateBackendWriter) (*CreateBackendCreated, *CreateBackendAccepted, error)
	DeleteBackend(writer *DeleteBackendWriter) (*DeleteBackendNoContent, *DeleteBackendAccepted, error)
	EditBackend(writer *EditBackendWriter) (*EditBackend, *EditBackendAccepted, error)
	GetBackend(writer *GetBackendWriter) (*GetBackend, error)
	GetBackends(writer *GetBackendsWriter) (*GetBackends, error)
}

func (c *Client) CreateBackend(writer *CreateBackendWriter) (*CreateBackendCreated, *CreateBackendAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CreateBackend",
		Reader:             NewCreateBackendReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "POST",
		PathPattern:        "services/haproxy/configuration/backends",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateBackendCreated:
		return value, nil, nil
	case *CreateBackendAccepted:
		return nil, value, nil
	case *CreateBackendBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while creating backend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateBackendConflict:
		return nil, nil, errors.New(fmt.Sprintf("error while creating backend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateBackendDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while creating backend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while creating backend")
}

func (c *Client) DeleteBackend(writer *DeleteBackendWriter) (*DeleteBackendNoContent, *DeleteBackendAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "DeleteBackend",
		Reader:             NewDeleteBackendReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "DELETE",
		PathPattern:        "services/haproxy/configuration/backends/{name}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteBackendNoContent:
		return value, nil, nil
	case *DeleteBackendAccepted:
		return nil, value, nil
	case *DeleteBackendNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while creating backend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *DeleteBackendDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while creating backend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while deleting backend")
}

func (c *Client) EditBackend(writer *EditBackendWriter) (*EditBackend, *EditBackendAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "EditBackend",
		Reader:             NewEditBackendReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "PUT",
		PathPattern:        "services/haproxy/configuration/backends/{name}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *EditBackend:
		return value, nil, nil
	case *EditBackendAccepted:
		return nil, value, nil
	case *EditBackendBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while editing backend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditBackendNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while editing backend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditBackendDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while editing backend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while editing backend")
}

func (c *Client) GetBackend(writer *GetBackendWriter) (*GetBackend, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetBackend",
		Reader:             NewGetBackendReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/backends/{name}",
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBackend)
	if ok {
		return success, nil
	}
	return nil, errors.New("not able to get backend")
}

func (c *Client) GetBackends(writer *GetBackendsWriter) (*GetBackends, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetBackends",
		Reader:             NewGetBackendsReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/backends",
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBackends)
	if ok {
		return success, nil
	}
	return nil, errors.New("not able to get backends")
}
