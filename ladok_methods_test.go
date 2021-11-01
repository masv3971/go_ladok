package goladok3

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLadokPermissionsSufficient(t *testing.T) {
	var (
		uid = "11111111-2222-0000-0000-000000000000"
	)
	type statusCode struct {
		egna, profile int
	}
	type payload struct {
		egna, profile []byte
	}
	tts := []struct {
		name       string
		have       Permissions
		want       interface{}
		param      string
		statusCode statusCode
		payload    payload
	}{
		{
			name:       "OK",
			have:       Permissions{61001: "rattighetsniva.las", 90019: "rattighetsniva.las"},
			want:       true,
			param:      uid,
			statusCode: statusCode{200, 200},
			payload:    payload{jsonEgna, jsonKataloginformationBehorighetsprofil},
		},
		{
			name:       "Missing id 0 with permission las",
			have:       Permissions{61001: "rattighetsniva.las", 0: "rattighetsniva.las"},
			want:       &Errors{Internal: []InternalError{{Msg: "Missing id: 0, value: \"rattighetsniva.las\"", Type: "Permission"}}},
			param:      uid,
			statusCode: statusCode{200, 200},
			payload:    payload{jsonEgna, jsonKataloginformationBehorighetsprofil},
		},
		{
			name:       "Empty input Permissions map",
			have:       Permissions{},
			want:       &Errors{Internal: []InternalError{{Msg: "No permissions provided", Type: "Permission"}}},
			param:      uid,
			statusCode: statusCode{200, 200},
			payload:    payload{jsonEgna, jsonKataloginformationBehorighetsprofil},
		},
		{
			name:       "Ladok does not have any permissions",
			have:       Permissions{61001: "rattighetsniva.las", 90019: "rattighetsniva.las"},
			want:       &Errors{Internal: []InternalError{{Msg: "No permissions found in ladok", Type: "Permission"}}},
			param:      uid,
			statusCode: statusCode{200, 200},
			payload:    payload{jsonEgna, jsonKataloginformationBehorighetsprofilNoPermissions},
		},
		{
			name: "Egna does not respond",
			have: Permissions{61001: "rattighetsniva.las", 90019: "rattighetsniva.las"},
			want: &Errors{Ladok: &LadokError{
				FelUID:          "c0f52d2c-3a5f-11ec-aa00-acd34b504da7",
				Felkategori:     "commons.fel.kategori.applikationsfel",
				FelkategoriText: "Generellt fel i applikationen",
				Meddelande:      "java.lang.NullPointerException null",
				Link:            []interface{}{},
			}},
			param:      uid,
			statusCode: statusCode{500, 200},
			payload:    payload{jsonErrors500, jsonKataloginformationBehorighetsprofil},
		},
		{
			name: "Profil does not respond",
			have: Permissions{61001: "rattighetsniva.las", 90019: "rattighetsniva.las"},
			want: &Errors{Ladok: &LadokError{
				Detaljkod:       "commons.domain.uid",
				DetaljkodText:   "Unik identifierare",
				FelUID:          "14c837fd-3a60-11ec-aa00-acd34b504da7",
				Felgrupp:        "commons.fel.grupp.felaktigt_format",
				FelgruppText:    "Felaktigt format",
				Felkategori:     "commons.fel.kategori.valideringsfel",
				FelkategoriText: "Valideringsfel",
				Meddelande:      "Uid: 6daf0d1e-114f-11ec-95ca-f52940734df",
				Link:            []interface{}{},
			}},
			param:      uid,
			statusCode: statusCode{200, 500},
			payload:    payload{jsonEgna, jsonErrorValideringsFel},
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, client := mockSetup(t, envIntTestAPI)
			defer server.Close()

			mockGenericEndpointServer(t, mux, contentTypeKataloginformationJSON, "GET", "/kataloginformation/anvandarbehorighet/egna", "", tt.payload.egna, tt.statusCode.egna)
			mockGenericEndpointServer(t, mux, contentTypeKataloginformationJSON, "GET", "/kataloginformation/behorighetsprofil", tt.param, tt.payload.profile, tt.statusCode.profile)

			got, err := client.IsLadokPermissionsSufficient(context.TODO(), tt.have)

			switch tt.want.(type) {
			case bool:
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, got)
			case *Errors:
				assert.Equal(t, tt.want, err)
			}
		})
	}
}

func TestFeedIDTrim(t *testing.T) {
	tts := []struct {
		name string
		have FeedID
		want FeedID
	}{
		{
			name: "OK",
			have: "urn:id:4856",
			want: "4856",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.have.trim())
		})
	}
}

func TestFeedIDInt(t *testing.T) {
	tts := []struct {
		name string
		have FeedID
		want int
	}{
		{
			name: "OK",
			have: "4856",
			want: 4856,
		},
		{
			name: "OK leading zero removed",
			have: "0856",
			want: 856,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.have.int()
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			assert.Equal(t, tt.want, got, "int")
		})
	}
}
