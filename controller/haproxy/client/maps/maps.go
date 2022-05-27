package maps

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
	CreateMapFile(writer *CreateMapFileWriter) (*CreateMapFileCreated, error)
	DeleteMapFile(writer *DeleteMapFileWriter) (*DeleteMapFileNoContent, error)
	GetMapFile(writer *GetMapFileWriter) (*GetMapFileOk, error)
}

func (c *Client) CreateMapFile(writer *CreateMapFileWriter) (*CreateMapFileCreated, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CreateMapFile",
		Reader:             NewCreateMapFileReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "POST",
		PathPattern:        "services/haproxy/runtime/maps_entries",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *CreateMapFileCreated:
		return value, nil
	case *CreateMapFileBadRequest:
		return nil, errors.New(fmt.Sprintf("error while creating map file code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateMapFileConflict:
		return nil, errors.New(fmt.Sprintf("error while creating map file code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateMapFileDefault:
		return nil, errors.New(fmt.Sprintf("error while creating map file code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while creating map file")
}

func (c *Client) DeleteMapFile(writer *DeleteMapFileWriter) (*DeleteMapFileNoContent, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "DeleteMapFile",
		Reader:             NewDeleteMapFileReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "DELETE",
		PathPattern:        "services/haproxy/runtime/maps/{name}",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *DeleteMapFileNoContent:
		return value, nil
	case *DeleteMapFileNotFound:
		return nil, errors.New(fmt.Sprintf("error while deleting map file code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *DeleteMapFileDefault:
		return nil, errors.New(fmt.Sprintf("error while deleting map file code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while deleting map file")
}

func (c *Client) GetMapFile(writer *GetMapFileWriter) (*GetMapFileOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetMapFile",
		Reader:             NewGetMapFileReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/runtime/maps/{name}",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetMapFileOk:
		return value, nil
	case *GetMapFileNotFound:
		return nil, errors.New(fmt.Sprintf("error while getting map file code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *GetMapFileDefault:
		return nil, errors.New(fmt.Sprintf("error while getting map file code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting map file")
}
