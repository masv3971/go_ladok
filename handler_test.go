package goladok3

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/masv3971/goladok3/ladokmocks"
	"github.com/masv3971/goladok3/ladoktypes"
	"github.com/stretchr/testify/assert"
)

func TestHTTPHandlers(t *testing.T) {
	var (
		client = mockNewClient(t, ladoktypes.EnvProdAPI, "test")
	)
	tts := []struct {
		name              string
		serverMethod      string
		serverURL         string
		serverReply       []byte
		serverStatusCode  int
		serverContentType string
		clientReq         interface{}
		clientReplyType   interface{}
		clientFn          interface{}
	}{
		{
			name:              "GetAnvandareAutentiserad",
			serverMethod:      "GET",
			serverURL:         "/kataloginformation/anvandare/autentiserad",
			serverReply:       ladokmocks.JSONKataloginformationAutentiserad,
			serverStatusCode:  200,
			serverContentType: ContentTypeKataloginformationJSON,
			clientReplyType:   &ladoktypes.KataloginformationAnvandareAutentiserad{},
			clientFn:          client.Kataloginformation.GetAnvandareAutentiserad,
		},
		{
			name:              "GetAnvandareAutentiserad",
			serverMethod:      "GET",
			serverURL:         "/kataloginformation/anvandare/autentiserad",
			serverContentType: ContentTypeKataloginformationJSON,
			serverReply:       ladokmocks.JSONErrors500,
			serverStatusCode:  500,
			clientReplyType:   &Errors{Ladok: ladokmocks.Errors500},
			clientFn:          client.Kataloginformation.GetAnvandareAutentiserad,
		},
		{
			name:              "GetAnvandarbehorighetEgna",
			serverMethod:      "GET",
			serverURL:         "/kataloginformation/anvandarbehorighet/egna",
			serverContentType: ContentTypeKataloginformationJSON,
			serverReply:       ladokmocks.JSONKataloginformationEgna,
			serverStatusCode:  200,
			clientReplyType:   &ladoktypes.KataloginformationAnvandarbehorighetEgna{},
			clientFn:          client.Kataloginformation.GetAnvandarbehorighetEgna,
		},
		{
			name:              "GetAnvandarbehorighetEgna",
			serverMethod:      "GET",
			serverURL:         "/kataloginformation/anvandarbehorighet/egna",
			serverReply:       ladokmocks.JSONErrors500,
			serverStatusCode:  500,
			serverContentType: ContentTypeKataloginformationJSON,
			clientReplyType:   &Errors{Ladok: ladokmocks.Errors500},
			clientFn:          client.Kataloginformation.GetAnvandarbehorighetEgna,
		},
		{
			name:              "GetBehorighetsprofil",
			serverMethod:      "GET",
			serverURL:         fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", ladokmocks.StudentUID),
			serverReply:       ladokmocks.JSONKataloginformationProfil,
			serverStatusCode:  200,
			serverContentType: ContentTypeKataloginformationJSON,
			clientReplyType:   &ladoktypes.KataloginformationBehorighetsprofil{},
			clientReq:         &GetBehorighetsprofilerReq{UID: ladokmocks.StudentUID},
			clientFn:          client.Kataloginformation.GetBehorighetsprofil,
		},
		{
			name:              "GetBehorighetsprofil",
			serverMethod:      "GET",
			serverURL:         fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", ladokmocks.StudentUID),
			serverReply:       ladokmocks.JSONErrors500,
			serverStatusCode:  500,
			serverContentType: ContentTypeKataloginformationJSON,
			clientReq:         &GetBehorighetsprofilerReq{UID: ladokmocks.StudentUID},
			clientReplyType:   &Errors{Ladok: ladokmocks.Errors500},
			clientFn:          client.Kataloginformation.GetBehorighetsprofil,
		},
		{
			name:              "GetGrunddataLarosatesinformation",
			serverMethod:      "GET",
			serverURL:         "/kataloginformation/grunddata/larosatesinformation",
			serverReply:       ladokmocks.JSONKataloginformationGrunddataLarosateinformation,
			serverStatusCode:  200,
			serverContentType: ContentTypeKataloginformationJSON,
			clientReplyType:   &ladoktypes.KataloginformationGrunddataLarosatesinformation{},
			clientFn:          client.Kataloginformation.GetGrunddataLarosatesinformation,
		},
		{
			name:              "GetGrunddataLarosatesinformation",
			serverMethod:      "GET",
			serverURL:         "/kataloginformation/grunddata/larosatesinformation",
			serverReply:       ladokmocks.JSONErrors500,
			serverStatusCode:  500,
			serverContentType: ContentTypeKataloginformationJSON,
			clientReplyType:   &Errors{Ladok: ladokmocks.Errors500},
			clientFn:          client.Kataloginformation.GetGrunddataLarosatesinformation,
		},
		{
			name:              "GetStudent UID",
			serverMethod:      "GET",
			serverURL:         fmt.Sprintf("/studentinformation/student/%s", ladokmocks.StudentUID),
			serverContentType: ContentTypeStudentinformationJSON,
			serverReply:       ladokmocks.StudentJSON(ladokmocks.StudentUID, ladokmocks.ExterntUID, ladokmocks.Personnummer),
			serverStatusCode:  200,
			clientReplyType:   &ladoktypes.Student{},
			clientReq:         &GetStudentReq{UID: ladokmocks.StudentUID},
			clientFn:          client.Studentinformation.GetStudent,
		},
		{
			name:              "GetStudent UID",
			serverMethod:      "GET",
			serverURL:         fmt.Sprintf("/studentinformation/student/%s", ladokmocks.StudentUID),
			serverReply:       ladokmocks.JSONErrors500,
			serverStatusCode:  500,
			serverContentType: ContentTypeStudentinformationJSON,
			clientReq:         &GetStudentReq{UID: ladokmocks.StudentUID},
			clientReplyType:   &Errors{Ladok: ladokmocks.Errors500},
			clientFn:          client.Studentinformation.GetStudent,
		},
		{
			name:              "GetStudent Personnummer",
			serverMethod:      "GET",
			serverURL:         fmt.Sprintf("/studentinformation/student/personnummer/%s", ladokmocks.Personnummer),
			serverContentType: ContentTypeStudentinformationJSON,
			serverReply:       ladokmocks.StudentJSON(ladokmocks.StudentUID, ladokmocks.ExterntUID, ladokmocks.Personnummer),
			serverStatusCode:  200,
			clientReplyType:   &ladoktypes.Student{},
			clientReq:         &GetStudentReq{Personnummer: ladokmocks.Personnummer},
			clientFn:          client.Studentinformation.GetStudent,
		},
		{
			name:              "GetStudent Personnummer",
			serverMethod:      "GET",
			serverURL:         fmt.Sprintf("/studentinformation/student/personnummer/%s", ladokmocks.Personnummer),
			serverReply:       ladokmocks.JSONErrors500,
			serverStatusCode:  500,
			serverContentType: ContentTypeStudentinformationJSON,
			clientReq:         &GetStudentReq{Personnummer: ladokmocks.Personnummer},
			clientReplyType:   &Errors{Ladok: ladokmocks.Errors500},
			clientFn:          client.Studentinformation.GetStudent,
		},
		{
			name:              "GetStudent ExterntUID",
			serverMethod:      "GET",
			serverURL:         fmt.Sprintf("/studentinformation/student/externtuuid/%s", ladokmocks.ExterntUID),
			serverContentType: ContentTypeStudentinformationJSON,
			serverReply:       ladokmocks.StudentJSON(ladokmocks.StudentUID, ladokmocks.ExterntUID, ladokmocks.Personnummer),
			serverStatusCode:  200,
			clientReplyType:   &ladoktypes.Student{},
			clientReq:         &GetStudentReq{ExterntUID: ladokmocks.ExterntUID},
			clientFn:          client.Studentinformation.GetStudent,
		},
		{
			name:              "GetStudent ExterntUID",
			serverMethod:      "GET",
			serverURL:         fmt.Sprintf("/studentinformation/student/externtuuid/%s", ladokmocks.ExterntUID),
			serverReply:       ladokmocks.JSONErrors500,
			serverStatusCode:  500,
			serverContentType: ContentTypeStudentinformationJSON,
			clientReq:         &GetStudentReq{ExterntUID: ladokmocks.ExterntUID},
			clientReplyType:   &Errors{Ladok: ladokmocks.Errors500},
			clientFn:          client.Studentinformation.GetStudent,
		},
		{
			name:              "GetAktivPaLarosate",
			serverMethod:      "GET",
			serverURL:         fmt.Sprintf("/studentinformation/student/%s/aktivpalarosaten", ladokmocks.StudentUID),
			serverContentType: ContentTypeStudentinformationJSON,
			serverReply:       ladokmocks.JSONAktivPaLarosate,
			serverStatusCode:  200,
			clientReplyType:   &ladoktypes.AktivPaLarosate{},
			clientReq:         &GetAktivPaLarosateReq{UID: ladokmocks.StudentUID},
			clientFn:          client.Studentinformation.GetAktivPaLarosate,
		},
		{
			name:              "GetAktivPaLarosate",
			serverMethod:      "GET",
			serverURL:         fmt.Sprintf("/studentinformation/student/%s/aktivpalarosaten", ladokmocks.StudentUID),
			serverReply:       ladokmocks.JSONErrors500,
			serverStatusCode:  500,
			serverContentType: ContentTypeStudentinformationJSON,
			clientReq:         &GetAktivPaLarosateReq{UID: ladokmocks.StudentUID},
			clientReplyType:   &Errors{Ladok: ladokmocks.Errors500},
			clientFn:          client.Studentinformation.GetAktivPaLarosate,
		},
	}

	for _, tt := range tts {
		t.Run(fmt.Sprintf("%s:%s %d -- %s", tt.serverMethod, tt.serverURL, tt.serverStatusCode, tt.name), func(t *testing.T) {
			mux, server, _ := mockSetup(t, ladoktypes.EnvIntTestAPI)
			client.url = server.URL
			defer server.Close() // Close server after each run

			mockGenericEndpointServer(t, mux, tt.serverContentType, tt.serverMethod, tt.serverURL, tt.serverReply, tt.serverStatusCode)

			// This will test if the reply will unmarshal into clientReplyType and compare it.
			err := json.Unmarshal(tt.serverReply, tt.clientReplyType)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			switch tt.clientFn.(type) {
			case func(context.Context) (*ladoktypes.KataloginformationAnvandareAutentiserad, *http.Response, error):
				f := tt.clientFn.(func(context.Context) (*ladoktypes.KataloginformationAnvandareAutentiserad, *http.Response, error))
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO())
					if !assert.NoError(t, err) {
						t.FailNow()
					}

					if !assert.Equal(t, tt.clientReplyType, reply, "Should be equal") {
						t.FailNow()
					}
				case 500:
					_, _, err = f(context.TODO())
					if !assert.Equal(t, err, tt.clientReplyType.(*Errors)) {
						t.FailNow()
					}
				}
			case func(context.Context) (*ladoktypes.KataloginformationAnvandarbehorighetEgna, *http.Response, error):
				f := tt.clientFn.(func(context.Context) (*ladoktypes.KataloginformationAnvandarbehorighetEgna, *http.Response, error))
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO())
					if !assert.NoError(t, err) {
						t.FailNow()
					}

					if !assert.Equal(t, tt.clientReplyType, reply, "Should be equal") {
						t.FailNow()
					}
				case 500:
					_, _, err = f(context.TODO())
					assert.Equal(t, err, tt.clientReplyType.(*Errors))
				}
			case func(context.Context, *GetBehorighetsprofilerReq) (*ladoktypes.KataloginformationBehorighetsprofil, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *GetBehorighetsprofilerReq) (*ladoktypes.KataloginformationBehorighetsprofil, *http.Response, error))
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), tt.clientReq.(*GetBehorighetsprofilerReq))
					if !assert.NoError(t, err) {
						t.FailNow()
					}

					if !assert.Equal(t, tt.clientReplyType, reply, "Should be equal") {
						t.FailNow()
					}
				case 500:
					_, _, err = f(context.TODO(), tt.clientReq.(*GetBehorighetsprofilerReq))
					assert.Equal(t, err, tt.clientReplyType.(*Errors))
				}
			case func(context.Context, *GetStudentReq) (*ladoktypes.Student, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *GetStudentReq) (*ladoktypes.Student, *http.Response, error))
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), tt.clientReq.(*GetStudentReq))
					if !assert.NoError(t, err) {
						t.FailNow()
					}

					if !assert.Equal(t, tt.clientReplyType, reply, "Should be equal") {
						t.FailNow()
					}
				case 500:
					_, _, err = f(context.TODO(), tt.clientReq.(*GetStudentReq))
					assert.Equal(t, err, tt.clientReplyType.(*Errors))
				}
			case func(context.Context, *GetAktivPaLarosateReq) (*ladoktypes.AktivPaLarosate, *http.Response, error):
				f := tt.clientFn.(func(context.Context, *GetAktivPaLarosateReq) (*ladoktypes.AktivPaLarosate, *http.Response, error))
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO(), tt.clientReq.(*GetAktivPaLarosateReq))
					if !assert.NoError(t, err) {
						t.FailNow()
					}

					if !assert.Equal(t, tt.clientReplyType, reply, "Should be equal") {
						t.FailNow()
					}

				case 500:
					_, _, err = f(context.TODO(), tt.clientReq.(*GetAktivPaLarosateReq))
					assert.Equal(t, err, tt.clientReplyType.(*Errors))
				}
			case func(context.Context) (*ladoktypes.KataloginformationGrunddataLarosatesinformation, *http.Response, error):
				f := tt.clientFn.(func(context.Context) (*ladoktypes.KataloginformationGrunddataLarosatesinformation, *http.Response, error))
				switch tt.serverStatusCode {
				case 200:
					reply, _, err := f(context.TODO())
					if !assert.NoError(t, err) {
						t.FailNow()
					}

					if !assert.Equal(t, tt.clientReplyType, reply, "Should be equal") {
						t.FailNow()
					}

				case 500:
					_, _, err = f(context.TODO())
					assert.Equal(t, err, tt.clientReplyType.(*Errors))
				}
			default:
				t.Fatalf("ERROR No function signature found! %T", tt.clientFn)
			}
		})
	}
}
