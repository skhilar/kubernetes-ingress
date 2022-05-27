package httprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetHttpResponseRulesOk struct {
	Payload *GetHttpResponseOkBody
}

type GetHttpResponseOkBody struct {
	Version int64                     `json:"_version,omitempty"`
	Data    *models.HTTPResponseRules `json:"data,omitempty"`
}

type GetHttpResponseRulesDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetHttpResponseRulesReader struct {
}

func NewGetHttpResponseRulesReader() *GetHttpResponseRulesReader {
	return &GetHttpResponseRulesReader{}
}

func (r *GetHttpResponseRulesReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		responseRule := &GetHttpResponseOkBody{}
		err := json.Unmarshal(response.Body(), responseRule)
		if err != nil {
			return nil, err
		}
		return &GetHttpResponseRulesOk{Payload: responseRule}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetHttpResponseRulesDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
