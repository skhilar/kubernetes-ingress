package storage

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type DeleteCertificateAccepted struct {
	ReloadID string
}

type DeleteCertificateNoContent struct {
}

type DeleteCertificateNotFound struct {
	Payload *models.Error
}

type DeleteCertificateDefault struct {
	Payload *models.Error
}

type DeleteCertificateReader struct {
}

func NewDeleteCertificateReader() *DeleteCertificateReader {
	return &DeleteCertificateReader{}
}

func (r *DeleteCertificateReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		return &DeleteCertificateAccepted{ReloadID: reloadID}, nil
	case 204:
		return &DeleteCertificateNoContent{}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteCertificateNotFound{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &DeleteCertificateDefault{Payload: error}, nil
	}
}
