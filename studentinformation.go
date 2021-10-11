package goladokrest

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// StudentinformationService handles studentinformation
type StudentinformationService struct {
	client      *Client
	contentType string
}

// StudentReply is ladok reply from /studentinformation/student/{studentuid}
type StudentReply struct {
	Avliden                           bool      `json:"Avliden"`
	Fornamn                           string    `json:"Fornamn"`
	Efternamn                         string    `json:"Efternamn"`
	ExterntUID                        string    `json:"ExterntUID"`
	FelVidEtableringExternt           bool      `json:"FelVidEtableringExternt"`
	Fodelsedata                       time.Time `json:"Fodelsedata"`
	FolkbokforingsbevakningTillOchMed time.Time `json:"FolkbokforingsbevakningTillOchMed"`
	KonID                             int       `json:"KonID"`
	LarosateID                        int       `json:"LarosateID"`
	Personnummer                      string    `json:"Personnummer"`
	SenastAndradAv                    string    `json:"SenastAndradAv"`
	SenastSparad                      time.Time `json:"SenastSparad"`
	UID                               string    `json:"Uid"`
	UnikaIdentifierare                struct {
		LarosateID        int `json:"LarosateID"`
		UnikIdentifierare []struct {
			LarosateID     int       `json:"LarosateID"`
			SenastAndradAv string    `json:"SenastAndradAv"`
			SenastSparad   time.Time `json:"SenastSparad"`
			Typ            string    `json:"Typ"`
			UID            string    `json:"Uid"`
			Varde          string    `json:"Varde"`
			Link           Link      `json:"link"`
		} `json:"UnikIdentifierare"`
		Link Link `json:"link`
	} `json:"UnikaIdentifierare"`
	Link Link `json:"link`
}

// StudentCfg config for GetStudent
type StudentCfg struct {
	UID string
}

type StudentReply struct{}

// GetStudent return student
func (s *StudentinformationService) GetStudent(ctx context.Context, indata StudentCfg) (*StudentReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s/%s", "studentinformation/student", indata.UID),
		LadokAcceptHeader[s.contentType][s.client.Format],
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
