package configuration

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
	"strconv"
	"strings"
)

type GetConfigurationVersionReader struct {
}

func NewGetConfigurationVersionReader() *GetConfigurationVersionReader {
	return &GetConfigurationVersionReader{}
}

func (r *GetConfigurationVersionReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		version, err := strconv.ParseInt(strings.TrimSuffix(string(response.Body()), "\n"), 10, 64)
		if err != nil {
			return nil, err
		}
		return &GetConfigurationVersion{Version: version}, nil
	case 404:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetConfigurationVersionNotFound{Error: error}, errors.New("configuration not found")
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetConfigurationVersionDefault{StatusCode: response.StatusCode(), Error: error}, nil

	}
}

type GetConfigurationVersion struct {
	Version int64
}

type GetConfigurationVersionNotFound struct {
	Error *models.Error
}

type GetConfigurationVersionDefault struct {
	StatusCode int
	Error      *models.Error
}
