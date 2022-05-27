package frontend

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type DeleteFrontendAccepted struct {
	ReloadID string
}

type DeleteFrontendNoContent struct {
}

type DeleteFrontendNotFound struct {
	Payload *models.Error
}

type DeleteFrontendDefault struct {
	StatusCode int
	Payload    *models.Error
}

type DeleteFrontendReader struct {
}

func NewDeleteFrontendReader() *DeleteFrontendReader {
	return &DeleteFrontendReader{}
}

func (r *DeleteFrontendReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		return &DeleteFrontendAccepted{ReloadID: reloadID}, nil
	case 204:
		return &DeleteFrontendNoContent{}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteFrontendNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteFrontendDefault{Payload: error, StatusCode: response.StatusCode()}, nil

	}

}
