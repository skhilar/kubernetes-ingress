package httprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetHttpResponseRuleOk struct {
	Payload *GetHttpResponseRuleOkBody
}

type GetHttpResponseRuleOkBody struct {
	Version int64                    `json:"_version,omitempty"`
	Data    *models.HTTPResponseRule `json:"data,omitempty"`
}

type GetHttpResponseRuleNotFound struct {
	Payload *models.Error
}

type GetHttpResponseDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetHttpResponseRuleReader struct {
}

func NewGetHttpResponseRuleReader() *GetHttpResponseRuleReader {
	return &GetHttpResponseRuleReader{}
}

func (r *GetHttpResponseRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		responseRule := &GetHttpResponseRuleOkBody{}
		err := json.Unmarshal(response.Body(), responseRule)
		if err != nil {
			return nil, err
		}
		return &GetHttpResponseRuleOk{Payload: responseRule}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetHttpResponseRuleNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetHttpResponseDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
