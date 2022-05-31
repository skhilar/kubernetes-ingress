package tcprule

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
	CreateTCPRequestRule(writer *CreateTCPRequestRuleWriter) (*CreateTCPRequestRuleCreated, *CreateTCPRequestRuleAccepted, error)
	DeleteTCPRequestRule(writer *DeleteTCPRequestRuleWriter) (*DeleteTCPRequestRuleAccepted, *DeleteTCPRequestRuleNoContent, error)
	GetTCPRequestRules(writer *GetTCPRulesWriter) (*GetTCPRulesRequestOk, error)
}

func (c *Client) CreateTCPRequestRule(writer *CreateTCPRequestRuleWriter) (*CreateTCPRequestRuleCreated, *CreateTCPRequestRuleAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CreateTCPRequestRule",
		Reader:             NewCreateTCPRequestRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "POST",
		PathPattern:        "services/haproxy/configuration/tcp_request_rules",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateTCPRequestRuleCreated:
		return value, nil, nil
	case *CreateTCPRequestRuleAccepted:
		return nil, value, nil
	case *CreateTCPRequestRuleBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while creating tcp rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateTCPRequestRuleConflict:
		return nil, nil, errors.New(fmt.Sprintf("error while creating tcp rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateTCPRequestRuleDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while creating tcp rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while creating tcp rule")
}

func (c *Client) DeleteTCPRequestRule(writer *DeleteTCPRequestRuleWriter) (*DeleteTCPRequestRuleAccepted, *DeleteTCPRequestRuleNoContent, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "DeleteTCPRequestRule",
		Reader:             NewDeleteTCPRequestRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "DELETE",
		PathPattern:        "services/haproxy/configuration/tcp_request_rules/{index}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteTCPRequestRuleAccepted:
		return value, nil, nil
	case *DeleteTCPRequestRuleNoContent:
		return nil, value, nil
	case *DeleteTCPRequestRuleNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting tcp rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *DeleteTCPRequestRuleDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting tcp rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while deleting tcp rule")
}

func (c *Client) GetTCPRequestRules(writer *GetTCPRulesWriter) (*GetTCPRulesRequestOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetTCPRequestRules",
		Reader:             NewGetTCPRulesRequestReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "multipart/form-data",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/tcp_request_rules",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetTCPRulesRequestOk:
		return value, nil
	case *GetTCPRulesRequestDefault:
		return nil, errors.New(fmt.Sprintf("error while getting tcp request rules code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting tcp request rules")
}
