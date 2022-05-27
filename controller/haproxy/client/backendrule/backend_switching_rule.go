package backendrule

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
	CreateBackendSwitchingRule(writer *CreateBackendSwitchingRuleWriter) (*CreateBackendSwitchingRuleCreated, *CreateBackendSwitchingRuleAccepted, error)
	DeleteBackendSwitchingRule(writer *DeleteBackendSwitchingWriter) (*DeleteBackendSwitchingRuleAccepted, *DeleteBackendSwitchingRuleNoContent, error)
	EditBackendSwitchingRule(writer *EditBackendSwitchingRuleWriter) (*EditBackendSwitchingRuleOK, *EditBackendSwitchingRuleAccepted, error)
	GetBackendSwitchingRule(writer *GetBackendSwitchingRuleWriter) (*GetBackendSwitchingRuleOK, error)
	GetBackendSwitchingRules(writer *GetBackendSwitchingRulesWriter) (*GetBackendSwitchingRulesOK, error)
}

func (c *Client) CreateBackendSwitchingRule(writer *CreateBackendSwitchingRuleWriter) (*CreateBackendSwitchingRuleCreated, *CreateBackendSwitchingRuleAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CreateBackendSwitchingRule",
		Reader:             NewCreateBackendSwitchingRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "POST",
		PathPattern:        "services/haproxy/configuration/backend_switching_rules",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateBackendSwitchingRuleCreated:
		return value, nil, nil
	case *CreateBackendSwitchingRuleAccepted:
		return nil, value, nil
	case *CreateBackendSwitchingRuleBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while creating backendswitchingrule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateBackendSwitchingRuleConflict:
		return nil, nil, errors.New(fmt.Sprintf("error while creating backendswitchingrule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateBackendSwitchingRuleDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while creating backendswitchingrule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))

	}
	return nil, nil, errors.New("unknown error")

}

func (c *Client) DeleteBackendSwitchingRule(writer *DeleteBackendSwitchingWriter) (*DeleteBackendSwitchingRuleAccepted, *DeleteBackendSwitchingRuleNoContent, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "DeleteBackendSwitchingRule",
		Reader:             NewDeleteBackendSwitchingRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "DELETE",
		PathPattern:        "services/haproxy/configuration/backend_switching_rules/{index}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteBackendSwitchingRuleAccepted:
		return value, nil, nil
	case *DeleteBackendSwitchingRuleNoContent:
		return nil, value, nil
	case *DeleteBackendSwitchingRuleNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting backendswitchingrule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *DeleteBackendSwitchingRuleDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting backendswitchingrule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}

	return nil, nil, errors.New("unknown error")
}

func (c *Client) EditBackendSwitchingRule(writer *EditBackendSwitchingRuleWriter) (*EditBackendSwitchingRuleOK, *EditBackendSwitchingRuleAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "EditBackendSwitchingRule",
		Reader:             NewEditBackendSwitchingRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "PUT",
		PathPattern:        "services/haproxy/configuration/backend_switching_rules/{index}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *EditBackendSwitchingRuleOK:
		return value, nil, nil
	case *EditBackendSwitchingRuleAccepted:
		return nil, value, nil
	case *EditBackendSwitchingRuleBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while editing backendswitchingrule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditBackendSwitchingRuleNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while editing backendswitchingrule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditBackendSwitchingRuleDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while editing backendswitchingrule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error")
}

func (c *Client) GetBackendSwitchingRule(writer *GetBackendSwitchingRuleWriter) (*GetBackendSwitchingRuleOK, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetBackendSwitchingRule",
		Reader:             NewGetBackendSwitchingRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/backend_switching_rules/{index}",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetBackendSwitchingRuleOK:
		return value, nil
	case *GetBackendSwitchingRuleNotFound:
		return nil, errors.New(fmt.Sprintf("error while getting backendswitchingrule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *GetBackendSwitchingRuleDefault:
		return nil, errors.New(fmt.Sprintf("error while getting backendswitchingrule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error")
}

func (c *Client) GetBackendSwitchingRules(writer *GetBackendSwitchingRulesWriter) (*GetBackendSwitchingRulesOK, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetBackendSwitchingRule",
		Reader:             NewGetBackendSwitchingRulesReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/backend_switching_rules",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetBackendSwitchingRulesOK:
		return value, nil
	case *GetBackendSwitchingRulesDefault:
		return nil, errors.New(fmt.Sprintf("error while getting backendswitchingrules code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error")
}
