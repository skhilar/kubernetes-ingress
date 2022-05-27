package server

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditServerOk struct {
	Payload *models.Server
}

type EditServerAccepted struct {
	ReloadID string
	Payload  *models.Server
}

type EditServerBadRequest struct {
	Payload *models.Error
}

type CreateServerNotFound struct {
	Payload *models.Error
}

type EditServerDefault struct {
	StatusCode int
	Payload    *models.Error
}

type EditServerReader struct {
}

func NewEditServerReader() *EditServerReader {
	return &EditServerReader{}
}

func (r *EditServerReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		server := &models.Server{}
		err := json.Unmarshal(response.Body(), server)
		if err != nil {
			return nil, err
		}
		return &EditServerOk{Payload: server}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		server := &models.Server{}
		err := json.Unmarshal(response.Body(), server)
		if err != nil {
			return nil, err
		}
		return &EditServerAccepted{Payload: server, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditServerBadRequest{Payload: error}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateServerNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditServerDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
