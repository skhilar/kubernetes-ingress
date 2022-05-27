package bind

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditBindOk struct {
	Payload *models.Bind
}

type EditBindAccepted struct {
	ReloadID string
	Payload  *models.Bind
}

type EditBindBadRequest struct {
	Payload *models.Error
}

type EditBindNotFound struct {
	Payload *models.Error
}

type EditBindDefault struct {
	StatusCode int
	Payload    *models.Error
}
type EditBindReader struct {
}

func NewEditBindReader() *EditBindReader {
	return &EditBindReader{}
}

func (r *EditBindReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		bind := &models.Bind{}
		err := json.Unmarshal(response.Body(), bind)
		if err != nil {
			return nil, err
		}
		return &EditBindOk{Payload: bind}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		bind := &models.Bind{}
		err := json.Unmarshal(response.Body(), bind)
		if err != nil {
			return nil, err
		}
		return &EditBindAccepted{Payload: bind, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditBindBadRequest{Payload: error}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditBindNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditBindDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
