package goladokrest

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// KataloginformationService handles kataloginformation
type KataloginformationService struct {
	client      *Client
	contentType string
}

// AnvandareAutentiseradReply is ladok response from /kataloginformation/anvandare/autentiserad
type AnvandareAutentiseradReply struct {
	Anvandarnamn   string    `json:"Anvandarnamn"`
	Efternamen     string    `json:"Efternamn"`
	Fornamn        string    `json:"Fornamn"`
	SenastAndradAv string    `json:"SenastAndradAv"`
	SenastSparad   time.Time `json:"SenastSparad"`
	LarosateID     int       `json:"LarosateID"`
	UID            string    `json:"Uid"`
	Link           []Link    `json:"link"`
}

// GetAnvandareAutentiserad gets kataloginformation/anvandare/autentiserad
func (s *KataloginformationService) GetAnvandareAutentiserad(ctx context.Context) (*AnvandareAutentiseradReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s", "kataloginformation/anvandare/autentiserad"),
		LadokAcceptHeader[s.contentType][s.client.Format],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &AnvandareAutentiseradReply{}
	resp, err := s.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// BehorighetsprofilReply is ladok reply from kataloginformation/behorighetsprofil/{uid}
type BehorighetsprofilReply struct {
	Behorighetsprofiler []struct {
		Benamning map[string]string `json:"Benaming"`
	} `json:"Behorighetsprofiler"`
	Dataavgransningar struct {
		LarosateID int `json:"LarosateID"`
		Lista      []struct {
			DataDimension  string    `json:"DataDimension"`
			DataID         string    `json:"DataId"`
			LarosateID     int       `json:"LarosateID"`
			SenastAndradAv string    `json:"SenastAndradAv"`
			SenastSparad   time.Time `json:"SenastSparad"`
			UID            string    `json:"Uid"`
			Link           []Link    `json:"link"`
		}
		SenastAndradAv string    `json"SenastAndradAv"`
		SenastSparad   time.Time `json:"SenastSparad"`
		UID            string    `json:"Uid"`
		Link           []Link    `json:"link"`
	}
	LarosateID        int       `json:"LarosateID"`
	Rattighetsniva    string    `json:"Rattighetsniva"`
	SenastAndradAv    string    `json:"SenastAndradAv"`
	SenastSparad      time.Time `json:"SenastSparad"`
	Systemaktiviteter []struct {
		Betafunktion      bool   `json:"Betafunktion"`
		I18nNyckel        string `json:"I18nNyckel"`
		ID                int    `json:"Id"`
		KlarForProduktion bool   `json:"KlarForProduktion"`
		LarosateID        int    `json:"LarosateID"`
		Rattighetsniva    string `json:"Rattighetsniva"`
		Link              []Link `json:"link"`
	}
	UID            string    `json:"Uid"`
	Link           []Link    `json:"link"`
	LarosateID     int       `json:"LarosateID"`
	SenastAndradAv string    `json:"SenastAndradAv"`
	SenastSparad   time.Time `json:"SenastSparad"`
	UID            string    `json:"Uid"`
	Link           []Link    `json:"links"`
}

// GetBehorighetsprofil return structure of rights for uid
func (s *KataloginformationService) GetBehorighetsprofil(ctx context.Context, uid string) (*BehorighetsprofilReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s/%s", "kataloginformation/behorighetsprofil", uid),
		LadokAcceptHeader[s.contentType][s.client.Format],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &BehorighetsprofilReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// AnvandarbehorighetEgnaReply is ladok response from kataloginformation/anvandarbehorighet/egna
type AnvandarbehorighetEgnaReply struct {
	Anvandarbehorighet []struct {
		AnvandareRef struct {
			Anvandarnamn string `json:"Anvandarnamn"`
			Efternamn    string `json:"Efternamn"`
			Fornamn      string `json:"Fornamn"`
			UID          string `json:"Uid"`
			Link         Link   `json:"link"`
		}
		BehorighetsprofilRef struct {
			Benamning []Benamning `json:"Benamning"`
			UID       string      `json:"Uid"`
			Link      Link        `json:"link"`
		}
	}
	BestalldTidpunkt time.Time `json:"BestalldTidpunkt"`
	LarosateID       int       `json:"LarosateID"`
	OrganisationRef  struct {
		Benamning []Benamning `json:"Benamning"`
		UID       string      `json:"Uid"`
		Link      Link        `json:"link"`
	}
	SenastAndradAv string    `json:"SenastAndradAv"`
	SenastSparad   time.Time `json:"SenastSparad"`
	Status         string    `json:"Status"`
	UID            string    `json:"Uid"`
	Link           []Link    `json:"link"`
}

// GetAnvandarbehorighetEgna return structure of ladok permission
func (s *KataloginformationService) GetAnvandarbehorighetEgna(ctx context.Context) (*AnvandarbehorighetEgnaReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s", "kataloginformation/anvandarbehorighet/egna"),
		LadokAcceptHeader[s.contentType][s.client.Format],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	reply := &AnvandarbehorighetEgnaReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}
