package httprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetHttpRequestRulesOk struct {
	Payload *GetHttpRequestRulesOkBody
}

type GetHttpRequestRulesOkBody struct {
	Version int64                    `json:"_version,omitempty"`
	Data    *models.HTTPRequestRules `json:"data,omitempty"`
}

type GetHttpRequestRulesDefault struct {
	StatusCode int
	Payload    *models.Error
}
type GetHttpRequestRulesReader struct {
}

func NewGetHttpRequestRulesReader() *GetHttpRequestRulesReader {
	return &GetHttpRequestRulesReader{}
}

func (r *GetHttpRequestRulesReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		requestRule := &GetHttpRequestRulesOkBody{}
		err := json.Unmarshal(response.Body(), requestRule)
		if err != nil {
			return nil, err
		}
		return &GetHttpRequestRulesOk{Payload: requestRule}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetHttpRequestRulesDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
