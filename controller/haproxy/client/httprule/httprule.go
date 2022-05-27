package httprule

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
	CreateHttpRequestRule(writer *CreateHttpRequestRuleWriter) (*CreateHttpRequestRuleCreated, *CreateHttpRequestRuleAccepted, error)
	CreateHttpResponseRule(writer *CreateHttpResponseRuleWriter) (*CreateHttpResponseRuleCreated, *CreateHttpResponseRuleAccepted, error)
	DeleteHttpRequestRule(writer *DeleteHttpRequestRuleWriter) (*DeleteHttpRequestRuleAccepted, *DeleteHttpRequestRuleNoContent, error)
	DeleteHttpResponseRule(writer *DeleteHttpResponseRuleWriter) (*DeleteHttpResponseRuleAccepted, *DeleteHttpResponseRuleNoContent, error)
	EditHttpRequestRule(writer *EditHttpRequestRuleWriter) (*EditHttpRequestRuleOk, *EditHttpRequestRuleAccepted, error)
	EditHttpResponseRule(writer *EditHttpResponseRuleWriter) (*EditHttpResponseRuleOk, *EditHttpResponseRuleAccepted, error)
	GetHttpRequestRule(writer *GetHttpRequestRuleWriter) (*GetHttpRequestRuleOk, error)
	GetHttpResponseRule(writer *GetHttpResponseRuleWriter) (*GetHttpResponseRuleOk, error)
	GetHttpRequestRules(writer *GetHttpRequestRulesWriter) (*GetHttpRequestRulesOk, error)
	GetHttpResponseRules(writer *GetHttpResponseRulesWriter) (*GetHttpResponseRulesOk, error)
}

func (c *Client) CreateHttpRequestRule(writer *CreateHttpRequestRuleWriter) (*CreateHttpRequestRuleCreated, *CreateHttpRequestRuleAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CreateHttpRequestRule",
		Reader:             NewCreateHttpRequestRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "POST",
		PathPattern:        "services/haproxy/configuration/http_request_rules",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateHttpRequestRuleCreated:
		return value, nil, nil
	case *CreateHttpRequestRuleAccepted:
		return nil, value, nil
	case *CreateHttpRequestRuleBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while creating http request rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateHttpRequestRuleConflict:
		return nil, nil, errors.New(fmt.Sprintf("error while creating http request rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateHttpRequestRuleDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while creating http request rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while creating http request rule")
}

func (c *Client) CreateHttpResponseRule(writer *CreateHttpResponseRuleWriter) (*CreateHttpResponseRuleCreated, *CreateHttpResponseRuleAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CreateHttpResponseRule",
		Reader:             NewCreateHttpResponseRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "POST",
		PathPattern:        "services/haproxy/configuration/http_response_rules",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateHttpResponseRuleCreated:
		return value, nil, nil
	case *CreateHttpResponseRuleAccepted:
		return nil, value, nil
	case *CreateHttpResponseRuleBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while creating http response rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateHttpResponseRuleConflict:
		return nil, nil, errors.New(fmt.Sprintf("error while creating http response rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateHttpResponseRuleDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while creating http response rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while creating http response rule")
}

func (c *Client) DeleteHttpRequestRule(writer *DeleteHttpRequestRuleWriter) (*DeleteHttpRequestRuleAccepted, *DeleteHttpRequestRuleNoContent, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "DeleteHttpRequestRule",
		Reader:             NewDeleteHttpRequestRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "DELETE",
		PathPattern:        "services/haproxy/configuration/http_request_rules/{index}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteHttpRequestRuleAccepted:
		return value, nil, nil
	case *DeleteHttpRequestRuleNoContent:
		return nil, value, nil
	case *DeleteHttpRequestRuleNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting http request rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *DeleteHttpRequestRuleDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while creating http request rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while deleting http request rule")
}

func (c *Client) DeleteHttpResponseRule(writer *DeleteHttpResponseRuleWriter) (*DeleteHttpResponseRuleAccepted, *DeleteHttpResponseRuleNoContent, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "DeleteHttpResponseRule",
		Reader:             NewDeleteHttpResponseRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "DELETE",
		PathPattern:        "services/haproxy/configuration/http_response_rules/{index}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteHttpResponseRuleAccepted:
		return value, nil, nil
	case *DeleteHttpResponseRuleNoContent:
		return nil, value, nil
	case *DeleteHttpResponseRuleNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting http response rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *DeleteHttpResponseRuleDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while response http request rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while deleting http response rule")
}

func (c *Client) EditHttpRequestRule(writer *EditHttpRequestRuleWriter) (*EditHttpRequestRuleOk, *EditHttpRequestRuleAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "EditHttpRequestRule",
		Reader:             NewEditHttpRequestRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "PUT",
		PathPattern:        "services/haproxy/configuration/http_request_rules/{index}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *EditHttpRequestRuleOk:
		return value, nil, nil
	case *EditHttpRequestRuleAccepted:
		return nil, value, nil
	case *EditHttpRequestRuleBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while editing http request rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditHttpRequestRuleNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while editing http request rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditHttpRequestRuleDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while editing http request rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while editing http request rule")
}

func (c *Client) EditHttpResponseRule(writer *EditHttpResponseRuleWriter) (*EditHttpResponseRuleOk, *EditHttpResponseRuleAccepted, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "EditHttpResponseRule",
		Reader:             NewEditHttpResponseRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "PUT",
		PathPattern:        "services/haproxy/configuration/http_response_rules/{index}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *EditHttpResponseRuleOk:
		return value, nil, nil
	case *EditHttpResponseRuleAccepted:
		return nil, value, nil
	case *EditHttpResponseRuleBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while editing http response rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditHttpResponseRuleNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while editing http response rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *EditHttpResponseRuleDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while editing http response rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while editing http response rule")
}

func (c *Client) GetHttpRequestRule(writer *GetHttpRequestRuleWriter) (*GetHttpRequestRuleOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetHttpRequestRule",
		Reader:             NewGetHttpRequestRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/http_request_rules/{index}",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetHttpRequestRuleOk:
		return value, nil
	case *GetHttpRequestRuleNotFound:
		return nil, errors.New(fmt.Sprintf("error while getting http request rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *GetHttpRequestRuleDefault:
		return nil, errors.New(fmt.Sprintf("error while getting http request rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting http request rule")
}

func (c *Client) GetHttpResponseRule(writer *GetHttpResponseRuleWriter) (*GetHttpResponseRuleOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetHttpResponseRule",
		Reader:             NewGetHttpResponseRuleReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/http_response_rules/{index}",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetHttpResponseRuleOk:
		return value, nil
	case *GetHttpResponseRuleNotFound:
		return nil, errors.New(fmt.Sprintf("error while getting http response rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *GetHttpResponseDefault:
		return nil, errors.New(fmt.Sprintf("error while getting http response rule code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting http response rule")
}

func (c *Client) GetHttpRequestRules(writer *GetHttpRequestRulesWriter) (*GetHttpRequestRulesOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetHttpRequestRules",
		Reader:             NewGetHttpRequestRulesReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/http_request_rules",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetHttpRequestRulesOk:
		return value, nil
	case *GetHttpRequestRulesDefault:
		return nil, errors.New(fmt.Sprintf("error while getting http request rules code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting http request rules")
}

func (c *Client) GetHttpResponseRules(writer *GetHttpResponseRulesWriter) (*GetHttpResponseRulesOk, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetHttpResponseRules",
		Reader:             NewGetHttpResponseRulesReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/http_response_rules",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *GetHttpResponseRulesOk:
		return value, nil
	case *GetHttpResponseRulesDefault:
		return nil, errors.New(fmt.Sprintf("error while getting http response rules code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while getting http response rules")
}
