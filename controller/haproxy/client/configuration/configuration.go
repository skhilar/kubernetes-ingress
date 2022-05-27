package configuration

import (
	"errors"
	"github.com/haproxytech/kubernetes-ingress/controller/haproxy/client"
)

const (
	TransactionID = "transaction_id"
)

type Client struct {
	client.HAProxyTransport
}

func NewClient(client client.HAProxyTransport) ClientService {
	return &Client{client}
}

type ClientService interface {
	GetConfigurationVersion(writer *GetConfigurationVersionWriter) (*GetConfigurationVersion, error)
}

func (c *Client) GetConfigurationVersion(writer *GetConfigurationVersionWriter) (*GetConfigurationVersion, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetConfigurationVersion",
		Reader:             NewGetConfigurationVersionReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/configuration/version",
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetConfigurationVersion)
	if ok {
		return success, nil
	}
	return nil, errors.New("not able to get configuration version")
}
