package storage

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateCertificateCreated struct {
	Payload *models.SslCertificate
}

type CreateCertificateBadRequest struct {
	Payload *models.Error
}

type CreateCertificateConflict struct {
	Payload *models.Error
}

type CreateCertificateDefault struct {
	Payload *models.Error
}

type CreateCertificateReader struct {
}

func NewCreateCertificateReader() *CreateCertificateReader {
	return &CreateCertificateReader{}
}

func (r *CreateCertificateReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		cert := &models.SslCertificate{}
		err := json.Unmarshal(response.Body(), cert)
		if err != nil {
			return nil, err
		}
		return &CreateCertificateCreated{Payload: cert}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateCertificateBadRequest{Payload: error}, nil
	case 409:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateCertificateConflict{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateCertificateDefault{Payload: error}, nil
	}
}
