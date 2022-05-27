package maps

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateMapFileCreated struct {
	Payload *models.MapEntry
}

type CreateMapFileBadRequest struct {
	Payload *models.Error
}

type CreateMapFileConflict struct {
	Payload *models.Error
}

type CreateMapFileDefault struct {
	Payload *models.Error
}

type CreateMapFileReader struct {
}

func NewCreateMapFileReader() *CreateMapFileReader {
	return &CreateMapFileReader{}
}

func (r *CreateMapFileReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		mapEntry := &models.MapEntry{}
		err := json.Unmarshal(response.Body(), mapEntry)
		if err != nil {
			return nil, err
		}
		return &CreateMapFileCreated{Payload: mapEntry}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateMapFileBadRequest{Payload: error}, nil
	case 409:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateMapFileConflict{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateMapFileDefault{Payload: error}, nil
	}
}
