package maps

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type DeleteMapFileNoContent struct {
}

type DeleteMapFileNotFound struct {
	Payload *models.Error
}

type DeleteMapFileDefault struct {
	Payload *models.Error
}

type DeleteMapFileReader struct {
}

func NewDeleteMapFileReader() *DeleteMapFileReader {
	return &DeleteMapFileReader{}
}

func (r *DeleteMapFileReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 204:
		return &DeleteMapFileNoContent{}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteMapFileNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteMapFileDefault{Payload: error}, nil
	}

}
