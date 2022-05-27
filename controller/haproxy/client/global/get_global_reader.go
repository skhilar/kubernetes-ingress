package global

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
	"strconv"
)

type GetGlobalOk struct {
	ConfigurationVersion int64
	Payload              *GetGlobalOKBody
}

type GetGlobalOKBody struct {
	Version int64          `json:"_version,omitempty"`
	Data    *models.Global `json:"data,omitempty"`
}

type GetGlobalDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetGlobalReader struct {
}

func NewGetGlobalReader() *GetGlobalReader {
	return &GetGlobalReader{}
}

func (r *GetGlobalReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		globals := &GetGlobalOKBody{}
		err := json.Unmarshal(response.Body(), globals)
		if err != nil {
			return nil, err
		}
		version, err := strconv.ParseInt(response.Header().Get("Configuration-Version"), 10, 64)
		if err != nil {
			return nil, err
		}
		return &GetGlobalOk{Payload: globals, ConfigurationVersion: version}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetGlobalDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
