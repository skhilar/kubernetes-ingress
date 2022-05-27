package httprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateHttpResponseRuleCreated struct {
	Payload *models.HTTPResponseRule
}

type CreateHttpResponseRuleAccepted struct {
	ReloadID string
	Payload  *models.HTTPResponseRule
}

type CreateHttpResponseRuleBadRequest struct {
	Payload *models.Error
}

type CreateHttpResponseRuleConflict struct {
	Payload *models.Error
}

type CreateHttpResponseRuleDefault struct {
	StatusCode int
	Payload    *models.Error
}

type CreateHttpResponseRuleReader struct {
}

func NewCreateHttpResponseRuleReader() *CreateHttpResponseRuleReader {
	return &CreateHttpResponseRuleReader{}
}

func (r *CreateHttpResponseRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		responseRule := &models.HTTPResponseRule{}
		err := json.Unmarshal(response.Body(), responseRule)
		if err != nil {
			return nil, err
		}
		return &CreateHttpResponseRuleCreated{Payload: responseRule}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		responseRule := &models.HTTPResponseRule{}
		err := json.Unmarshal(response.Body(), responseRule)
		if err != nil {
			return nil, err
		}
		return &CreateHttpResponseRuleAccepted{Payload: responseRule, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateHttpResponseRuleBadRequest{Payload: error}, nil
	case 409:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateHttpResponseRuleConflict{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateHttpResponseRuleDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
