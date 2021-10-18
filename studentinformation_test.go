package goladok3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var payloadStudent = []byte(`{
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
			"Uid": "11111111-2222-0000-0000-000000000000",
			"UnikaIdentifierare": {
				"LarosateID": 96,
				"UnikIdentifierare": [{
					"LarosateID": 96,
					"SenastAndradAv": "testEppn@example.com",
					"SenastSparad": "2012-01-11T12:45:45",
					"Typ": "studentinformation.domain.unikidentifieraretyp.passnummer",
					"Uid": "11111111-2222-0000-0000-000000000000",
					"Varde": "1234",
					"link": [{
						"method": "POST",
						"uri": "https://api.mit.ladok.se:443/test",
						"mediaType": "application/vnd.ladok+xml",
						"rel": "http://schemas.ladok.se"
					}]
				}],
				"link": [{
					"method": "POST",
					"uri": "https://api.mit.ladok.se:443/test",
					"mediaType": "application/vnd.ladok+xml",
					"rel": "http://schemas.ladok.se"
				}]
			},
			"link": [{
				"method": "POST",
				"uri": "https://api.mit.ladok.se:443/test",
				"mediaType": "application/vnd.ladok+xml",
				"rel": "http://schemas.ladok.se"
			}]
		}`)

func TestGetStudent(t *testing.T) {

	d := &GetStudentReply{}
	if err := json.Unmarshal(payloadStudent, d); err != nil {
		assert.NoError(t, err)
		t.Fatal()
	}

	got, err := json.Marshal(d)
	assert.NoError(t, err)

	require.JSONEq(t, string(payloadStudent), string(got))

	mux, server, client := mockSetup(t, envIntTestAPI)
	defer takeDown(server)

	cfg := &GetStudentCfg{
		UID: newUUID(),
	}

	mux.HandleFunc(fmt.Sprintf("/studentinformation/student/%s", cfg.UID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentTypeStudentinformationJSON)
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/studentinformation/student/%s", cfg.UID))
			w.Write(payloadStudent)
		},
	)

	reply, _, err := client.StudentinformationService.GetStudent(context.TODO(), cfg)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")

}
