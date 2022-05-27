package logs

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetLogTargetsOk struct {
	Payload *GetLogTargetsOkBody
}

type GetLogTargetsOkBody struct {
	Version int64              `json:"_version,omitempty"`
	Data    *models.LogTargets `json:"data,omitempty"`
}

type GetLogTargetsDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetLogTargetsReader struct {
}

func NewGetLogTargetsReader() *GetLogTargetsReader {
	return &GetLogTargetsReader{}
}

func (r *GetLogTargetsReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		logTargets := &GetLogTargetsOkBody{}
		err := json.Unmarshal(response.Body(), logTargets)
		if err != nil {
			return nil, err
		}
		return &GetLogTargetsOk{Payload: logTargets}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetLogTargetsDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
