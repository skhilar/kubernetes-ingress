package server

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetServerOk struct {
	Payload *GetServerOkBody
}

type GetServerOkBody struct {
	Version int64          `json:"_version,omitempty"`
	Data    *models.Server `json:"data,omitempty"`
}

type GetServerNotFound struct {
	Payload *models.Error
}

type GetServerDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetServerReader struct {
}

func NewGetServerReader() *GetServerReader {
	return &GetServerReader{}
}

func (r *GetServerReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		server := &GetServerOkBody{}
		err := json.Unmarshal(response.Body(), server)
		if err != nil {
			return nil, err
		}
		return &GetServerOk{Payload: server}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetServerNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetServerDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
