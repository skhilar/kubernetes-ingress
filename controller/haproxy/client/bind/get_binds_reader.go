package bind

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetBindsOk struct {
	Payload *GetBindsOkBody
}

type GetBindsOkBody struct {
	Version int64         `json:"_version,omitempty"`
	Data    *models.Binds `json:"data,omitempty"`
}
type GetBindsDefault struct {
	StatusCode int
	Payload    *models.Error
}

type GetBindsReader struct {
}

func NewGetBindsReader() *GetBindsReader {
	return &GetBindsReader{}
}

func (r *GetBindsReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		binds := &GetBindsOkBody{}
		err := json.Unmarshal(response.Body(), binds)
		if err != nil {
			return nil, err
		}
		return &GetBindsOk{Payload: binds}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetBindsDefault{Payload: error, StatusCode: response.StatusCode()}, nil

	}

}
