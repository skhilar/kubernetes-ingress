package transactions

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type DeleteTransactionNoContent struct {
}

type DeleteTransactionNotFound struct {
	Payload *models.Error
}

type DeleteTransactionDefault struct {
	StatusCode int
	Payload    *models.Error
}

type DeleteTransactionReader struct {
}

func NewDeleteTransactionReader() *DeleteTransactionReader {
	return &DeleteTransactionReader{}
}

func (r *DeleteTransactionReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 204:
		return &DeleteTransactionNoContent{}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteTransactionNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteTransactionDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
