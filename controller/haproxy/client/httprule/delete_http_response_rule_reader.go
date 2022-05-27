package httprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type DeleteHttpResponseRuleAccepted struct {
	ReloadID string
}

type DeleteHttpResponseRuleNoContent struct {
}

type DeleteHttpResponseRuleNotFound struct {
	Payload *models.Error
}

type DeleteHttpResponseRuleDefault struct {
	StatusCode int
	Payload    *models.Error
}

type DeleteHttpResponseRuleReader struct {
}

func NewDeleteHttpResponseRuleReader() *DeleteHttpResponseRuleReader {
	return &DeleteHttpResponseRuleReader{}
}

func (r *DeleteHttpResponseRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		return &DeleteHttpResponseRuleAccepted{ReloadID: reloadID}, nil
	case 204:
		return &DeleteHttpResponseRuleNoContent{}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteHttpResponseRuleNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteHttpResponseRuleDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
