package goladok3

import (
	"context"
	"fmt"
	"testing"

	"github.com/masv3971/goladok3/ladokmocks"
	"github.com/masv3971/goladok3/ladoktypes"
	"github.com/stretchr/testify/assert"
)

func TestIsLadokPermissionsSufficient(t *testing.T) {
	var (
		uid = "11111111-2222-0000-0000-000000000000"
	)
	type serverStatusCode struct {
		egna, profil int
	}
	type serverReply struct {
		egna, profil []byte
	}
	type serverURL struct {
		egna, profil string
	}
	tts := []struct {
		name             string
		serverStatusCode serverStatusCode
		serverReply      serverReply
		serverURL        serverURL
		have             Permissions
		want             interface{}
	}{
		{
			name:             "OK",
			serverURL:        serverURL{"/kataloginformation/anvandarbehorighet/egna", fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", uid)},
			serverStatusCode: serverStatusCode{200, 200},
			serverReply:      serverReply{ladokmocks.JSONKataloginformationEgna, ladokmocks.JSONKataloginformationBehorighetsprofil},
			have:             Permissions{61001: "rattighetsniva.las", 90019: "rattighetsniva.las"},
			want:             true,
		},
		{
			name:             "Missing id 0 with permission las",
			serverURL:        serverURL{"/kataloginformation/anvandarbehorighet/egna", fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", uid)},
			have:             Permissions{61001: "rattighetsniva.las", 0: "rattighetsniva.las"},
			want:             &Errors{Internal: []ladoktypes.InternalError{{Msg: "Missing id: 0, value: \"rattighetsniva.las\"", Type: "Permission"}}},
			serverStatusCode: serverStatusCode{200, 200},
			serverReply:      serverReply{ladokmocks.JSONKataloginformationEgna, ladokmocks.JSONKataloginformationBehorighetsprofil},
		},
		{
			name:             "Empty input Permissions map",
			serverURL:        serverURL{"/kataloginformation/anvandarbehorighet/egna", fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", uid)},
			have:             Permissions{},
			want:             &Errors{Internal: []ladoktypes.InternalError{{Msg: "No permissions provided", Type: "Permission"}}},
			serverStatusCode: serverStatusCode{200, 200},
			serverReply:      serverReply{ladokmocks.JSONKataloginformationEgna, ladokmocks.JSONKataloginformationBehorighetsprofil},
		},
		{
			name:             "Ladok does not have any permissions",
			serverURL:        serverURL{"/kataloginformation/anvandarbehorighet/egna", fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", uid)},
			have:             Permissions{61001: "rattighetsniva.las", 90019: "rattighetsniva.las"},
			want:             &Errors{Internal: []ladoktypes.InternalError{{Msg: "No permissions found in ladok", Type: "Permission"}}},
			serverStatusCode: serverStatusCode{200, 200},
			serverReply:      serverReply{ladokmocks.JSONKataloginformationEgna, ladokmocks.JSONKataloginformationBehorighetsprofilNoPermissions},
		},
		{
			name:      "Egna does not respond",
			serverURL: serverURL{"/kataloginformation/anvandarbehorighet/egna", fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", uid)},
			have:      Permissions{61001: "rattighetsniva.las", 90019: "rattighetsniva.las"},
			want: &Errors{Ladok: &ladoktypes.LadokError{
				FelUID:          "c0f52d2c-3a5f-11ec-aa00-acd34b504da7",
				Felkategori:     "commons.fel.kategori.applikationsfel",
				FelkategoriText: "Generellt fel i applikationen",
				Meddelande:      "java.lang.NullPointerException null",
				Link:            []interface{}{},
			}},
			serverStatusCode: serverStatusCode{500, 200},
			serverReply:      serverReply{ladokmocks.JSONErrors500, ladokmocks.JSONKataloginformationBehorighetsprofil},
		},
		{
			name:      "Profil does not respond",
			serverURL: serverURL{"/kataloginformation/anvandarbehorighet/egna", fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", uid)},
			have:      Permissions{61001: "rattighetsniva.las", 90019: "rattighetsniva.las"},
			want: &Errors{Ladok: &ladoktypes.LadokError{
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
			serverStatusCode: serverStatusCode{200, 500},
			serverReply:      serverReply{ladokmocks.JSONKataloginformationEgna, ladokmocks.JSONErrorsValideringsFel},
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, client := mockSetup(t, ladoktypes.EnvIntTestAPI)
			defer server.Close()

			mockGenericEndpointServer(t, mux, ContentTypeKataloginformationJSON, "GET", tt.serverURL.egna, tt.serverReply.egna, tt.serverStatusCode.egna)
			mockGenericEndpointServer(t, mux, ContentTypeKataloginformationJSON, "GET", tt.serverURL.profil, tt.serverReply.profil, tt.serverStatusCode.profil)

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
