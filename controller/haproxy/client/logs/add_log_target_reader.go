package logs

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateLogTargetCreated struct {
	Payload *models.LogTarget
}

type CreateLogTargetAccepted struct {
	ReloadID string
	Payload  *models.LogTarget
}

type CreateLogTargetBadRequest struct {
	Payload *models.Error
}

type CreateLogTargetConflict struct {
	Payload *models.Error
}

type CreateLogTargetDefault struct {
	StatusCode int
	Payload    *models.Error
}

type CreateLogTargetReader struct {
}

func NewCreateLogTargetReader() *CreateLogTargetReader {
	return &CreateLogTargetReader{}
}

func (r *CreateLogTargetReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		logTarget := &models.LogTarget{}
		err := json.Unmarshal(response.Body(), logTarget)
		if err != nil {
			return nil, err
		}
		return &CreateLogTargetCreated{Payload: logTarget}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		logTarget := &models.LogTarget{}
		err := json.Unmarshal(response.Body(), logTarget)
		if err != nil {
			return nil, err
		}
		return &CreateLogTargetAccepted{Payload: logTarget, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateLogTargetBadRequest{Payload: error}, nil
	case 409:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateLogTargetConflict{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateLogTargetDefault{Payload: error, StatusCode: response.StatusCode()}, nil
	}

}
