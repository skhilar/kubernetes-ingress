package backend

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateBackendCreated struct {
	Payload *models.Backend
}

type CreateBackendAccepted struct {
	ReloadID string
	Payload  *models.Backend
}

type CreateBackendBadRequest struct {
	Payload *models.Error
}

type CreateBackendConflict struct {
	Payload *models.Error
}

type CreateBackendDefault struct {
	StatusCode int
	Payload    *models.Error
}

type CreateBackendReader struct {
}

func NewCreateBackendReader() *CreateBackendReader {
	return &CreateBackendReader{}
}

func (r *CreateBackendReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		backend := &models.Backend{}
		err := json.Unmarshal(response.Body(), backend)
		if err != nil {
			return nil, err
		}
		return &CreateBackendCreated{Payload: backend}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		backend := &models.Backend{}
		err := json.Unmarshal(response.Body(), backend)
		if err != nil {
			return nil, err
		}
		return &CreateBackendAccepted{Payload: backend, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateBackendBadRequest{Payload: error}, nil
	case 409:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateBackendConflict{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateBackendDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
