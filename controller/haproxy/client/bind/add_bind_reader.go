package bind

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateBindCreated struct {
	Payload *models.Bind
}

type CreateBindAccepted struct {
	ReloadID string
	Payload  *models.Bind
}

type CreateBindBadRequest struct {
	Payload *models.Error
}

type CreateBindConflict struct {
	Payload *models.Error
}

type CreateBindDefault struct {
	StatusCode int
	Payload    *models.Error
}

type CreateBindReader struct {
}

func NewCreateBindReader() *CreateBindReader {
	return &CreateBindReader{}
}

func (r *CreateBindReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		bind := &models.Bind{}
		err := json.Unmarshal(response.Body(), bind)
		if err != nil {
			return nil, err
		}
		return &CreateBindCreated{Payload: bind}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		bind := &models.Bind{}
		err := json.Unmarshal(response.Body(), bind)
		if err != nil {
			return nil, err
		}
		return &CreateBindAccepted{Payload: bind, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateBindBadRequest{Payload: error}, nil
	case 409:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateBindConflict{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateBindDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
