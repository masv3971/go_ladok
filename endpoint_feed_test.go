package goladok3

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
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
	client := mockNewClient(t, env, "test")

	type reply struct {
		client, server []byte
	}
	tts := []struct {
		name             string
		serverURL        string
		reply            reply
		clientReplyType  interface{}
		serverStatusCode int
		clientReq        interface{}
		clientFn         interface{}
	}{
		{
			name:             "recent",
			serverURL:        fmt.Sprintf("%s/%s", url, "recent"),
			serverStatusCode: 200,
			reply:            reply{ladokmocks.JSONSuperFeed(4856), ladokmocks.XMLFeedRecent},
			clientReplyType:  &ladoktypes.SuperFeed{},
			clientFn:         client.Feed.Recent,
		},
		{
			name:             "recent",
			serverURL:        fmt.Sprintf("%s/%s", url, "recent"),
			reply:            reply{ladokmocks.JSONSuperFeed(4856), ladokmocks.JSONErrors500},
			clientReplyType:  ladokmocks.Errors500,
			serverStatusCode: 500,
			clientFn:         client.Feed.Recent,
		},
		{
			name:             "historical",
			serverURL:        fmt.Sprintf("%s/%d", url, 100),
			reply:            reply{ladokmocks.JSONSuperFeed(4856), ladokmocks.XMLFeedRecent},
			clientReplyType:  &ladoktypes.SuperFeed{},
			serverStatusCode: 200,
			clientReq:        &HistoricalReq{ID: 100},
			clientFn:         client.Feed.Historical,
		},
		{
			name:             "historical",
			serverURL:        fmt.Sprintf("%s/%d", url, 100),
			reply:            reply{ladokmocks.JSONSuperFeed(4856), ladokmocks.JSONErrors500},
			clientReplyType:  ladokmocks.Errors500,
			serverStatusCode: 500,
			clientReq:        &HistoricalReq{ID: 100},
			clientFn:         client.Feed.Historical,
		},
		{
			name:             "first",
			serverURL:        fmt.Sprintf("%s/%s", url, "first"),
			reply:            reply{ladokmocks.JSONSuperFeed(4856), ladokmocks.XMLFeedRecent},
			clientReplyType:  &ladoktypes.SuperFeed{},
			serverStatusCode: 200,
			clientFn:         client.Feed.First,
		},
		{
			name:             "first",
			serverURL:        fmt.Sprintf("%s/%s", url, "first"),
			reply:            reply{ladokmocks.JSONSuperFeed(4856), ladokmocks.JSONErrors500},
			clientReplyType:  ladokmocks.Errors500,
			serverStatusCode: 500,
			clientFn:         client.Feed.First,
		},
	}

	for _, tt := range tts {
		t.Run(fmt.Sprintf("%s %d-- %s", tt.serverURL, tt.serverStatusCode, tt.name), func(t *testing.T) {
			mux, server, _ := mockSetup(t, env)
			client.url = server.URL
			mockGenericEndpointServer(t, mux, ContentTypeAtomXML, "GET", tt.serverURL, tt.reply.server, tt.serverStatusCode)

			err := json.Unmarshal(tt.reply.client, tt.clientReplyType)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			switch tt.clientFn.(type) {
			// Recent
			case func(context.Context) (*ladoktypes.SuperFeed, *http.Response, error):
				fn := tt.clientFn.(func(context.Context) (*ladoktypes.SuperFeed, *http.Response, error))
				switch tt.serverStatusCode {
				case 200:
					got, _, _ := fn(context.TODO())
					assert.Equal(t, tt.clientReplyType, got, "Should be equal")
				case 500:
					_, _, err = fn(context.TODO())
					assert.Equal(t, tt.clientReplyType.(*ladoktypes.LadokError), err)
				}
				// Historical
			case func(context.Context, *HistoricalReq) (*ladoktypes.SuperFeed, *http.Response, error):
				fn := tt.clientFn.(func(context.Context, *HistoricalReq) (*ladoktypes.SuperFeed, *http.Response, error))
				switch tt.serverStatusCode {
				case 200:
					got, _, _ := fn(context.TODO(), tt.clientReq.(*HistoricalReq))
					assert.Equal(t, tt.clientReplyType, got, "Should be equal")
				case 500:
					_, _, err = fn(context.TODO(), tt.clientReq.(*HistoricalReq))
					assert.Equal(t, tt.clientReplyType.(*ladoktypes.LadokError), err)
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
			case *ladoktypes.LadokError:
				assert.Equal(t, tt.want, err)
			}
		})
	}
}
