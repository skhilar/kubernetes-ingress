package httprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateHttpRequestRuleCreated struct {
	Payload *models.HTTPRequestRule
}

type CreateHttpRequestRuleAccepted struct {
	ReloadID string
	Payload  *models.HTTPRequestRule
}

type CreateHttpRequestRuleBadRequest struct {
	Payload *models.Error
}

type CreateHttpRequestRuleConflict struct {
	Payload *models.Error
}

type CreateHttpRequestRuleDefault struct {
	StatusCode int
	Payload    *models.Error
}

type CreateHttpRequestRuleReader struct {
}

func NewCreateHttpRequestRuleReader() *CreateHttpRequestRuleReader {
	return &CreateHttpRequestRuleReader{}
}

func (r *CreateHttpRequestRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		requestRule := &models.HTTPRequestRule{}
		err := json.Unmarshal(response.Body(), requestRule)
		if err != nil {
			return nil, err
		}
		return &CreateHttpRequestRuleCreated{Payload: requestRule}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		requestRule := &models.HTTPRequestRule{}
		err := json.Unmarshal(response.Body(), requestRule)
		if err != nil {
			return nil, err
		}
		return &CreateHttpRequestRuleAccepted{Payload: requestRule, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateHttpRequestRuleBadRequest{Payload: error}, nil
	case 409:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateHttpRequestRuleConflict{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateHttpRequestRuleDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
