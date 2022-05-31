package storage

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
	CreateCertificate(writer *CreateCertificateWriter) (*CreateCertificateCreated, error)
	DeleteCertificate(writer *DeleteCertificateWriter) (*DeleteCertificateAccepted, *DeleteCertificateNoContent, error)
}

func (c *Client) CreateCertificate(writer *CreateCertificateWriter) (*CreateCertificateCreated, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CreateCertificate",
		Reader:             NewCreateCertificateReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "multipart/form-data",
		Method:             "POST",
		PathPattern:        "services/haproxy/storage/ssl_certificates",
	})
	if err != nil {
		return nil, err
	}
	switch value := result.(type) {
	case *CreateCertificateCreated:
		return value, nil
	case *CreateCertificateBadRequest:
		return nil, errors.New(fmt.Sprintf("error while creating certificate code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateCertificateConflict:
		return nil, errors.New(fmt.Sprintf("error while creating certificate code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CreateCertificateDefault:
		return nil, errors.New(fmt.Sprintf("error while creating certificate code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, errors.New("unknown error while creating certificate")
}

func (c *Client) DeleteCertificate(writer *DeleteCertificateWriter) (*DeleteCertificateAccepted, *DeleteCertificateNoContent, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "DeleteCertificate",
		Reader:             NewDeleteCertificateReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "DELETE",
		PathPattern:        "services/haproxy/storage/ssl_certificates/{name}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteCertificateAccepted:
		return value, nil, nil
	case *DeleteCertificateNoContent:
		return nil, value, nil
	case *DeleteCertificateNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting certificate code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *DeleteCertificateDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while deleting certificate code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while deleting certificate")
}
