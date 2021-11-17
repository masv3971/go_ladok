package ladokmocks

import (
	"encoding/json"

	"github.com/masv3971/goladok3/ladoktypes"
)

var (
	// StudentUID uid for testing student
	StudentUID = "339A47C0-426D-4012-B83A-6427E9587352"
	// ExterntUID uid for testing student
	ExterntUID = "72460B4B-8F15-442C-A464-0743BDFB1429"
	// Personnummer is a mock number from skatteverket
	Personnummer = "198602179882"
)

// JSONStudentinformationStudent mock ladok reply
var JSONStudentinformationStudent = []byte(`
{
	"Avliden": false,
	"Efternamn": "TestEfternamn",
	"ExterntUID": "11111111-2222-0000-0000-000000000000",
	"FelVidEtableringExternt": false,
	"Fodelsedata": "1996-11-05",
	"FolkbokforingsbevakningTillOchMed": "2020-03-13",
	"Fornamn": "TestFornamn",
	"KonID": 1,
	"LarosateID": 96,
	"Personnummer": "199611052383",
	"SenastAndradAv": "testEppn@ladok3.ladok.se",
	"SenastSparad": "2012-01-11T12:45:45",
	"Uid": "11111111-2222-0000-0000-000000000000"
}
`)

// JSONAktivPaLarosate mock ladok reply
var JSONAktivPaLarosate = []byte(` 
		{
			"Studentkopplingar": [
			  {
				"LarosateID": 27,
				"link": [],
				"studentUID": "339A47C0-426D-4012-B83A-6427E9587352"
			  },
			  {
				"LarosateID": 39,
				"link": [],
				"studentUID": "339A47C0-426D-4012-B83A-6427E9587352"
			  }
			],
			"link": []
		  }
		`)

// MockStudentinformationStudent return mock
func MockStudentinformationStudent() *ladoktypes.Student {
	s := &ladoktypes.Student{}
	json.Unmarshal(JSONStudentinformationStudent, s)
	return s
}

// StudentJSON return JSON object of a student
func StudentJSON(uid, externtUID, personnummer string) []byte {
	s := &ladoktypes.Student{}

	if err := json.Unmarshal(JSONStudentinformationStudent, s); err != nil {
		return nil
	}

	s.Personnummer = personnummer
	s.UID = uid
	s.ExterntUID = externtUID

	b, err := json.Marshal(s)
	if err != nil {
		return nil
	}
	return b
}
