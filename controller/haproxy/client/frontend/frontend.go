package frontend

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
	CreateFrontend(writer *CreateFrontendWriter) (*CreateFrontendCreated, *CreateFrontendAccepted, error)
	DeleteFrontend(writer *DeleteFrontendWriter) (*DeleteFrontendAccepted, *DeleteFrontendNoContent, error)
	EditFrontend(writer *EditFrontendWriter) (*EditFrontendOk, *EditFrontendAccepted, error)
	GetFrontend(writer *GetFrontendWriter) (*GetFrontendOk, error)
	GetFrontends(writer *GetFrontendsWriter) (*GetFrontendsOk, error)
}

func (c *Client) CreateFrontend(writer *CreateFrontendWriter) (*CreateFrontendCreated, *CreateFrontendAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CreateFrontend",
		Reader:             NewCreateFrontendReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "POST",
		PathPattern:        "services/haproxy/configuration/frontends",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateFrontendCreated:
		return value, nil, nil
	case *CreateFrontendAccepted:
		return nil, value, nil
	case *CreateFrontendBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while creating frontend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateFrontendConflict:
		return nil, nil, errors.New(fmt.Sprintf("error while creating frontend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateFrontendDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while creating frontend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while creating frontend")
}

func (c *Client) DeleteFrontend(writer *DeleteFrontendWriter) (*DeleteFrontendAccepted, *DeleteFrontendNoContent, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "DeleteFrontend",
		Reader:             NewDeleteFrontendReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "DELETE",
		PathPattern:        "services/haproxy/configuration/frontends/{name}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteFrontendAccepted:
		return value, nil, nil
	case *DeleteFrontendNoContent:
		return nil, value, nil
	case *DeleteFrontendNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting frontend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *DeleteFrontendDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting frontend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while deleting frontend")
}

func (c *Client) EditFrontend(writer *EditFrontendWriter) (*EditFrontendOk, *EditFrontendAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "EditFrontend",
		Reader:             NewEditFrontendReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "PUT",
		PathPattern:        "services/haproxy/configuration/frontends/{name}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *EditFrontendOk:
		return value, nil, nil
	case *EditFrontendAccepted:
		return nil, value, nil
	case *EditFrontendBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while editing frontend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditFrontendNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while editing frontend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditFrontendDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while editing frontend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))

	}
	return nil, nil, errors.New("unknown error while editing frontend")
}

func (c *Client) GetFrontend(writer *GetFrontendWriter) (*GetFrontendOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetFrontend",
		Reader:             NewGetFrontendReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/frontends/{name}",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetFrontendOk:
		return value, nil
	case *GetFrontendNotFound:
		return nil, errors.New(fmt.Sprintf("error while getting frontend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *GetFrontendDefault:
		return nil, errors.New(fmt.Sprintf("error while getting frontend code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting frontend")
}

func (c *Client) GetFrontends(writer *GetFrontendsWriter) (*GetFrontendsOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetFrontends",
		Reader:             NewGetFrontendsReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/frontends",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetFrontendsOk:
		return value, nil
	case *GetFrontendsDefault:
		return nil, errors.New(fmt.Sprintf("error while getting frontends code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting frontends")
}
