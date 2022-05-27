package bind

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type DeleteBindAccepted struct {
	ReloadID string
}

type DeleteBindNoContent struct {
}

type DeleteBindNotFound struct {
	Payload *models.Error
}

type DeleteBindDefault struct {
	StatusCode int
	Payload    *models.Error
}

type DeleteBindReader struct {
}

func NewDeleteBindReader() *DeleteBindReader {
	return &DeleteBindReader{}
}

func (r *DeleteBindReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		return &DeleteBindAccepted{ReloadID: reloadID}, nil
	case 204:
		return &DeleteBindNoContent{}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteBindNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteBindDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}
}
