package goladok3

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
