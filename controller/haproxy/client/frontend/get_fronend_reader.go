package frontend

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
	"strconv"
)

type GetFrontendOk struct {
	ConfigurationVersion int64
	Payload              *GetFrontendOkBody
}

type GetFrontendOkBody struct {
	Version int64            `json:"_version,omitempty"`
	Data    *models.Frontend `json:"data,omitempty"`
}

type GetFrontendNotFound struct {
	Payload *models.Error
}

type GetFrontendDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetFrontendReader struct {
}

func NewGetFrontendReader() *GetFrontendReader {
	return &GetFrontendReader{}
}

func (r *GetFrontendReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		frontend := &GetFrontendOkBody{}
		err := json.Unmarshal(response.Body(), frontend)
		if err != nil {
			return nil, err
		}
		version, err := strconv.ParseInt(response.Header().Get("Configuration-Version"), 10, 64)
		if err != nil {
			return nil, err
		}
		return &GetFrontendOk{ConfigurationVersion: version, Payload: frontend}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetFrontendNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetFrontendDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
