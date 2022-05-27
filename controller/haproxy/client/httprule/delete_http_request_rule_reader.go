package httprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type DeleteHttpRequestRuleAccepted struct {
	ReloadID string
}

type DeleteHttpRequestRuleNoContent struct {
}

type DeleteHttpRequestRuleNotFound struct {
	Payload *models.Error
}

type DeleteHttpRequestRuleDefault struct {
	StatusCode int
	Payload    *models.Error
}

type DeleteHttpRequestRuleReader struct {
}

func NewDeleteHttpRequestRuleReader() *DeleteHttpRequestRuleReader {
	return &DeleteHttpRequestRuleReader{}
}

func (r *DeleteHttpRequestRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		return &DeleteHttpRequestRuleAccepted{ReloadID: reloadID}, nil
	case 204:
		return &DeleteHttpRequestRuleNoContent{}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteHttpRequestRuleNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteHttpRequestRuleDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
