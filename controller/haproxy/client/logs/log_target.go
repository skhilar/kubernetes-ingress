package logs

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
	CreateLogTarget(writer *CreateLogTargetWriter) (*CreateLogTargetCreated, *CreateLogTargetAccepted, error)
	DeleteLogTarget(writer *DeleteLogTargetWriter) (*DeleteLogTargetAccepted, *DeleteLogTargetNoContent, error)
	EditLogTarget(writer *EditLogTargetWriter) (*EditLogTargetCreated, *EditLogTargetAccepted, error)
	GetLogTarget(writer *GetLogTargetWriter) (*GetLogTargetOk, error)
	GetLogTargets(writer *GetLogTargetsWriter) (*GetLogTargetsOk, error)
}

func (c *Client) CreateLogTarget(writer *CreateLogTargetWriter) (*CreateLogTargetCreated, *CreateLogTargetAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CreateLogTarget",
		Reader:             NewCreateLogTargetReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "POST",
		PathPattern:        "services/haproxy/configuration/log_targets",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateLogTargetCreated:
		return value, nil, nil
	case *CreateLogTargetAccepted:
		return nil, value, nil
	case *CreateLogTargetBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while creating log target code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateLogTargetConflict:
		return nil, nil, errors.New(fmt.Sprintf("error while creating log target code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateLogTargetDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while creating log target code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while creating log target")
}

func (c *Client) DeleteLogTarget(writer *DeleteLogTargetWriter) (*DeleteLogTargetAccepted, *DeleteLogTargetNoContent, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "DeleteLogTarget",
		Reader:             NewDeleteLogTargetReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "DELETE",
		PathPattern:        "services/haproxy/configuration/log_targets/{index}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteLogTargetAccepted:
		return value, nil, nil
	case *DeleteLogTargetNoContent:
		return nil, value, nil
	case *DeleteLogTargetNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting log target code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *DeleteLogTargetDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting log target code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while deleting log target")
}

func (c *Client) EditLogTarget(writer *EditLogTargetWriter) (*EditLogTargetCreated, *EditLogTargetAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "EditLogTarget",
		Reader:             NewEditLogTargetReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "PUT",
		PathPattern:        "services/haproxy/configuration/log_targets/{index}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *EditLogTargetCreated:
		return value, nil, nil
	case *EditLogTargetAccepted:
		return nil, value, nil
	case *EditLogTargetBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while editing log target code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditLogTargetDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while editing log target code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while editing log target")
}

func (c *Client) GetLogTarget(writer *GetLogTargetWriter) (*GetLogTargetOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetLogTarget",
		Reader:             NewGetLogTargetReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/log_targets/{index}",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetLogTargetOk:
		return value, nil
	case *GetLogTargetNotFound:
		return nil, errors.New(fmt.Sprintf("error while getting log target code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *GetLogTargetDefault:
		return nil, errors.New(fmt.Sprintf("error while getting log target code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting log target")
}

func (c *Client) GetLogTargets(writer *GetLogTargetsWriter) (*GetLogTargetsOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetLogTargets",
		Reader:             NewGetLogTargetsReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/log_targets",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetLogTargetsOk:
		return value, nil
	case *GetLogTargetsDefault:
		return nil, errors.New(fmt.Sprintf("error while getting log targets code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting log targets")
}
