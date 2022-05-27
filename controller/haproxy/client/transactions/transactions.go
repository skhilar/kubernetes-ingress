package transactions

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
	CreateTransaction(writer *CreateTransactionWriter) (*TransactionCreated, error)
	GetTransactions(writer *GetTransactionsWriter) (*Transactions, error)
	GetTransaction(writer *GetTransactionWriter) (*Transaction, error)
	DeleteTransaction(writer *DeleteTransactionWriter) (*DeleteTransactionNoContent, error)
	CommitTransaction(writer *CommitTransactionWriter) (*CommitTransactionAccepted, *CommitTransaction, error)
}

func (c *Client) CreateTransaction(writer *CreateTransactionWriter) (*TransactionCreated, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CreateTransaction",
		Reader:             NewCreateTransactionReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "POST",
		PathPattern:        "services/haproxy/transactions",
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TransactionCreated)
	if ok {
		return success, nil
	}
	return nil, errors.New("not able to create transactions")
}

func (c *Client) GetTransactions(writer *GetTransactionsWriter) (*Transactions, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetTransactions",
		Reader:             NewGetTransactionsReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/transactions",
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Transactions)
	if ok {
		return success, nil
	}
	return nil, errors.New("not able to get transactions")
}

func (c *Client) GetTransaction(writer *GetTransactionWriter) (*Transaction, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "GetTransaction",
		Reader:             NewGetTransactionReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "GET",
		PathPattern:        "services/haproxy/transactions/{transaction_id}",
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Transaction)
	if ok {
		return success, nil
	}
	return nil, errors.New("not able to get transaction")
}

func (c *Client) DeleteTransaction(writer *DeleteTransactionWriter) (*DeleteTransactionNoContent, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "DeleteTransaction",
		Reader:             NewDeleteTransactionReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "DELETE",
		PathPattern:        "services/haproxy/transactions/{transaction_id}",
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTransactionNoContent)
	if ok {
		return success, nil
	}
	return nil, errors.New("not able to delete transaction")
}

func (c *Client) CommitTransaction(writer *CommitTransactionWriter) (*CommitTransactionAccepted, *CommitTransaction, error) {
	result, err := c.Execute(&client.HAProxyClientOperation{
		ID:                 "CommitTransaction",
		Reader:             NewCommitTransactionReader(),
		Writer:             writer,
		Context:            writer.Context,
		ConsumesMediaTypes: "application/json",
		ProducesMediaTypes: "application/json",
		Method:             "PUT",
		PathPattern:        "services/haproxy/transactions/{transaction_id}",
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CommitTransactionAccepted:
		return value, nil, nil
	case *CommitTransaction:
		return nil, value, nil
	case *CommitTransactionBadRequest:
		return nil, nil, errors.New(fmt.Sprintf("error while commiting transaction code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CommitTransactionNotFound:
		return nil, nil, errors.New(fmt.Sprintf("error while commiting transaction code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	case *CommitTransactionDefault:
		return nil, nil, errors.New(fmt.Sprintf("error while commiting transaction code=%d , message=%s", *value.Payload.Code, *value.Payload.Message))
	}
	return nil, nil, errors.New("unknown error while committing transaction")
}
