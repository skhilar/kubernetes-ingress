package transactions

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CommitTransaction struct {
	Payload *models.Transaction
}

type CommitTransactionAccepted struct {
	ReloadID string
	Payload  *models.Transaction
}

type CommitTransactionBadRequest struct {
	Payload *models.Error
}

type CommitTransactionNotFound struct {
	Payload *models.Error
}

type CommitTransactionNotHandled struct {
	Payload *models.Error
}

type CommitTransactionDefault struct {
	StatusCode int
	Payload    *models.Error
}

type CommitTransactionReader struct {
}

func NewCommitTransactionReader() *CommitTransactionReader {
	return &CommitTransactionReader{}
}

func (r *CommitTransactionReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		transaction := &models.Transaction{}
		err := json.Unmarshal(response.Body(), transaction)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return &CommitTransaction{Payload: transaction}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		transaction := &models.Transaction{}
		err := json.Unmarshal(response.Body(), transaction)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return &CommitTransactionAccepted{ReloadID: reloadID, Payload: transaction}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CommitTransactionBadRequest{Payload: error}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CommitTransactionNotFound{Payload: error}, nil
	case 406:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CommitTransactionNotHandled{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CommitTransactionDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
