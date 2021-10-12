package goladokrest

import (
	"context"
	"fmt"
	"net/http"
)

// StudentinformationService handles studentinformation
type StudentinformationService struct {
	client      *Client
	contentType string
}

// StudentReply is ladok reply from /studentinformation/student/{studentuid}
type StudentReply struct {
	Avliden                           bool   `json:"Avliden"`
	Efternamn                         string `json:"Efternamn"`
	ExterntUID                        string `json:"ExterntUID"`
	FelVidEtableringExternt           bool   `json:"FelVidEtableringExternt"`
	Fodelsedata                       string `json:"Fodelsedata"`
	FolkbokforingsbevakningTillOchMed string `json:"FolkbokforingsbevakningTillOchMed"`
	Fornamn                           string `json:"Fornamn"`
	KonID                             int    `json:"KonID"`
	LarosateID                        int    `json:"LarosateID"`
	Personnummer                      string `json:"Personnummer"`
	SenastAndradAv                    string `json:"SenastAndradAv"`
	SenastSparad                      string `json:"SenastSparad"`
	UID                               string `json:"Uid"`
	UnikaIdentifierare                struct {
		LarosateID        int `json:"LarosateID"`
		UnikIdentifierare []struct {
			LarosateID     int    `json:"LarosateID"`
			SenastAndradAv string `json:"SenastAndradAv"`
			SenastSparad   string `json:"SenastSparad"`
			Typ            string `json:"Typ"`
			UID            string `json:"Uid"`
			Varde          string `json:"Varde"`
			Link           []Link `json:"link"`
		} `json:"UnikIdentifierare"`
		Link []Link `json:"link"`
	} `json:"UnikaIdentifierare"`
	Link []Link `json:"link"`
}

// StudentCfg config for GetStudent
type StudentCfg struct {
	UID string `validate:"required,uuid"`
}

// GetStudent return student
func (s *StudentinformationService) GetStudent(ctx context.Context, cfg *StudentCfg) (*StudentReply, *http.Response, error) {
	if err := validate(cfg); err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s/%s", "studentinformation/student", cfg.UID),
		LadokAcceptHeader[s.contentType][s.client.format],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &StudentReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
