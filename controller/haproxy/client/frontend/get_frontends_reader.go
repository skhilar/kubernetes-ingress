package frontend

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetFrontendsOk struct {
	Payload *GetFrontendsOkBody
}

type GetFrontendsOkBody struct {
	Version int64             `json:"_version,omitempty"`
	Data    *models.Frontends `json:"data,omitempty"`
}

type GetFrontendsDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetFrontendsReader struct {
}

func NewGetFrontendsReader() *GetFrontendsReader {
	return &GetFrontendsReader{}
}

func (r *GetFrontendsReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		frontends := &GetFrontendsOkBody{}
		err := json.Unmarshal(response.Body(), frontends)
		if err != nil {
			return nil, err
		}
		return &GetFrontendsOk{Payload: frontends}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetFrontendsDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
