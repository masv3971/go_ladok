package goladok3

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var jsonStudent = []byte(`{
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
	var (
		c = mockNewClient(t, envIntTestAPI, "")
	)
	tts := []struct {
		name       string
		url        string
		payload    []byte
		reply      interface{}
		req        *GetStudentReq
		statusCode int
		fn         func(context.Context, *GetStudentReq) (*Student, *http.Response, error)
	}{
		{
			name:       "Get:/student 200",
			url:        "/studentinformation/student",
			payload:    jsonStudent,
			reply:      &Student{},
			req:        &GetStudentReq{UID: uuid.NewString()},
			statusCode: 200,
			fn:         c.Studentinformation.GetStudent,
		},
		{
			name:    "Get:/student 500",
			url:     "/studentinformation/student",
			payload: jsonErrors500,
			reply: &Errors{Ladok: &LadokError{
				FelUID:          "c0f52d2c-3a5f-11ec-aa00-acd34b504da7",
				Felkategori:     "commons.fel.kategori.applikationsfel",
				FelkategoriText: "Generellt fel i applikationen",
				Meddelande:      "java.lang.NullPointerException null",
				Link:            []interface{}{},
			}},
			req:        &GetStudentReq{UID: uuid.NewString()},
			statusCode: 500,
			fn:         c.Studentinformation.GetStudent,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, _ := mockSetup(t, envIntTestAPI)
			c.url = server.URL

			mockGenericEndpointServer(t, mux, contentTypeStudentinformationJSON, "GET", tt.url, tt.req.UID, tt.payload, tt.statusCode)

			err := json.Unmarshal(tt.payload, tt.reply)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			switch tt.statusCode {
			case 200:
				reply, _, err := tt.fn(context.TODO(), tt.req)
				if !assert.NoError(t, err) {
					t.Fatal(err)
				}

				if !assert.Equal(t, tt.reply, reply, "Should be equal") {
					t.FailNow()
				}
			case 500:
				_, _, err = tt.fn(context.TODO(), tt.req)
				assert.Equal(t, err, tt.reply.(*Errors))
			}

			server.Close() // Close server after each run
		})
	}
}

func TestGenderString(t *testing.T) {
	tts := []struct {
		name string
		have *Student
		want string
	}{
		{
			name: "female",
			have: &Student{
				KonID: 1,
			},
			want: "female",
		},
		{
			name: "male",
			have: &Student{
				KonID: 2,
			},
			want: "male",
		},
		{
			name: "n/a",
			have: &Student{
				KonID: 10,
			},
			want: "n/a",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.have.GenderString()
			assert.Equal(t, tt.want, got)
		})
	}
}
