package backendrule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type DeleteBackendSwitchingRuleAccepted struct {
	ReloadID string
}

type DeleteBackendSwitchingRuleNoContent struct {
}

type DeleteBackendSwitchingRuleNotFound struct {
	Payload *models.Error
}

type DeleteBackendSwitchingRuleDefault struct {
	StatusCode int
	Payload    *models.Error
}

type DeleteBackendSwitchingRuleReader struct {
}

func NewDeleteBackendSwitchingRuleReader() *DeleteBackendSwitchingRuleReader {
	return &DeleteBackendSwitchingRuleReader{}
}

func (r *DeleteBackendSwitchingRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		return &DeleteBackendSwitchingRuleAccepted{ReloadID: reloadID}, nil
	case 204:
		return &DeleteBackendSwitchingRuleNoContent{}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteBackendSwitchingRuleNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteBackendSwitchingRuleDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
