package ladoktypes

import "encoding/xml"

// FeedID is the feeds id
type FeedID string

// Feed is the returning ladok type for atom feed
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Title   struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"title"`
	Link []struct {
		Text string `xml:",chardata"`
		Rel  string `xml:"rel,attr"`
		Type string `xml:"type,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	ID        FeedID `xml:"id"`
	Generator struct {
		Text string `xml:",chardata"`
		URI  string `xml:"uri,attr"`
	} `xml:"generator"`
	Updated string `xml:"updated"`
	Entry   []struct {
		Text     string `xml:",chardata"`
		Category struct {
			Text  string `xml:",chardata"`
			Term  string `xml:"term,attr"`
			Label string `xml:"label,attr"`
		} `xml:"category"`
		ID      string `xml:"id"`
		Updated string `xml:"updated"`
		Content struct {
			Text                             string                 `xml:",chardata"`
			Type                             string                 `xml:"type,attr"`
			AnvandareAndradEvent             *AnvandareEvent        `xml:"AnvandareAndradEvent,omitempty"`
			AnvandareSkapadEvent             *AnvandareEvent        `xml:"AnvandareSkapadEvent,omitempty"`
			KontaktuppgifterEvent            *KontaktuppgifterEvent `xml:"KontaktuppgifterEvent,omitempty"`
			ExternPartEvent                  *ExternPartEvent       `xml:"ExternPartEvent,omitempty"`
			LokalStudentEvent                *LokalStudentEvent     `xml:"LokalStudentEvent,omitempty"`
			ResultatPaModulAttesteratEvent   *ResultatEvent         `xml:"ResultatPaModulAttesteratEvent,omitempty"`
			ResultatPaHelKursAttesteratEvent *ResultatEvent         `xml:"ResultatPaHelKursAttesteratEvent,omitempty"`
		} `xml:"content"`
	} `xml:"entry"`
}

// AnvandareEvent ladok user event type
type AnvandareEvent struct {
	Text           string       `xml:",chardata"`
	Ki             string       `xml:"ki,attr"`
	Base           string       `xml:"base,attr"`
	Dap            string       `xml:"dap,attr"`
	Events         string       `xml:"events,attr"`
	HandelseUID    string       `xml:"HandelseUID"`
	EventContext   EventContext `xml:"EventContext"`
	AnvandareUID   string       `xml:"AnvandareUID"`
	Anvandarnamnet string       `xml:"Anvandarnamnet"`
	Efternamn      string       `xml:"Efternamn"`
	Fornamn        string       `xml:"Fornamn"`
	Email          string       `xml:"Email"`
}
type KontaktuppgifterEvent struct {
	Text         string       `xml:",chardata"`
	Si           string       `xml:"si,attr"`
	Base         string       `xml:"base,attr"`
	Dap          string       `xml:"dap,attr"`
	Events       string       `xml:"events,attr"`
	HandelseUID  string       `xml:"HandelseUID"`
	EventContext EventContext `xml:"EventContext"`
	Handelsetyp  string       `xml:"Handelsetyp"`
	Epostadress  string       `xml:"Epostadress"`
	Postadresser []struct {
		Text             string `xml:",chardata"`
		Land             string `xml:"Land"`
		PostadressTyp    string `xml:"PostadressTyp"`
		Postnummer       string `xml:"Postnummer"`
		Postort          string `xml:"Postort"`
		Utdelningsadress string `xml:"Utdelningsadress"`
		CareOf           string `xml:"CareOf"`
	} `xml:"Postadresser"`
	StudentUID    string `xml:"StudentUID"`
	Telefonnummer string `xml:"Telefonnummer"`
}
type LokalStudentEvent struct {
	Text         string `xml:",chardata"`
	Si           string `xml:"si,attr"`
	Base         string `xml:"base,attr"`
	Dap          string `xml:"dap,attr"`
	Events       string `xml:"events,attr"`
	HandelseUID  string `xml:"HandelseUID"`
	EventContext struct {
		Text         string `xml:",chardata"`
		AnvandareUID string `xml:"AnvandareUID"`
		Anvandarnamn string `xml:"Anvandarnamn"`
		LarosateID   string `xml:"LarosateID"`
	} `xml:"EventContext"`
	Handelsetyp       string `xml:"Handelsetyp"`
	Efternamn         string `xml:"Efternamn"`
	ExterntStudentUID string `xml:"ExterntStudentUID"`
	Fodelsedata       string `xml:"Fodelsedata"`
	Fornamn           string `xml:"Fornamn"`
	Kon               string `xml:"Kon"`
	Personnummer      string `xml:"Personnummer"`
	StudentUID        string `xml:"StudentUID"`
}

type ExternPartEvent struct {
	Text         string       `xml:",chardata"`
	Ki           string       `xml:"ki,attr"`
	Base         string       `xml:"base,attr"`
	Dap          string       `xml:"dap,attr"`
	Events       string       `xml:"events,attr"`
	HandelseUID  string       `xml:"HandelseUID"`
	EventContext EventContext `xml:"EventContext"`
	Benamningar  struct {
		Text      string `xml:",chardata"`
		Benamning []struct {
			Chardata string `xml:",chardata"`
			Sprakkod string `xml:"Sprakkod"`
			Text     string `xml:"Text"`
		} `xml:"Benamning"`
	} `xml:"Benamningar"`
	Beskrivningar struct {
		Text      string `xml:",chardata"`
		Benamning struct {
			Chardata string `xml:",chardata"`
			Sprakkod string `xml:"Sprakkod"`
			Text     string `xml:"Text"`
		} `xml:"Benamning"`
	} `xml:"Beskrivningar"`
	EventTyp          string `xml:"EventTyp"`
	Giltighetsperiod  string `xml:"Giltighetsperiod"`
	ID                string `xml:"Id"`
	Kod               string `xml:"Kod"`
	LandID            string `xml:"LandID"`
	TypAvExternPartID string `xml:"TypAvExternPartID"`
}
type ResultatEvent struct {
	Text         string `xml:",chardata"`
	Rr           string `xml:"rr,attr"`
	Base         string `xml:"base,attr"`
	Dap          string `xml:"dap,attr"`
	Events       string `xml:"events,attr"`
	HandelseUID  string `xml:"HandelseUID"`
	EventContext struct {
		Text         string `xml:",chardata"`
		AnvandareUID string `xml:"AnvandareUID"`
		Anvandarnamn string `xml:"Anvandarnamn"`
		LarosateID   string `xml:"LarosateID"`
	} `xml:"EventContext"`
	Beslut struct {
		Text              string `xml:",chardata"`
		BeslutUID         string `xml:"BeslutUID"`
		Beslutsdatum      string `xml:"Beslutsdatum"`
		Beslutsfattare    string `xml:"Beslutsfattare"`
		BeslutsfattareUID string `xml:"BeslutsfattareUID"`
	} `xml:"Beslut"`
	KursUID          string `xml:"KursUID"`
	KursinstansUID   string `xml:"KursinstansUID"`
	KurstillfalleUID string `xml:"KurstillfalleUID"`
	Resultat         struct {
		Text               string `xml:",chardata"`
		BetygsgradID       string `xml:"BetygsgradID"`
		BetygsskalaID      string `xml:"BetygsskalaID"`
		Examinationsdatum  string `xml:"Examinationsdatum"`
		GiltigSomSlutbetyg string `xml:"GiltigSomSlutbetyg"`
		OmfattningsPoang   string `xml:"OmfattningsPoang"`
		PrestationsPoang   string `xml:"PrestationsPoang"`
		ResultatUID        string `xml:"ResultatUID"`
	} `xml:"Resultat"`
	StudentUID            string `xml:"StudentUID"`
	UtbildningsinstansUID string `xml:"UtbildningsinstansUID"`
}

// EventContext is a common ladok type
type EventContext struct {
	Text         string `xml:",chardata"`
	AnvandareUID string `xml:"AnvandareUID"`
	Anvandarnamn string `xml:"Anvandarnamn"`
	LarosateID   string `xml:"LarosateID"`
}
