package backendrule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditBackendSwitchingRuleOK struct {
	Payload *models.BackendSwitchingRule
}

type EditBackendSwitchingRuleAccepted struct {
	ReloadID string
	Payload  *models.BackendSwitchingRule
}

type EditBackendSwitchingRuleBadRequest struct {
	Payload *models.Error
}

type EditBackendSwitchingRuleNotFound struct {
	Payload *models.Error
}

type EditBackendSwitchingRuleDefault struct {
	StatusCode int
	Payload    *models.Error
}

type EditBackendSwitchingRuleReader struct {
}

func NewEditBackendSwitchingRuleReader() *EditBackendSwitchingRuleReader {
	return &EditBackendSwitchingRuleReader{}
}

func (r *EditBackendSwitchingRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		backendRule := &models.BackendSwitchingRule{}
		err := json.Unmarshal(response.Body(), backendRule)
		if err != nil {
			return nil, err
		}
		return &EditBackendSwitchingRuleOK{Payload: backendRule}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		backendRule := &models.BackendSwitchingRule{}
		err := json.Unmarshal(response.Body(), backendRule)
		if err != nil {
			return nil, err
		}
		return &EditBackendSwitchingRuleAccepted{ReloadID: reloadID, Payload: backendRule}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditBackendSwitchingRuleBadRequest{Payload: error}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditBackendSwitchingRuleNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditBackendSwitchingRuleDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
