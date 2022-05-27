package httprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetHttpRequestRuleOk struct {
	Payload *GetHttpRequestRuleOkBody
}

type GetHttpRequestRuleOkBody struct {
	Version int64                   `json:"_version,omitempty"`
	Data    *models.HTTPRequestRule `json:"data,omitempty"`
}

type GetHttpRequestRuleNotFound struct {
	Payload *models.Error
}

type GetHttpRequestRuleDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetHttpRequestRuleReader struct {
}

func NewGetHttpRequestRuleReader() *GetHttpRequestRuleReader {
	return &GetHttpRequestRuleReader{}
}

func (r *GetHttpRequestRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		requestRule := &GetHttpRequestRuleOkBody{}
		err := json.Unmarshal(response.Body(), requestRule)
		if err != nil {
			return nil, err
		}
		return &GetHttpRequestRuleOk{Payload: requestRule}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetHttpRequestRuleNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetHttpRequestRuleDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
