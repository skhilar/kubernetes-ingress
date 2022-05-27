package backend

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
	"strconv"
)

type GetBackend struct {
	ConfigurationVersion int64
	Payload              GetBackendPayload
}

type GetBackendNotFound struct {
	Payload models.Error
}

type GetBackendDefault struct {
	StatusCode int
	Payload    models.Error
}

type GetBackendPayload struct {
	Data models.Backend `json:"data,omitempty"`
}

type GetBackendReader struct {
}

func NewGetBackendReader() *GetBackendReader {
	return &GetBackendReader{}
}

func (r *GetBackendReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		version, err := strconv.ParseInt(response.Header().Get("Configuration-Version"), 10, 64)
		backend := GetBackendPayload{}
		err = json.Unmarshal(response.Body(), &backend)
		if err != nil {
			return nil, err
		}
		return &GetBackend{ConfigurationVersion: version, Payload: backend}, nil
	case 404:
		error := models.Error{}
		err := json.Unmarshal(response.Body(), &error)
		if err != nil {
			return nil, err
		}
		return &GetBackendNotFound{Payload: error}, nil
	default:
		error := models.Error{}
		err := json.Unmarshal(response.Body(), &error)
		if err != nil {
			return nil, err
		}
		return &GetBackendDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
