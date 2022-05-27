package transactions

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type Transactions struct {
	Payload models.Transactions
}

type GetTransactionsDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetTransactionsReader struct {
}

func NewGetTransactionsReader() *GetTransactionsReader {
	return &GetTransactionsReader{}
}

func (r *GetTransactionsReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		transactions := &models.Transactions{}
		err := json.Unmarshal(response.Body(), transactions)
		if err != nil {
			return nil, err
		}
		return &Transactions{Payload: *transactions}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetTransactionsDefault{Payload: error, StatusCode: response.StatusCode()}, nil

	}

}
