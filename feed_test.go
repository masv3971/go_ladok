package goladok3

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/masv3971/goladok3/ladokmocks"
	"github.com/masv3971/goladok3/ladoktypes"
	"github.com/stretchr/testify/assert"
)

func TestFeed(t *testing.T) {
	tts := []struct {
		env string
		url string
	}{
		{
			env: ladoktypes.EnvIntTestAPI,
			url: "/handelser/feed",
		},
		{
			env: ladoktypes.EnvTestAPI,
			url: "/uppfoljning/feed",
		},
		{
			env: ladoktypes.EnvProdAPI,
			url: "/uppfoljning/feed",
		},
	}

	for _, tt := range tts {
		t.Run(tt.env, func(t *testing.T) {
			testFeed(t, tt.env, tt.url)
		})
	}
}

func testFeed(t *testing.T, env string, url string) {
	client := mockNewClient(t, env, "")

	type payload struct {
		client, server []byte
	}
	tts := []struct {
		name       string
		serverURL  string
		param      int
		payload    payload
		reply      interface{}
		statusCode int
		fn         interface{}
	}{
		{
			name:       "recent 200",
			serverURL:  fmt.Sprintf("%s/%s", url, "recent"),
			payload:    payload{ladokmocks.JSONSuperFeed(4856), ladokmocks.XMLFeedRecent},
			reply:      &ladoktypes.SuperFeed{},
			statusCode: 200,
			fn:         client.Feed.Recent,
		},
		{
			name:       "recent 500",
			serverURL:  fmt.Sprintf("%s/%s", url, "recent"),
			payload:    payload{ladokmocks.JSONSuperFeed(4856), ladokmocks.JSONErrors500},
			reply:      &Errors{Ladok: ladokmocks.GeneralErrorMessage},
			statusCode: 500,
			fn:         client.Feed.Recent,
		},
		{
			name:       "historical 200",
			serverURL:  url,
			param:      100,
			payload:    payload{ladokmocks.JSONSuperFeed(4856), ladokmocks.XMLFeedRecent},
			reply:      &ladoktypes.SuperFeed{},
			statusCode: 200,
			fn:         client.Feed.Historical,
		},
		{
			name:       "historical 500",
			serverURL:  url,
			param:      100,
			payload:    payload{ladokmocks.JSONSuperFeed(4856), ladokmocks.JSONErrors500},
			reply:      &Errors{Ladok: ladokmocks.GeneralErrorMessage},
			statusCode: 500,
			fn:         client.Feed.Historical,
		},
		{
			name:       "first 200",
			serverURL:  fmt.Sprintf("%s/%s", url, "first"),
			payload:    payload{ladokmocks.JSONSuperFeed(4856), ladokmocks.XMLFeedRecent},
			reply:      &ladoktypes.SuperFeed{},
			statusCode: 200,
			fn:         client.Feed.First,
		},
		{
			name:       "first 500",
			serverURL:  fmt.Sprintf("%s/%s", url, "first"),
			payload:    payload{ladokmocks.JSONSuperFeed(4856), ladokmocks.JSONErrors500},
			reply:      &Errors{Ladok: ladokmocks.GeneralErrorMessage},
			statusCode: 500,
			fn:         client.Feed.First,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, _ := mockSetup(t, env)
			client.url = server.URL
			if tt.param != 0 {
				mockGenericEndpointServer(t, mux, ContentTypeAtomXML, "GET", tt.serverURL, strconv.Itoa(tt.param), tt.payload.server, tt.statusCode)
			} else {
				mockGenericEndpointServer(t, mux, ContentTypeAtomXML, "GET", tt.serverURL, "", tt.payload.server, tt.statusCode)
			}

			err := json.Unmarshal(tt.payload.client, tt.reply)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			switch tt.fn.(type) {
			case func(context.Context) (*ladoktypes.SuperFeed, *http.Response, error):
				f := tt.fn.(func(context.Context) (*ladoktypes.SuperFeed, *http.Response, error))
				switch tt.reply.(type) {
				case *ladoktypes.SuperFeed:
					got, _, _ := f(context.TODO())

					assert.Equal(t, tt.reply, got, "Should be equal")
				case *Errors:
					_, _, err = f(context.TODO())
					assert.Equal(t, tt.reply.(*Errors), err)
				}
			case func(context.Context, int) (*ladoktypes.SuperFeed, *http.Response, error):
				f := tt.fn.(func(context.Context, int) (*ladoktypes.SuperFeed, *http.Response, error))
				switch tt.reply.(type) {
				case *ladoktypes.SuperFeed:
					got, _, _ := f(context.TODO(), tt.param)

					assert.Equal(t, tt.reply, got, "Should be equal")
				case *Errors:
					_, _, err = f(context.TODO(), tt.param)
					assert.Equal(t, tt.reply.(*Errors), err)
				}
			}

			server.Close() // Close server after each run
		})
	}
}

func TestParse(t *testing.T) {
	tts := []struct {
		name    string
		event   interface{}
		entryID string
		payload []byte
		want    *ladoktypes.SuperEvent
	}{
		{
			name:    "AnvandareAndradEvent",
			event:   &ladoktypes.AnvandareEvent{},
			entryID: ladokmocks.AnvandareAndradEventID,
			want:    ladokmocks.MockAnvandareAndradEvent,
			payload: ladokmocks.XMLAnvandareAndraEvent,
		},
		{
			name:    "AnvandareSkapadEvent",
			event:   &ladoktypes.AnvandareEvent{},
			entryID: ladokmocks.AnvandareSkapadEventID,
			want:    ladokmocks.MockAnvandareSkapadEventSuperEvent,
			payload: ladokmocks.XMLAnvandareSkapadEvent,
		},
		{
			name:    "ExternPartEvent",
			event:   &ladoktypes.ExternPartEvent{},
			entryID: ladokmocks.ExternPartEventID,
			want:    ladokmocks.MockExternPartEvent,
			payload: ladokmocks.XMLExternPartEvent,
		},
		{
			name:    "KontaktuppgifterEvent",
			event:   &ladoktypes.KontaktuppgifterEvent{},
			entryID: ladokmocks.KontaktuppgifterEventID,
			want:    ladokmocks.MockKontaktuppgifterEvent,
			payload: ladokmocks.XMLKontaktuppgifterEvent,
		},
		{
			name:    "ResultatPaModulAttesteratEvent",
			event:   &ladoktypes.ResultatEvent{},
			entryID: ladokmocks.ResultatPaModulAttesteratEventID,
			want:    ladokmocks.MockResultatPaModulAttesteratEvent,
			payload: ladokmocks.XMLResultatPaModulAttesteratEvent,
		},
		{
			name:    "ResultatPaHelKursAttesteratEvent",
			event:   &ladoktypes.ResultatEvent{},
			entryID: ladokmocks.ResultatPaHelKursAttesteratEventID,
			want:    ladokmocks.MockResultatPaHelKursAttesteratEvent,
			payload: ladokmocks.XMLResultatPaHelKursAttesteratEvent,
		},
		{
			name:    "LokalStudentEvent",
			event:   &ladoktypes.LokalStudentEvent{},
			entryID: ladokmocks.LokalStudentEventID,
			want:    ladokmocks.MockLokalStudentEvent,
			payload: ladokmocks.XMLLokalStudentEvent,
		},
	}
	// ExternPartEvent.parse()
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			var got = &ladoktypes.SuperEvent{}

			err := xml.Unmarshal(tt.payload, tt.event)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			switch tt.event.(type) {
			case *ladoktypes.AnvandareEvent:
				got = tt.event.(*ladoktypes.AnvandareEvent).Parse(tt.name, tt.entryID)
			case *ladoktypes.ExternPartEvent:
				got = tt.event.(*ladoktypes.ExternPartEvent).Parse(tt.entryID)
			case *ladoktypes.KontaktuppgifterEvent:
				got = tt.event.(*ladoktypes.KontaktuppgifterEvent).Parse(tt.entryID)
			case *ladoktypes.ResultatEvent:
				got = tt.event.(*ladoktypes.ResultatEvent).Parse(tt.name, tt.entryID)
			case *ladoktypes.LokalStudentEvent:
				got = tt.event.(*ladoktypes.LokalStudentEvent).Parse(tt.entryID)
			default:
				t.Fatalf("ERROR: type: %T not found", tt.event)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMotherParser(t *testing.T) {
	tts := []struct {
		name    string
		payload []byte
		event   *ladoktypes.Feed
		want    interface{}
	}{
		{
			name:    "OK",
			payload: ladokmocks.XMLFeedRecent,
			event:   &ladoktypes.Feed{},
			want:    ladokmocks.MockSuperFeed(4856),
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			err := xml.Unmarshal(tt.payload, tt.event)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			got, err := tt.event.Parse()

			switch tt.want.(type) {
			case *ladoktypes.SuperFeed:
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
