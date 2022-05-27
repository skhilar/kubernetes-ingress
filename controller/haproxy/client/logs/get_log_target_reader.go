package logs

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetLogTargetOk struct {
	Payload *GetLogTargetOkBody
}

type GetLogTargetOkBody struct {
	Version int64             `json:"_version,omitempty"`
	Data    *models.LogTarget `json:"data,omitempty"`
}

type GetLogTargetNotFound struct {
	Payload *models.Error
}

type GetLogTargetDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetLogTargetReader struct {
}

func NewGetLogTargetReader() *GetLogTargetReader {
	return &GetLogTargetReader{}
}

func (r *GetLogTargetReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		logTarget := &GetLogTargetOkBody{}
		err := json.Unmarshal(response.Body(), logTarget)
		if err != nil {
			return nil, err
		}
		return &GetLogTargetOk{Payload: logTarget}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetLogTargetNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetLogTargetDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
