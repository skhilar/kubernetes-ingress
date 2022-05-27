package httprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditHttpResponseRuleOk struct {
	Payload *models.HTTPResponseRule
}

type EditHttpResponseRuleAccepted struct {
	Payload  *models.HTTPResponseRule
	ReloadID string
}

type EditHttpResponseRuleBadRequest struct {
	Payload *models.Error
}

type EditHttpResponseRuleNotFound struct {
	Payload *models.Error
}

type EditHttpResponseRuleDefault struct {
	StatusCode int
	Payload    *models.Error
}
type EditHttpResponseRuleReader struct {
}

func NewEditHttpResponseRuleReader() *EditHttpResponseRuleReader {
	return &EditHttpResponseRuleReader{}
}

func (r *EditHttpResponseRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		responseRule := &models.HTTPResponseRule{}
		err := json.Unmarshal(response.Body(), responseRule)
		if err != nil {
			return nil, err
		}
		return &EditHttpResponseRuleOk{Payload: responseRule}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		responseRule := &models.HTTPResponseRule{}
		err := json.Unmarshal(response.Body(), responseRule)
		if err != nil {
			return nil, err
		}
		return &EditHttpResponseRuleAccepted{Payload: responseRule, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditHttpResponseRuleBadRequest{Payload: error}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditHttpResponseRuleNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditHttpResponseRuleDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
