package uppfoljning

import "encoding/xml"

// FeedRecent xx
type FeedRecent struct {
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
			Text                           string                          `xml:",chardata"`
			Type                           string                          `xml:"type,attr"`
			AnvandareAndradEvent           *AnvandareEvent                 `xml:"AnvandareAndradEvent,omitempty"`
			AnvandareSkapadEvent           *AnvandareEvent                 `xml:"AnvandareSkapadEvent,omitempty"`
			KontaktuppgifterEvent          *KontaktuppgifterEvent          `xml:"KontaktuppgifterEvent,omitempty"`
			ResultatPaModulAttesteratEvent *ResultatPaModulAttesteratEvent `xml:"ResultatPaModulAttesteratEvent,omitempty"`
			ExternPartEvent                *ExternPartEvent                `xml:"ExternPartEvent,omitempty"`
		} `xml:"content"`
	} `xml:"entry"`
}

// EventContext ladok atom type
type EventContext struct {
	Text         string `xml:",chardata"`
	AnvandareUID string `xml:"AnvandareUID"`
	Anvandarnamn string `xml:"Anvandarnamn"`
	LarosateID   string `xml:"LarosateID"`
}

// SuperFeed is a made up type, in order to make unstructured data to structured data.
type SuperFeed struct {
	ID          int           `json:"id,omitempty"`
	SuperEvents []*SuperEvent `json:"super_events,omitempty"`
}

// SuperEventContext is a made up type
type SuperEventContext struct {
	AnvandareUID string `xml:"AnvandareUID"`
	Anvandarnamn string `xml:"Anvandarnamn"`
	LarosateID   string `xml:"LarosateID"`
}

// SuperPostadresser is a made up type
type SuperPostadresser struct {
	Land             string `json:"land"`
	PostadressTyp    string `json:"postadress_typ"`
	Postnummer       string `json:"postnummer"`
	Postort          string `json:"postort"`
	Utdelningsadress string `json:"utdelningsadress"`
}

// SuperBeslut is a made up type
type SuperBeslut struct {
	BeslutUID         string `json:"beslut_uid"`
	Beslutsdatum      string `json:"beslutsdatum"`
	Beslutsfattare    string `json:"beslutsdattare"`
	BeslutsfattareUID string `json:"beslutsfattare_uid"`
}

// SuperResultat is a made up type
type SuperResultat struct {
	BetygsgradID       string `json:"betygsgrad_id"`
	BetygsskalaID      string `json:"betygsskala_id"`
	Examinationsdatum  string `json:"examinationsdatum"`
	GiltigSomSlutbetyg string `json:"giltig_som_slutbetyg"`
	OmfattningsPoang   string `json:"omfattnings_poang"`
	PrestationsPoang   string `json:"prestations_poang"`
	ResultatUID        string `json:"resultat_uid"`
}

// SuperEvent is a made up type consists of all the aviable ladok attributes
type SuperEvent struct {
	EventTypeName         string            `json:"event_type_name"`
	EventContext          SuperEventContext `json:"event_context"`
	HandelseUID           string            `json:"handelse_id"`
	ID                    string            `json:"id"`
	AnvandareUID          string            `json:"anvandare_uid"`
	Efternamn             string            `json:"efternamn"`
	Email                 string            `json:"email"`
	Fornamn               string            `json:"fornamn"`
	Handelsetyp           string            `json:"handelsetype"`
	Epostadress           string            `json:"epostaddress"`
	StudentUID            string            `json:"student_uid"`
	Postadresser          SuperPostadresser `json:"postadresser"`
	Telefonnummer         string            `json:"telefonnummer"`
	Beslut                SuperBeslut       `json:"beslut"`
	Resultat              SuperResultat     `json:"resultat"`
	UtbildningsinstansUID string            `json:"utbildningsinstans_uid"`
	Anvandarnamnet        string            `json:"anvandarnamnet"`
	EventTyp              string            `json:"event_type"`
	Giltighetsperiod      string            `json:"giltighetsperiod"`
	Kod                   string            `json:"kod"`
	LandID                string            `json:"land_id"`
	TypAvExternPartID     string            `json:"typ_av_extern_part_id"`
	KursUID               string            `json:"kurs_uid"`
	KursinstansUID        string            `json:"kursinstans_uid"`
	KurstillfalleUID      string            `json:"kurstillfalle_uid"`
}

// AnvandareEvent event
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
}

//KontaktuppgifterEvent event
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
	Postadresser struct {
		Text             string `xml:",chardata"`
		Land             string `xml:"Land"`
		PostadressTyp    string `xml:"PostadressTyp"`
		Postnummer       string `xml:"Postnummer"`
		Postort          string `xml:"Postort"`
		Utdelningsadress string `xml:"Utdelningsadress"`
	} `xml:"Postadresser"`
	StudentUID    string `xml:"StudentUID"`
	Telefonnummer string `xml:"Telefonnummer"`
}

// ResultatPaModulAttesteratEvent event
type ResultatPaModulAttesteratEvent struct {
	Text         string       `xml:",chardata"`
	Rr           string       `xml:"rr,attr"`
	Base         string       `xml:"base,attr"`
	Dap          string       `xml:"dap,attr"`
	Events       string       `xml:"events,attr"`
	HandelseUID  string       `xml:"HandelseUID"`
	EventContext EventContext `xml:"EventContext"`
	Beslut       struct {
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

//ExternPartEvent event
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
