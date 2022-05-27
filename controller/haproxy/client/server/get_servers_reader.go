package server

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetServersOk struct {
	Payload *GetServersOkBody
}

type GetServersOkBody struct {
	Version int64           `json:"_version,omitempty"`
	Data    *models.Servers `json:"data,omitempty"`
}

type GetServersDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetServersReader struct {
}

func NewGetServersReader() *GetServersReader {
	return &GetServersReader{}
}

func (r *GetServersReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		servers := &GetServersOkBody{}
		err := json.Unmarshal(response.Body(), servers)
		if err != nil {
			return nil, err
		}
		return &GetServersOk{Payload: servers}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetServersDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
