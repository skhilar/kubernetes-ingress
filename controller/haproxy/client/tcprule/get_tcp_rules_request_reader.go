package tcprule

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/haproxytech/models"
)

type GetTCPRulesRequestOk struct {
	Payload *models.TCPRequestRules
}

type GetTCPRulesRequestDefault struct {
	Payload *models.Error
}

type GetTCPRulesRequestReader struct {
}

func NewGetTCPRulesRequestReader() *GetTCPRulesRequestReader {
	return &GetTCPRulesRequestReader{}
}

func (r *GetTCPRulesRequestReader) ReadResponse(response *resty.Response) (interface{}, error) {
	switch response.StatusCode() {
	case 200:
		tcpRules := &models.TCPRequestRules{}
		err := json.Unmarshal(response.Body(), tcpRules)
		if err != nil {
			return nil, err
		}
		return &GetTCPRulesRequestOk{Payload: tcpRules}, nil
	default:
		error := &models.Error{}
		err := json.Unmarshal(response.Body(), error)
		if err != nil {
			return nil, err
		}
		return &GetTCPRulesRequestDefault{Payload: error}, nil
	}
}
