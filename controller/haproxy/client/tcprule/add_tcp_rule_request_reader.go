package tcprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type CreateTCPRequestRuleCreated struct {
	Payload *models.TCPRequestRule
}

type CreateTCPRequestRuleAccepted struct {
	ReloadID string
	Payload  *models.TCPRequestRule
}

type CreateTCPRequestRuleConflict struct {
	Payload *models.Error
}

type CreateTCPRequestRuleDefault struct {
	Payload *models.Error
}

type CreateTCPRequestRuleBadRequest struct {
	Payload *models.Error
}

type CreateTCPRequestRuleReader struct {
}

func NewCreateTCPRequestRuleReader() *CreateTCPRequestRuleReader {
	return &CreateTCPRequestRuleReader{}
}

func (r *CreateTCPRequestRuleReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 201:
		tcpRequestRule := &models.TCPRequestRule{}
		err := json.Unmarshal(response.Body(), tcpRequestRule)
		if err != nil {
			return nil, err
		}
		return &CreateTCPRequestRuleCreated{Payload: tcpRequestRule}, nil
	case 202:
		reloadID := response.Header().Get("Reload-ID")
		tcpRequestRule := &models.TCPRequestRule{}
		err := json.Unmarshal(response.Body(), tcpRequestRule)
		if err != nil {
			return nil, err
		}
		return &CreateTCPRequestRuleAccepted{Payload: tcpRequestRule, ReloadID: reloadID}, nil
	case 400:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateTCPRequestRuleBadRequest{Payload: error}, nil
	case 409:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateTCPRequestRuleConflict{Payload: error}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &CreateTCPRequestRuleDefault{Payload: error}, nil
	}
}
