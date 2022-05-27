package backendrule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetBackendSwitchingRulesOK struct {
	Payload *GetBackendSwitchingRulesOKBody
}

type GetBackendSwitchingRulesOKBody struct {
	Version int64                        `json:"_version,omitempty"`
	Data    models.BackendSwitchingRules `json:"data"`
}

type GetBackendSwitchingRulesDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetBackendSwitchingRulesReader struct {
}

func NewGetBackendSwitchingRulesReader() *GetBackendSwitchingRulesReader {
	return &GetBackendSwitchingRulesReader{}
}

func (r *GetBackendSwitchingRulesReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		payload := &GetBackendSwitchingRulesOKBody{}
		err := json.Unmarshal(response.Body(), payload)
		if err != nil {
			return nil, err
		}
		return &GetBackendSwitchingRulesOK{Payload: payload}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetBackendSwitchingRulesDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
