package server

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateServerCreated struct {
	Payload *models.Server
}

type CreateServerAccepted struct {
	ReloadID string
	Payload  *models.Server
}

type CreateServerBadRequest struct {
	Payload *models.Error
}

type CreateServerConflict struct {
	Payload *models.Error
}

type CreateServerDefault struct {
	StatusCode int
	Payload    *models.Error
}

type CreateServerReader struct {
}

func NewCreateServerReader() *CreateServerReader {
	return &CreateServerReader{}
}

func (r *CreateServerReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		server := &models.Server{}
		err := json.Unmarshal(response.Body(), server)
		if err != nil {
			return nil, err
		}
		return &CreateServerCreated{Payload: server}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		server := &models.Server{}
		err := json.Unmarshal(response.Body(), server)
		if err != nil {
			return nil, err
		}
		return &CreateServerAccepted{Payload: server, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		return &CreateServerBadRequest{Payload: error}, nil
	case 409:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		return &CreateServerConflict{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		return &CreateServerDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
