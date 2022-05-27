package transactions

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type TransactionCreated struct {
	Payload *models.Transaction
}

type TransactionsDefault struct {
	StatusCode int
	Payload    *models.Error
}

type TransactionTooManyRequests struct {
	Payload *models.Error
}

type CreateTransactionReader struct {
}

func NewCreateTransactionReader() *CreateTransactionReader {
	return &CreateTransactionReader{}
}

func (r *CreateTransactionReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		transactions := &models.Transaction{}
		err := json.Unmarshal(response.Body(), transactions)
		if err != nil {
			return nil, err
		}
		return &TransactionCreated{Payload: transactions}, nil
	case 429:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &TransactionTooManyRequests{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &TransactionsDefault{StatusCode: response.StatusCode(), Payload: error}, nil
	}
}
