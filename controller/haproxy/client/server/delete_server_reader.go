package server

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type DeleteServerAccepted struct {
	ReloadID string
}

type DeleteServerNoContent struct {
}

type DeleteServerNotFound struct {
	Payload *models.Error
}

type DeleteServerDefault struct {
	StatusCode int
	Payload    *models.Error
}

type DeleteServerReader struct {
}

func NewDeleteServerReader() *DeleteServerReader {
	return &DeleteServerReader{}
}

func (r *DeleteServerReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		return &DeleteServerAccepted{ReloadID: reloadID}, nil
	case 204:
		return &DeleteServerNoContent{}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteServerNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteServerDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
