package backendrule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateBackendSwitchingRuleCreated struct {
	Payload *models.BackendSwitchingRule
}

type CreateBackendSwitchingRuleAccepted struct {
	ReloadID string
	Payload  *models.BackendSwitchingRule
}

type CreateBackendSwitchingRuleBadRequest struct {
	Payload *models.Error
}

type CreateBackendSwitchingRuleConflict struct {
	Payload *models.Error
}

type CreateBackendSwitchingRuleDefault struct {
	StatusCode int
	Payload    *models.Error
}

type CreateBackendSwitchingRuleReader struct {
}

func NewCreateBackendSwitchingRuleReader() *CreateBackendSwitchingRuleReader {
	return &CreateBackendSwitchingRuleReader{}
}

func (r *CreateBackendSwitchingRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		backendSwitchingRule := &models.BackendSwitchingRule{}
		err := json.Unmarshal(response.Body(), backendSwitchingRule)
		if err != nil {
			return nil, err
		}
		return &CreateBackendSwitchingRuleCreated{Payload: backendSwitchingRule}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		backendSwitchingRule := &models.BackendSwitchingRule{}
		err := json.Unmarshal(response.Body(), backendSwitchingRule)
		if err != nil {
			return nil, err
		}
		return &CreateBackendSwitchingRuleAccepted{ReloadID: reloadID, Payload: backendSwitchingRule}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateBackendSwitchingRuleBadRequest{Payload: error}, nil
	case 409:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateBackendSwitchingRuleConflict{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateBackendSwitchingRuleDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
