package logs

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type EditLogTargetCreated struct {
	Payload *models.LogTarget
}

type EditLogTargetAccepted struct {
	ReloadID string
	Payload  *models.LogTarget
}

type EditLogTargetBadRequest struct {
	Payload *models.Error
}

type EditLogTargetDefault struct {
	StatusCode int
	Payload    *models.Error
}

type EditLogTargetReader struct {
}

func NewEditLogTargetReader() *EditLogTargetReader {
	return &EditLogTargetReader{}
}

func (r *EditLogTargetReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		logTarget := &models.LogTarget{}
		err := json.Unmarshal(response.Body(), logTarget)
		if err != nil {
			return nil, err
		}
		return &EditLogTargetCreated{Payload: logTarget}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		logTarget := &models.LogTarget{}
		err := json.Unmarshal(response.Body(), logTarget)
		if err != nil {
			return nil, err
		}
		return &EditLogTargetAccepted{Payload: logTarget, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditLogTargetBadRequest{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &EditLogTargetDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
