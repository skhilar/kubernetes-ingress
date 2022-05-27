package maps

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetMapFileOk struct {
	Payload *models.Map
}

type GetMapFileNotFound struct {
	Payload *models.Error
}

type GetMapFileDefault struct {
	Payload *models.Error
}

type GetMapFileReader struct {
}

func NewGetMapFileReader() *GetMapFileReader {
	return &GetMapFileReader{}
}

func (r *GetMapFileReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		mapFile := &models.Map{}
		err := json.Unmarshal(response.Body(), mapFile)
		if err != nil {
			return nil, err
		}
		return &GetMapFileOk{Payload: mapFile}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetMapFileNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetMapFileDefault{Payload: error}, nil
	}
}
