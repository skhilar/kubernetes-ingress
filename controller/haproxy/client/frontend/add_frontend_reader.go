package frontend

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateFrontendCreated struct {
	Payload *models.Frontend
}

type CreateFrontendAccepted struct {
	Payload  *models.Frontend
	ReloadID string
}

type CreateFrontendBadRequest struct {
	Payload *models.Error
}

type CreateFrontendConflict struct {
	Payload *models.Error
}

type CreateFrontendDefault struct {
	StatusCode int
	Payload    *models.Error
}

type CreateFrontendReader struct {
}

func NewCreateFrontendReader() *CreateFrontendReader {
	return &CreateFrontendReader{}
}

func (r *CreateFrontendReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		frontend := &models.Frontend{}
		err := json.Unmarshal(response.Body(), frontend)
		if err != nil {
			return nil, err
		}
		return &CreateFrontendCreated{Payload: frontend}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		frontend := &models.Frontend{}
		err := json.Unmarshal(response.Body(), frontend)
		if err != nil {
			return nil, err
		}
		return &CreateFrontendAccepted{Payload: frontend, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateFrontendBadRequest{Payload: error}, nil
	case 409:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateFrontendConflict{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateFrontendDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
