package bind

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetBindOk struct {
	Payload *GetBindOkBody
}

type GetBindNotFound struct {
	Payload *models.Error
}

type GetBindOkBody struct {
	Version int64        `json:"_version,omitempty"`
	Data    *models.Bind `json:"data,omitempty"`
}

type GetBindDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetBindReader struct {
}

func NewGetBindReader() *GetBindReader {
	return &GetBindReader{}
}

func (r *GetBindReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		bind := &GetBindOkBody{}
		err := json.Unmarshal(response.Body(), bind)
		if err != nil {
			return nil, err
		}
		return &GetBindOk{Payload: bind}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetBindNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetBindDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
