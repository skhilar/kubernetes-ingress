package frontend

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditFrontendOk struct {
	Payload *models.Frontend
}

type EditFrontendAccepted struct {
	ReloadID string
	Payload  *models.Frontend
}

type EditFrontendBadRequest struct {
	Payload *models.Error
}

type EditFrontendNotFound struct {
	Payload *models.Error
}

type EditFrontendDefault struct {
	StatusCode int
	Payload    *models.Error
}

type EditFrontendReader struct {
}

func NewEditFrontendReader() *EditFrontendReader {
	return &EditFrontendReader{}
}

func (r *EditFrontendReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		frontend := &models.Frontend{}
		err := json.Unmarshal(response.Body(), frontend)
		if err != nil {
			return nil, err
		}
		return &EditFrontendOk{Payload: frontend}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		frontend := &models.Frontend{}
		err := json.Unmarshal(response.Body(), frontend)
		if err != nil {
			return nil, err
		}
		return &EditFrontendAccepted{Payload: frontend, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditFrontendBadRequest{Payload: error}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditFrontendNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditFrontendDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
