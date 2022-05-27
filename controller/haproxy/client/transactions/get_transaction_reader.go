package transactions

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type Transaction struct {
	PayLoad *models.Transaction
}

type TransactionNotFound struct {
	Payload *models.Error
}

type TransactionDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetTransactionReader struct {
}

func NewGetTransactionReader() *GetTransactionReader {
	return &GetTransactionReader{}
}

func (r *GetTransactionReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		transaction := &models.Transaction{}
		err := json.Unmarshal(response.Body(), transaction)
		if err != nil {
			return nil, err
		}
		return &Transaction{PayLoad: transaction}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &TransactionNotFound{Payload: error}, nil

	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &TransactionDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
