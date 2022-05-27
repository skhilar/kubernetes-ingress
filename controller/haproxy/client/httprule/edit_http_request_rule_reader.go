package httprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditHttpRequestRuleOk struct {
	Payload *models.HTTPRequestRule
}

type EditHttpRequestRuleAccepted struct {
	Payload  *models.HTTPRequestRule
	ReloadID string
}

type EditHttpRequestRuleBadRequest struct {
	Payload *models.Error
}

type EditHttpRequestRuleNotFound struct {
	Payload *models.Error
}

type EditHttpRequestRuleDefault struct {
	StatusCode int
	Payload    *models.Error
}

type EditHttpRequestRuleReader struct {
}

func NewEditHttpRequestRuleReader() *EditHttpRequestRuleReader {
	return &EditHttpRequestRuleReader{}
}

func (r *EditHttpRequestRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		requestRule := &models.HTTPRequestRule{}
		err := json.Unmarshal(response.Body(), requestRule)
		if err != nil {
			return nil, err
		}
		return &EditHttpRequestRuleOk{Payload: requestRule}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		requestRule := &models.HTTPRequestRule{}
		err := json.Unmarshal(response.Body(), requestRule)
		if err != nil {
			return nil, err
		}
		return &EditHttpRequestRuleAccepted{Payload: requestRule, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditHttpRequestRuleBadRequest{Payload: error}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditHttpRequestRuleNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditHttpRequestRuleDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
