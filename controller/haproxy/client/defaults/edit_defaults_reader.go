package defaults

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditDefaultsOk struct {
	Payload *models.Defaults
}

type EditDefaultsAccepted struct {
	ReloadID string
	Payload  *models.Defaults
}

type EditDefaultsBadRequest struct {
	Payload *models.Error
}

type EditDefaultsDefault struct {
	StatusCode int
	Payload    *models.Error
}

type EditDefaultsReader struct {
}

func NewEditDefaultsReader() *EditDefaultsReader {
	return &EditDefaultsReader{}
}

func (r *EditDefaultsReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		defaults := &models.Defaults{}
		err := json.Unmarshal(response.Body(), defaults)
		if err != nil {
			return nil, err
		}
		return &EditDefaultsOk{Payload: defaults}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		defaults := &models.Defaults{}
		err := json.Unmarshal(response.Body(), defaults)
		if err != nil {
			return nil, err
		}
		return &EditDefaultsAccepted{Payload: defaults, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditDefaultsBadRequest{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditDefaultsDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
