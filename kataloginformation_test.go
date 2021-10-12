package ladok3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAnvandareAutentiserad(t *testing.T) {
	payload := []byte(`{
		"Anvandarnamn" : "adbe0001@umu.se",
		"Efternamn" : "Andersson",
		"Fornamn" : "Georg",
		"LarosateID" : 96,
		"SenastAndradAv" : "eva@ladok3.ladok.se",
		"SenastSparad" : "2012-01-11T12:45:45",
		"Uid" : "11111111-2222-0000-0000-000000000000",
		"link" : [ {
		  "method" : "POST",
		  "uri" : "https://api.mit.ladok.se:443/test",
		  "mediaType" : "application/vnd.ladok+xml",
		  "rel" : "http://schemas.ladok.se"
		} ]
	  }`)

	d := &AnvandareAutentiseradReply{}
	if err := json.Unmarshal(payload, d); err != nil {
		assert.NoError(t, err)
	}

	got, err := json.Marshal(d)
	assert.NoError(t, err)

	require.JSONEq(t, string(payload), string(got))

	mux, server, client := mockSetup(t)
	defer takeDown(server)

	mux.HandleFunc("/kataloginformation/anvandare/autentiserad",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testURL(t, r, "/kataloginformation/anvandare/autentiserad")
			fmt.Fprint(w, string(payload))
		},
	)
	reply, _, err := client.KataloginformationService.GetAnvandareAutentiserad(context.TODO())
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestGetAnvandarbehorighetEgna(t *testing.T) {
	payload := []byte(`{
		"Anvandarbehorighet": [{
		  "AnvandareRef": {
			"Anvandarnamn": "testEppn@example.com",
			"Efternamn": "TestEfternamn",
			"Fornamn": "TestFornamn",
			"Uid": "11111111-2222-0000-0000-000000000000",
			"link": {
			  "method": "POST",
			  "uri": "https://api.mit.ladok.se:443/test",
			  "mediaType": "application/vnd.ladok+xml",
			  "rel": "http://schemas.ladok.se"
			}
		  },
		  "BehorighetsprofilRef":{
			"Benamning":[{
			  "Sprakkod":"sv",
			  "Text": "Svenska",
			  "link": [ ]
			}, {
			  "Sprakkod": "en",
			  "Text": "English",
			  "link": [ ]
			}],
			"Uid": "11111111-2222-0000-0000-000000000000",
			"link": {
			  "method": "POST",
			  "uri": "https://api.mit.ladok.se:443/test",
			  "mediaType": "application/vnd.ladok+xml",
			  "rel": "http://schemas.ladok.se"
			}
		  },
		  "BestalldTidpunkt": "2013-10-14T12:45:45",
		  "LarosateID": 96,
		  "OrganisationRef": {
			"Benamning": [{
			  "Sprakkod": "sv",
			  "Text": "Svenska",
			  "link": [ ]
			}, {
			  "Sprakkod": "en",
			  "Text": "English",
			  "link": [ ]
			}],
			"Uid": "11111111-2222-0000-0000-000000000000",
			"link": {
			  "method": "POST",
			  "uri": "https://api.mit.ladok.se:443/test",
			  "mediaType": "application/vnd.ladok+xml",
			  "rel": "http://schemas.ladok.se"
			}
		  },
		  "SenastAndradAv": "testEppn@example.com",
		  "SenastSparad": "2012-01-11T12:45:45",
		  "Status": "AKTIV",
		  "Uid": "11111111-2222-0000-0000-000000000000",
		  "link": [{
			"method": "POST",
			"uri": "https://api.mit.ladok.se:443/test",
			"mediaType": "application/vnd.ladok+xml",
			"rel": "http://schemas.ladok.se"
		  }]
		}],
		"LarosateID": 96,
		"SenastAndradAv": "testEppn@example.com",
		"SenastSparad": "2012-01-11T12:45:45",
		"Uid": "11111111-2222-0000-0000-000000000000",
		"link": [{
		  "method": "POST",
		  "uri": "https://api.mit.ladok.se:443/test",
		  "mediaType": "application/vnd.ladok+xml",
		  "rel": "http://schemas.ladok.se"
		}]
	  }`)

	d := &AnvandarbehorighetEgnaReply{}
	if err := json.Unmarshal(payload, d); err != nil {
		assert.NoError(t, err)
	}

	got, err := json.Marshal(d)
	assert.NoError(t, err)

	require.JSONEq(t, string(payload), string(got))

	mux, server, client := mockSetup(t)
	defer takeDown(server)

	mux.HandleFunc("/kataloginformation/anvandarbehorighet/egna",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testURL(t, r, "/kataloginformation/anvandarbehorighet/egna")
			fmt.Fprint(w, string(payload))
		},
	)
	reply, _, err := client.KataloginformationService.GetAnvandarbehorighetEgna(context.TODO())
	assert.NoError(t, err)

	assert.Equal(t, d, reply, "Should be equal")
}

func TestGetBehorighetsprofil(t *testing.T) {
	payload := []byte(`{
	"Behorighetsprofiler": [{
		"Benamning": {
			"sv": "Svensk benämning"
		},
		"Dataavgransningar": {
			"LarosateID": 96,
			"Lista": [{
				"DataDimension": "ORGANISATION",
				"DataId": "01234567-1234-abcd-ef01-1234567890abcd",
				"LarosateID": 96,
				"SenastAndradAv": "testMail@example.com",
				"SenastSparad": "2012-01-11T12:45:45",
				"Uid": "11111111-2222-0000-0000-000000000000",
				"link": [{
					"method": "POST",
					"uri": "https://api.mit.ladok.se:443/test",
					"mediaType": "application/vnd.ladok+xml",
					"rel": "http://schemas.ladok.se"
				}]
			}],
			"SenastAndradAv": "testMail@example.com",
			"SenastSparad": "2012-01-11T12:45:45",
			"Uid": "11111111-2222-0000-0000-000000000000",
			"link": [{
				"method": "POST",
				"uri": "https://api.mit.ladok.se:443/test",
				"mediaType": "application/vnd.ladok+xml",
				"rel": "http://schemas.ladok.se"
			}]
		},
		"LarosateID": 96,
		"Rattighetsniva": "rattighetsniva.support",
		"SenastAndradAv": "testMail@example.com",
		"SenastSparad": "2012-01-11T12:45:45",
		"Systemaktiviteter": [{
			"Betafunktion": false,
			"I18nNyckel": "systemaktivitet.resultatrapportering",
			"Id": 2147483647,
			"KlarForProduktion": false,
			"LarosateID": 96,
			"Rattighetsniva": "rattighetsniva.support",
			"link": [{
				"method": "POST",
				"uri": "https://api.mit.ladok.se:443/test",
				"mediaType": "application/vnd.ladok+xml",
				"rel": "http://schemas.ladok.se"
			}]
		}],
		"Uid": "11111111-2222-0000-0000-000000000000",
		"link": [{
			"method": "POST",
			"uri": "https://api.mit.ladok.se:443/test",
			"mediaType": "application/vnd.ladok+xml",
			"rel": "http://schemas.ladok.se"
		}]
	}],
	"LarosateID": 96,
	"SenastAndradAv": "testMail@example.com",
	"SenastSparad": "2012-01-11T12:45:45",
	"Uid": "11111111-2222-0000-0000-000000000000",
	"link": [{
		"method": "POST",
		"uri": "https://api.mit.ladok.se:443/test",
		"mediaType": "application/vnd.ladok+xml",
		"rel": "http://schemas.ladok.se"
	}]
}`)

	d := &BehorighetsprofilReply{}
	if err := json.Unmarshal(payload, d); err != nil {
		assert.NoError(t, err)
	}

	got, err := json.Marshal(d)
	assert.NoError(t, err)

	require.JSONEq(t, string(payload), string(got))

	mux, server, client := mockSetup(t)
	defer takeDown(server)

	cfg := &BehorighetsprofilerCfg{
		UID: newUUID(),
	}

	mux.HandleFunc(fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", cfg.UID),
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", cfg.UID))
			fmt.Fprint(w, string(payload))
		},
	)
	reply, _, err := client.KataloginformationService.GetBehorighetsprofil(context.TODO(), cfg)
	assert.NoError(t, err)

	assert.Equal(t, d, reply, "Should be equal")
}
