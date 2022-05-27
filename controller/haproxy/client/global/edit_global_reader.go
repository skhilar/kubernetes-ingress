package global

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditGlobalOk struct {
	Payload *models.Global
}

type EditGlobalAccepted struct {
	ReloadID string
	Payload  *models.Global
}

type EditGlobalBadRequest struct {
	Payload *models.Error
}

type EditGlobalDefault struct {
	StatusCode int
	Payload    *models.Error
}

type EditGlobalReader struct {
}

func NewEditGlobalReader() *EditGlobalReader {
	return &EditGlobalReader{}
}

func (r *EditGlobalReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		globals := &models.Global{}
		err := json.Unmarshal(response.Body(), globals)
		if err != nil {
			return nil, err
		}
		return &EditGlobalOk{Payload: globals}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		globals := &models.Global{}
		err := json.Unmarshal(response.Body(), globals)
		if err != nil {
			return nil, err
		}
		return &EditGlobalAccepted{Payload: globals, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditGlobalBadRequest{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditGlobalDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
