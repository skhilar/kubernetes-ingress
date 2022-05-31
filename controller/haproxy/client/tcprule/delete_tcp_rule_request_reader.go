package tcprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type DeleteTCPRequestRuleAccepted struct {
	ReloadID string
}

type DeleteTCPRequestRuleNoContent struct {
}

type DeleteTCPRequestRuleNotFound struct {
	Payload *models.Error
}

type DeleteTCPRequestRuleDefault struct {
	Payload *models.Error
}

type DeleteTCPRequestRuleReader struct {
}

func NewDeleteTCPRequestRuleReader() *DeleteTCPRequestRuleReader {
	return &DeleteTCPRequestRuleReader{}
}

func (r *DeleteTCPRequestRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		return &DeleteTCPRequestRuleAccepted{ReloadID: reloadID}, nil
	case 204:
		return &DeleteTCPRequestRuleNoContent{}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteTCPRequestRuleNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteTCPRequestRuleDefault{Payload: error}, nil
	}
}
