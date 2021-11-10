package goladok3

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/masv3971/goladok3/ladokmocks"
	"github.com/masv3971/goladok3/ladoktypes"
	"github.com/stretchr/testify/assert"
)

func TestKataloginformation(t *testing.T) {
	var (
		client = mockNewClient(t, ladoktypes.EnvProdAPI, "")
	)
	tts := []struct {
		name       string
		url        string
		payload    []byte
		statusCode int
		param      string
		reply      interface{}
		fn         interface{}
	}{
		{
			name:       "GetAnvandareAutentiserad 200",
			url:        "/kataloginformation/anvandare/autentiserad",
			payload:    ladokmocks.JSONKataloginformationAutentiserad,
			statusCode: 200,
			reply:      &ladoktypes.KataloginformationAnvandareAutentiserad{},
			param:      "",
			fn:         client.Kataloginformation.GetAnvandareAutentiserad,
		},
		{
			name:       "GetAnvandareAutentiserad 500",
			url:        "/kataloginformation/anvandare/autentiserad",
			payload:    ladokmocks.JSONErrors500,
			statusCode: 500,
			reply: &Errors{Ladok: &ladoktypes.LadokError{
				FelUID:          "c0f52d2c-3a5f-11ec-aa00-acd34b504da7",
				Felkategori:     "commons.fel.kategori.applikationsfel",
				FelkategoriText: "Generellt fel i applikationen",
				Meddelande:      "java.lang.NullPointerException null",
				Link:            []interface{}{},
			}},
			param: "",
			fn:    client.Kataloginformation.GetAnvandareAutentiserad,
		},
		{
			name:       "GetAnvandarbehorighetEgna 200",
			url:        "/kataloginformation/anvandarbehorighet/egna",
			payload:    ladokmocks.JSONKataloginformationEgna,
			statusCode: 200,
			reply:      &ladoktypes.KataloginformationAnvandarbehorighetEgna{},
			param:      "",
			fn:         client.Kataloginformation.GetAnvandarbehorighetEgna,
		},
		{
			name:       "GetAnvandarbehorighetEgna 500",
			url:        "/kataloginformation/anvandarbehorighet/egna",
			payload:    ladokmocks.JSONErrors500,
			statusCode: 500,
			reply: &Errors{Ladok: &ladoktypes.LadokError{
				FelUID:          "c0f52d2c-3a5f-11ec-aa00-acd34b504da7",
				Felkategori:     "commons.fel.kategori.applikationsfel",
				FelkategoriText: "Generellt fel i applikationen",
				Meddelande:      "java.lang.NullPointerException null",
				Link:            []interface{}{},
			}},
			param: "",
			fn:    client.Kataloginformation.GetAnvandarbehorighetEgna,
		},
		{
			name:       "GetBehorighetsprofil 200",
			url:        "/kataloginformation/behorighetsprofil",
			payload:    ladokmocks.JSONKataloginformationProfil,
			statusCode: 200,
			reply:      &ladoktypes.KataloginformationBehorighetsprofil{},
			param:      uuid.NewString(),
			fn:         client.Kataloginformation.GetBehorighetsprofil,
		},
		{
			name:    "GetBehorighetsprofil 500",
			url:     "/kataloginformation/behorighetsprofil",
			payload: ladokmocks.JSONErrors500,
			reply: &Errors{Ladok: &ladoktypes.LadokError{
				FelUID:          "c0f52d2c-3a5f-11ec-aa00-acd34b504da7",
				Felkategori:     "commons.fel.kategori.applikationsfel",
				FelkategoriText: "Generellt fel i applikationen",
				Meddelande:      "java.lang.NullPointerException null",
				Link:            []interface{}{},
			}},
			param: uuid.NewString(),
			fn:    client.Kataloginformation.GetBehorighetsprofil,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, _ := mockSetup(t, ladoktypes.EnvIntTestAPI)
			client.url = server.URL

			mockGenericEndpointServer(t, mux, ContentTypeKataloginformationJSON, "GET", tt.url, tt.param, tt.payload, tt.statusCode)

			err := json.Unmarshal(tt.payload, tt.reply)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			switch tt.fn.(type) {
			case func(context.Context) (*ladoktypes.KataloginformationAnvandareAutentiserad, *http.Response, error):
				f := tt.fn.(func(context.Context) (*ladoktypes.KataloginformationAnvandareAutentiserad, *http.Response, error))
				switch tt.statusCode {
				case 200:
					reply, _, err := f(context.TODO())
					if !assert.NoError(t, err) {
						t.Fatal(err)
					}

					if !assert.Equal(t, tt.reply, reply, "Should be equal") {
						t.FailNow()
					}
				case 500:
					_, _, err = f(context.TODO())
					assert.Equal(t, err, tt.reply.(*Errors))
				}
			case func(context.Context) (*ladoktypes.KataloginformationAnvandarbehorighetEgna, *http.Response, error):
				f := tt.fn.(func(context.Context) (*ladoktypes.KataloginformationAnvandarbehorighetEgna, *http.Response, error))
				switch tt.statusCode {
				case 200:
					reply, _, err := f(context.TODO())
					if !assert.NoError(t, err) {
						t.Fatal(err)
					}

					if !assert.Equal(t, tt.reply, reply, "Should be equal") {
						t.FailNow()
					}
				case 500:
					_, _, err = f(context.TODO())
					assert.Equal(t, err, tt.reply.(*Errors))
				}
			case func(context.Context, *GetBehorighetsprofilerCfg) (*ladoktypes.KataloginformationBehorighetsprofil, *http.Response, error):
				f := tt.fn.(func(context.Context, *GetBehorighetsprofilerCfg) (*ladoktypes.KataloginformationBehorighetsprofil, *http.Response, error))
				switch tt.statusCode {
				case 200:
					reply, _, err := f(context.TODO(), &GetBehorighetsprofilerCfg{UID: tt.param})
					if !assert.NoError(t, err) {
						t.Fatal(err)
					}

					if !assert.Equal(t, tt.reply, reply, "Should be equal") {
						t.FailNow()
					}
				case 500:
					_, _, err = f(context.TODO(), &GetBehorighetsprofilerCfg{UID: tt.param})
					assert.Equal(t, err, tt.reply.(*Errors))
				}
			default:
				t.Fatalf("ERROR No function signature found! %T", tt.fn)
			}

			server.Close() // Close server after each run
		})
	}
}
