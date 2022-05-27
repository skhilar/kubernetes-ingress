package backend

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type DeleteBackendAccepted struct {
	ReloadID string
}

type DeleteBackendNoContent struct {
}

type DeleteBackendNotFound struct {
	Payload *models.Error
}

type DeleteBackendDefault struct {
	StatusCode int
	Payload    *models.Error
}

type DeleteBackendReader struct {
}

func NewDeleteBackendReader() *DeleteBackendReader {
	return &DeleteBackendReader{}
}

func (r *DeleteBackendReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		return &DeleteBackendAccepted{ReloadID: reloadID}, nil
	case 204:
		return &DeleteBackendNoContent{}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteBackendNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteBackendDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
