package studentinformation

// GetStudentReply is ladok reply from /studentinformation/student/{studentuid}
type GetStudentReply struct {
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

// Link is a general ladok link structure
type Link struct {
	Method    string `json:"method"`
	URI       string `json:"uri"`
	MediaType string `json:"mediaType"`
	Rel       string `json:"rel"`
}
