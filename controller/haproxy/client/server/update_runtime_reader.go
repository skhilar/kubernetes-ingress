package server

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type UpdateRuntimeOk struct {
	Payload *models.RuntimeServer
}

type UpdateRuntimeBadRequest struct {
	Payload *models.Error
}

type UpdateRuntimeNotFound struct {
	Payload *models.Error
}

type UpdateRuntimeDefault struct {
	Payload *models.Error
}

type UpdateRuntimeReader struct {
}

func NewUpdateRuntimeReader() *UpdateRuntimeReader {
	return &UpdateRuntimeReader{}
}

func (r *UpdateRuntimeReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		runtimeServer := &models.RuntimeServer{}
		err := json.Unmarshal(response.Body(), runtimeServer)
		if err != nil {
			return nil, err
		}
		return &UpdateRuntimeOk{Payload: runtimeServer}, err
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &UpdateRuntimeBadRequest{Payload: error}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &UpdateRuntimeNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &UpdateRuntimeDefault{Payload: error}, nil

	}
}
