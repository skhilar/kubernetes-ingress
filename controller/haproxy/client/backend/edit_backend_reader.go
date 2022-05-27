package backend

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditBackend struct {
	PayLoad *models.Backend
}

type EditBackendAccepted struct {
	ReloadID string
	Payload  *models.Backend
}

type EditBackendBadRequest struct {
	Payload *models.Error
}

type EditBackendNotFound struct {
	Payload *models.Error
}

type EditBackendDefault struct {
	StatusCode int
	Payload    *models.Error
}

type EditBackendReader struct {
}

func NewEditBackendReader() *EditBackendReader {
	return &EditBackendReader{}
}

func (r *EditBackendReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		backend := &models.Backend{}
		err := json.Unmarshal(response.Body(), backend)
		if err != nil {
			return nil, err
		}
		return &EditBackend{PayLoad: backend}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		backend := &models.Backend{}
		err := json.Unmarshal(response.Body(), backend)
		if err != nil {
			return nil, err
		}
		return &EditBackendAccepted{ReloadID: reloadID, Payload: backend}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditBackendBadRequest{Payload: error}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditBackendNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditBackendDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
