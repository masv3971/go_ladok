package goladok3

import (
	"context"
	"fmt"
	"net/http"
)

// kataloginformationService handles kataloginformation
type kataloginformationService struct {
	client  *Client
	service string
}

func (s *kataloginformationService) acceptHeader() string {
	return ladokAcceptHeader[s.service][s.client.format]
}

// AnvandareAutentiserad is ladok response from /kataloginformation/anvandare/autentiserad
type AnvandareAutentiserad struct {
	Anvandarnamn   string `json:"Anvandarnamn"`
	Efternamen     string `json:"Efternamn"`
	Fornamn        string `json:"Fornamn"`
	SenastAndradAv string `json:"SenastAndradAv"`
	SenastSparad   string `json:"SenastSparad"`
	LarosateID     int    `json:"LarosateID"`
	UID            string `json:"Uid"`
	Link           []Link `json:"link"`
}

// Behorighetsprofil is ladok reply from kataloginformation/behorighetsprofil/{uid}
type Behorighetsprofil struct {
	Behorighetsprofiler []struct {
		Benamning struct {
			Sv string `json:"sv"`
		} `json:"Benamning"`
		Dataavgransningar struct {
			LarosateID int `json:"LarosateID"`
			Lista      []struct {
				DataDimension  string `json:"DataDimension"`
				DataID         string `json:"DataId"`
				LarosateID     int    `json:"LarosateID"`
				SenastAndradAv string `json:"SenastAndradAv"`
				SenastSparad   string `json:"SenastSparad"`
				UID            string `json:"Uid"`
				Link           []Link `json:"link"`
			} `json:"Lista"`
			SenastAndradAv string `json:"SenastAndradAv"`
			SenastSparad   string `json:"SenastSparad"`
			UID            string `json:"Uid"`
			Link           []Link `json:"link"`
		} `json:"Dataavgransningar"`
		LarosateID        int    `json:"LarosateID"`
		Rattighetsniva    string `json:"Rattighetsniva"`
		SenastAndradAv    string `json:"SenastAndradAv"`
		SenastSparad      string `json:"SenastSparad"`
		Systemaktiviteter []struct {
			Betafunktion      bool   `json:"Betafunktion"`
			I18NNyckel        string `json:"I18nNyckel"`
			ID                int64  `json:"Id"`
			KlarForProduktion bool   `json:"KlarForProduktion"`
			LarosateID        int    `json:"LarosateID"`
			Rattighetsniva    string `json:"Rattighetsniva"`
			Link              []Link `json:"link"`
		} `json:"Systemaktiviteter"`
		UID  string `json:"Uid"`
		Link []Link `json:"link"`
	} `json:"Behorighetsprofiler"`
	LarosateID     int    `json:"LarosateID"`
	SenastAndradAv string `json:"SenastAndradAv"`
	SenastSparad   string `json:"SenastSparad"`
	UID            string `json:"Uid"`
	Link           []Link `json:"link"`
}

// AnvandarbehorighetEgna is ladok response from kataloginformation/anvandarbehorighet/egna
type AnvandarbehorighetEgna struct {
	Anvandarbehorighet []struct {
		AnvandareRef struct {
			Anvandarnamn string `json:"Anvandarnamn"`
			Efternamn    string `json:"Efternamn"`
			Fornamn      string `json:"Fornamn"`
			UID          string `json:"Uid"`
			Link         Link   `json:"link"`
		} `json:"AnvandareRef"`
		BehorighetsprofilRef struct {
			Benamning []Benamning `json:"Benamning"`
			UID       string      `json:"Uid"`
			Link      Link        `json:"link"`
		} `json:"BehorighetsprofilRef"`
		BestalldTidpunkt string `json:"BestalldTidpunkt"`
		LarosateID       int    `json:"LarosateID"`
		OrganisationRef  struct {
			Benamning []Benamning `json:"Benamning"`
			UID       string      `json:"Uid"`
			Link      Link        `json:"link"`
		} `json:"OrganisationRef"`
		SenastAndradAv string `json:"SenastAndradAv"`
		SenastSparad   string `json:"SenastSparad"`
		Status         string `json:"Status"`
		UID            string `json:"Uid"`
		Link           []Link `json:"link"`
	} `json:"Anvandarbehorighet"`
	LarosateID     int    `json:"LarosateID"`
	SenastAndradAv string `json:"SenastAndradAv"`
	SenastSparad   string `json:"SenastSparad"`
	UID            string `json:"Uid"`
	Link           []Link `json:"link"`
}

// GetAnvandareAutentiserad gets kataloginformation/anvandare/autentiserad
func (s *kataloginformationService) GetAnvandareAutentiserad(ctx context.Context) (*AnvandareAutentiserad, *http.Response, error) {
	url := fmt.Sprintf("%s/%s", s.service, "anvandare/autentiserad")
	reply := &AnvandareAutentiserad{}
	resp, err := s.client.call(ctx, s.acceptHeader(), http.MethodGet, url, "", nil, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// GetBehorighetsprofilerCfg configuration for GetBehorighetsprofil
type GetBehorighetsprofilerCfg struct {
	UID string `validate:"required,uuid"`
}

// GetBehorighetsprofil return structure of rights for uid
func (s *kataloginformationService) GetBehorighetsprofil(ctx context.Context, req *GetBehorighetsprofilerCfg) (*Behorighetsprofil, *http.Response, error) {
	url := fmt.Sprintf("%s/%s", s.service, "behorighetsprofil")
	reply := &Behorighetsprofil{}
	resp, err := s.client.call(ctx, s.acceptHeader(), http.MethodGet, url, req.UID, nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// GetAnvandarbehorighetEgna return structure of ladok permission
func (s *kataloginformationService) GetAnvandarbehorighetEgna(ctx context.Context) (*AnvandarbehorighetEgna, *http.Response, error) {
	url := fmt.Sprintf("%s/%s", s.service, "anvandarbehorighet/egna")
	reply := &AnvandarbehorighetEgna{}
	resp, err := s.client.call(ctx, s.acceptHeader(), http.MethodGet, url, "", nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}
