package backendrule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetBackendSwitchingRuleOK struct {
	Payload *GetBackendSwitchingRuleOKBody
}

type GetBackendSwitchingRuleNotFound struct {
	Payload *models.Error
}

type GetBackendSwitchingRuleDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetBackendSwitchingRuleOKBody struct {
	Version int64                        `json:"_version,omitempty"`
	Data    *models.BackendSwitchingRule `json:"data,omitempty"`
}

type GetBackendSwitchingRuleReader struct {
}

func NewGetBackendSwitchingRuleReader() *GetBackendSwitchingRuleReader {
	return &GetBackendSwitchingRuleReader{}
}

func (r *GetBackendSwitchingRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		payload := &GetBackendSwitchingRuleOKBody{}
		err := json.Unmarshal(response.Body(), payload)
		if err != nil {
			return nil, err
		}
		return &GetBackendSwitchingRuleOK{Payload: payload}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetBackendSwitchingRuleNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetBackendSwitchingRuleDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
