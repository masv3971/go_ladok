package ladok3

import (
	"context"
	"fmt"
	"net/http"
)

// KataloginformationService handles kataloginformation
type KataloginformationService struct {
	client      *Client
	contentType string
}

// GetAnvandareAutentiseradReply is ladok response from /kataloginformation/anvandare/autentiserad
type GetAnvandareAutentiseradReply struct {
	Anvandarnamn   string `json:"Anvandarnamn"`
	Efternamen     string `json:"Efternamn"`
	Fornamn        string `json:"Fornamn"`
	SenastAndradAv string `json:"SenastAndradAv"`
	SenastSparad   string `json:"SenastSparad"`
	LarosateID     int    `json:"LarosateID"`
	UID            string `json:"Uid"`
	Link           []Link `json:"link"`
}

// GetAnvandareAutentiserad gets kataloginformation/anvandare/autentiserad
func (s *KataloginformationService) GetAnvandareAutentiserad(ctx context.Context) (*GetAnvandareAutentiseradReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s", "kataloginformation/anvandare/autentiserad"),
		ladokAcceptHeader[s.contentType][s.client.format],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GetAnvandareAutentiseradReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// GetBehorighetsprofilReply is ladok reply from kataloginformation/behorighetsprofil/{uid}
type GetBehorighetsprofilReply struct {
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

// GetBehorighetsprofilerCfg configuration for GetBehorighetsprofil
type GetBehorighetsprofilerCfg struct {
	UID string `validate:"required,uuid"`
}

// GetBehorighetsprofil return structure of rights for uid
func (s *KataloginformationService) GetBehorighetsprofil(ctx context.Context, cfg *GetBehorighetsprofilerCfg) (*GetBehorighetsprofilReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s/%s", "kataloginformation/behorighetsprofil", cfg.UID),
		ladokAcceptHeader[s.contentType][s.client.format],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GetBehorighetsprofilReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// GetAnvandarbehorighetEgnaReply is ladok response from kataloginformation/anvandarbehorighet/egna
type GetAnvandarbehorighetEgnaReply struct {
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

// GetAnvandarbehorighetEgna return structure of ladok permission
func (s *KataloginformationService) GetAnvandarbehorighetEgna(ctx context.Context) (*GetAnvandarbehorighetEgnaReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s", "kataloginformation/anvandarbehorighet/egna"),
		ladokAcceptHeader[s.contentType][s.client.format],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	reply := &GetAnvandarbehorighetEgnaReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}
