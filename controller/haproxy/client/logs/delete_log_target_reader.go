package logs

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type DeleteLogTargetAccepted struct {
	ReloadID string
}

type DeleteLogTargetNoContent struct {
}

type DeleteLogTargetNotFound struct {
	Payload *models.Error
}

type DeleteLogTargetDefault struct {
	StatusCode int
	Payload    *models.Error
}

type DeleteLogTargetReader struct {
}

func NewDeleteLogTargetReader() *DeleteLogTargetReader {
	return &DeleteLogTargetReader{}
}

func (r *DeleteLogTargetReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		return &DeleteLogTargetAccepted{ReloadID: reloadID}, nil
	case 204:
		return &DeleteLogTargetNoContent{}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteLogTargetNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteLogTargetDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
