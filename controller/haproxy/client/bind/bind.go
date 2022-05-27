package bind

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
	CreateBind(writer *CreateBindWriter) (*CreateBindCreated, *CreateBindAccepted, error)
	DeleteBind(writer *DeleteBindWriter) (*DeleteBindAccepted, *DeleteBindNoContent, error)
	EditBind(writer *EditBindWriter) (*EditBindOk, *EditBindAccepted, error)
	GetBind(writer *GetBindWriter) (*GetBindOk, error)
	GetBinds(writer *GetBindsWriter) (*GetBindsOk, error)
}

func (c *Client) CreateBind(writer *CreateBindWriter) (*CreateBindCreated, *CreateBindAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CreateBind",
		Reader:             NewCreateBindReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "POST",
		PathPattern:        "services/haproxy/configuration/binds",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateBindCreated:
		return value, nil, nil
	case *CreateBindAccepted:
		return nil, value, nil
	case *CreateBindBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while creating bind code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateBindConflict:
		return nil, nil, errors.New(fmt.Sprintf("error while creating bind code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateBindDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while creating bind code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while creating bind")
}

func (c *Client) DeleteBind(writer *DeleteBindWriter) (*DeleteBindAccepted, *DeleteBindNoContent, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "DeleteBind",
		Reader:             NewDeleteBindReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "DELETE",
		PathPattern:        "services/haproxy/configuration/binds/{name}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteBindAccepted:
		return value, nil, nil
	case *DeleteBindNoContent:
		return nil, value, nil
	case *DeleteBindNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting bind code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *DeleteBindDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting bind code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while deleting bind")
}

func (c *Client) EditBind(writer *EditBindWriter) (*EditBindOk, *EditBindAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "EditBind",
		Reader:             NewEditBindReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "PUT",
		PathPattern:        "services/haproxy/configuration/binds/{name}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *EditBindOk:
		return value, nil, nil
	case *EditBindAccepted:
		return nil, value, nil
	case *EditBindBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while editing bind code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditBindNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while editing bind code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditBindDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while editing bind code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))

	}
	return nil, nil, errors.New("unknown error while editing bind")
}

func (c *Client) GetBind(writer *GetBindWriter) (*GetBindOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetBind",
		Reader:             NewGetBindReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/binds/{name}",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetBindOk:
		return value, nil
	case *GetBindNotFound:
		return nil, errors.New(fmt.Sprintf("error while getting bind code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *GetBindDefault:
		return nil, errors.New(fmt.Sprintf("error while getting bind code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting bind")
}

func (c *Client) GetBinds(writer *GetBindsWriter) (*GetBindsOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetBinds",
		Reader:             NewGetBindsReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/binds",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetBindsOk:
		return value, nil
	case *GetBindsDefault:
		return nil, errors.New(fmt.Sprintf("error while getting binds code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while creating binds")
}
