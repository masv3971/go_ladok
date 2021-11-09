package ladoktypes

// KataloginformationAnvandareAutentiserad is ladok response from /kataloginformation/anvandare/autentiserad
type KataloginformationAnvandareAutentiserad struct {
	Anvandarnamn   string `json:"Anvandarnamn"`
	Efternamen     string `json:"Efternamn"`
	Fornamn        string `json:"Fornamn"`
	SenastAndradAv string `json:"SenastAndradAv"`
	SenastSparad   string `json:"SenastSparad"`
	LarosateID     int    `json:"LarosateID"`
	UID            string `json:"Uid"`
	Link           []Link `json:"link"`
}

// KataloginformationAnvandarbehorighetEgna is ladok response from kataloginformation/anvandarbehorighet/egna
type KataloginformationAnvandarbehorighetEgna struct {
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

// Systemaktiviteter type
type Systemaktiviteter struct {
	Betafunktion      bool          `json:"Betafunktion"`
	I18NNyckel        string        `json:"I18nNyckel"`
	ID                int64         `json:"Id"`
	KlarForProduktion bool          `json:"KlarForProduktion"`
	Rattighetsniva    string        `json:"Rattighetsniva"`
	Link              []interface{} `json:"link"`
}

// KataloginformationBehorighetsprofil type
type KataloginformationBehorighetsprofil struct {
	Benamning struct {
		Sv string `json:"sv"`
		En string `json:"en"`
	} `json:"Benamning"`
	Dataavgransningar struct {
		Lista []interface{} `json:"Lista"`
		Link  []interface{} `json:"link"`
	} `json:"Dataavgransningar"`
	LarosateID        int                 `json:"LarosateID"`
	Rattighetsniva    string              `json:"Rattighetsniva"`
	Systemaktiviteter []Systemaktiviteter `json:"Systemaktiviteter"`
	UID               string              `json:"Uid"`
	Link              []Link              `json:"link"`
}
