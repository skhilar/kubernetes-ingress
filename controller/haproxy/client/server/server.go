package server

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
	CreateServer(writer *CreateServerWriter) (*CreateServerCreated, *CreateServerAccepted, error)
	DeleteServer(writer *DeleteServerWriter) (*DeleteServerAccepted, *DeleteServerNoContent, error)
	EditServer(writer *EditServerWriter) (*EditServerOk, *EditServerAccepted, error)
	GetServer(writer *GetServerWriter) (*GetServerOkBody, error)
	GetServers(writer *GetServersWriter) (*GetServersOk, error)
	UpdateRuntimeServer(writer *UpdateRuntimeWriter) (*UpdateRuntimeOk, error)
}

func (c *Client) CreateServer(writer *CreateServerWriter) (*CreateServerCreated, *CreateServerAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CreateServer",
		Reader:             NewCreateServerReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "POST",
		PathPattern:        "services/haproxy/configuration/servers",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateServerCreated:
		return value, nil, nil
	case *CreateServerAccepted:
		return nil, value, nil
	case *CreateServerBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while creating server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateServerConflict:
		return nil, nil, errors.New(fmt.Sprintf("error while creating server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateServerDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while creating server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while creating server")
}

func (c *Client) DeleteServer(writer *DeleteServerWriter) (*DeleteServerAccepted, *DeleteServerNoContent, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "DeleteServer",
		Reader:             NewDeleteServerReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "DELETE",
		PathPattern:        "services/haproxy/configuration/servers/{name}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteServerAccepted:
		return value, nil, nil
	case *DeleteServerNoContent:
		return nil, value, nil
	case *DeleteServerNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *DeleteServerDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while deleting server")
}

func (c *Client) EditServer(writer *EditServerWriter) (*EditServerOk, *EditServerAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "EditServer",
		Reader:             NewEditServerReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "PUT",
		PathPattern:        "services/haproxy/configuration/servers/{name}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *EditServerOk:
		return value, nil, nil
	case *EditServerAccepted:
		return nil, value, nil
	case *EditServerBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while editing server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateServerNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while editing server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditServerDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while editing server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while editing server")

}

func (c *Client) GetServer(writer *GetServerWriter) (*GetServerOkBody, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetServer",
		Reader:             NewGetServerReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/servers/{name}",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetServerOkBody:
		return value, nil
	case *GetServerNotFound:
		return nil, errors.New(fmt.Sprintf("error while getting server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *GetServerDefault:
		return nil, errors.New(fmt.Sprintf("error while getting server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting server")
}

func (c *Client) GetServers(writer *GetServersWriter) (*GetServersOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetServers",
		Reader:             NewGetServersReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/servers",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetServersOk:
		return value, nil
	case *GetServersDefault:
		return nil, errors.New(fmt.Sprintf("error while getting servers code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting servers")
}

func (c *Client) UpdateRuntimeServer(writer *UpdateRuntimeWriter) (*UpdateRuntimeOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "UpdateRuntimeServer",
		Reader:             NewUpdateRuntimeReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "PUT",
		PathPattern:        "services/haproxy/runtime/servers/{name}",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *UpdateRuntimeOk:
		return value, nil
	case *UpdateRuntimeBadRequest:
		return nil, errors.New(fmt.Sprintf("error while updating runtime server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *UpdateRuntimeNotFound:
		return nil, errors.New(fmt.Sprintf("error while updating runtime server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *UpdateRuntimeDefault:
		return nil, errors.New(fmt.Sprintf("error while updating runtime server code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))

	}
	return nil, errors.New("unknown error while updating runtime server")
}
