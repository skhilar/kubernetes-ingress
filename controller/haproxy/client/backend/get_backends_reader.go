package backend

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetBackends struct {
	Payload GetBackendsPayload
}

type GetBackendsDefault struct {
	StatusCode int
	Payload    models.Error
}

type GetBackendsPayload struct {
	Data models.Backends `json:"data,omitempty"`
}

type GetBackendsReader struct {
}

func NewGetBackendsReader() *GetBackendsReader {
	return &GetBackendsReader{}
}

func (r *GetBackendsReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		backends := GetBackendsPayload{}
		fmt.Println(string(response.Body()))
		err := json.Unmarshal(response.Body(), &backends)
		if err != nil {
			return nil, err
		}
		return &GetBackends{Payload: backends}, nil
	default:
		error := models.Error{}
		err := json.Unmarshal(response.Body(), &error)
		if err != nil {
			return nil, err
		}
		return &GetBackendsDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
