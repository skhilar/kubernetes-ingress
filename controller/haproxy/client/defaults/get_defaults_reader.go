package defaults

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetDefaultsOk struct {
	Payload *GetDefaultsOKBody
}

type GetDefaultsOKBody struct {
	Version int64            `json:"_version,omitempty"`
	Data    *models.Defaults `json:"data,omitempty"`
}

type GetDefaultsDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetDefaultsReader struct {
}

func NewGetDefaultsReader() *GetDefaultsReader {
	return &GetDefaultsReader{}
}

func (r *GetDefaultsReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		defaults := &GetDefaultsOKBody{}
		err := json.Unmarshal(response.Body(), defaults)
		if err != nil {
			return nil, err
		}
		return &GetDefaultsOk{Payload: defaults}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetDefaultsDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
