package goladok3

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/masv3971/goladok3/testinginfra"
	"github.com/stretchr/testify/assert"
)

func TestFeedRecent(t *testing.T) {
	type payload struct {
		client, server []byte
	}
	tts := []struct {
		name       string
		url        string
		payload    payload
		reply      interface{}
		statusCode int
		env        string
	}{
		{
			name:       "Prod_GET:/uppfoljning/feed/recent 200",
			url:        "/uppfoljning/feed/recent",
			payload:    payload{jsonSuperFeed(t), testinginfra.XMLFeedRecent},
			reply:      &SuperFeed{},
			statusCode: 200,
			env:        envProdAPI,
		},
		{
			name:    "Prod_GET:/uppfoljning/feed/recent 500",
			url:     "/uppfoljning/feed/recent",
			payload: payload{jsonSuperFeed(t), testinginfra.JSONErrors500},
			reply: &Errors{Ladok: &LadokError{
				FelUID:          "c0f52d2c-3a5f-11ec-aa00-acd34b504da7",
				Felkategori:     "commons.fel.kategori.applikationsfel",
				FelkategoriText: "Generellt fel i applikationen",
				Meddelande:      "java.lang.NullPointerException null",
				Link:            []interface{}{},
			}},
			statusCode: 500,
			env:        envProdAPI,
		},
		{
			name:       "IntTest_GET:/handelser/feed/recent 200",
			url:        "/handelser/feed/recent",
			payload:    payload{jsonSuperFeed(t), testinginfra.XMLFeedRecent},
			reply:      &SuperFeed{},
			statusCode: 200,
			env:        envIntTestAPI,
		},
		{
			name:       "Test_GET:/uppfoljning/feed/recent 200",
			url:        "/uppfoljning/feed/recent",
			payload:    payload{jsonSuperFeed(t), testinginfra.XMLFeedRecent},
			reply:      &SuperFeed{},
			statusCode: 200,
			env:        envTestAPI,
		},
		{
			name:       "Invalid ladok-environment",
			url:        "/uppfoljning/feed/recent",
			payload:    payload{jsonSuperFeed(t), testinginfra.XMLFeedRecent},
			reply:      &Errors{Internal: []InternalError{{Msg: "No valid ladok-environment (OU) found in certificate"}}},
			statusCode: 200,
			env:        "test",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, client := mockSetup(t, tt.env)

			mockGenericEndpointServer(t, mux, ContentTypeAtomXML, "GET", tt.url, "", tt.payload.server, tt.statusCode)

			err := json.Unmarshal(tt.payload.client, tt.reply)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			switch tt.reply.(type) {
			case *SuperFeed:
				got, _, _ := client.Feed.Recent(context.TODO())

				assert.Equal(t, tt.reply, got, "Should be equal")
			case *Errors:
				_, _, err = client.Feed.Recent(context.TODO())
				assert.Equal(t, tt.reply.(*Errors), err)
			}

			server.Close() // Close server after each run
		})
	}
}

var mockAnvandareAndradEvent = &SuperEvent{
	EntryID:       "e01ec574-2815-11ec-989a-cc769fd346b3",
	EventTypeName: "AnvandareAndradEvent",
	HandelseUID:   "df3ca2cd-2815-11ec-b525-441c04d24542",
	EventContext: SuperEventContext{
		AnvandareUID: "de624944-2815-11ec-b525-441c04d24542",
		Anvandarnamn: "system@ladokintern.se",
		LarosateID:   "27",
	},
	AnvandareUID:   "db20a822-2814-11ec-b525-441c04d24542",
	Anvandarnamnet: "konsortiesupport-mape5338@konstfack.se",
	Efternamn:      "Konsortiesupport TestEfternamn",
	Fornamn:        "testFornamn",
}

var mockAnvandareSkapadEventSuperEvent = &SuperEvent{
	EntryID:       "df7ae52e-2815-11ec-989a-cc769fd346b3",
	EventTypeName: "AnvandareSkapadEvent",
	HandelseUID:   "deeef7f0-2815-11ec-b525-441c04d24542",
	EventContext: SuperEventContext{
		AnvandareUID: "de624944-2815-11ec-b525-441c04d24542",
		Anvandarnamn: "system@ladokintern.se",
		LarosateID:   "27",
	},
	AnvandareUID:   "db17f56c-2814-11ec-b525-441c04d24542",
	Anvandarnamnet: "sunet@kf",
	Fornamn:        "sunet@KF",
}

var mockExternPartEvent = &SuperEvent{
	EntryID:       "4994B8E2-B4E9-41CB-B73D-F9A26D454294",
	EventTypeName: "ExternPartEvent",
	EventContext: SuperEventContext{
		AnvandareUID: "3d284b5a-8dc6-11e5-923c-c49715df4966",
		Anvandarnamn: "testName@example.com",
		LarosateID:   "-1",
	},
	HandelseUID:       "e849148a-276b-11ec-a912-d80914c94ada",
	ID:                "152447",
	EventTyp:          "SKAPAD",
	Kod:               "MINSK10",
	LandID:            "25",
	TypAvExternPartID: "1",
}

var mockKontaktuppgifterEvent = &SuperEvent{
	EntryID:       "63073d13-27c2-11ec-a5df-22713cb94088",
	EventTypeName: "KontaktuppgifterEvent",
	EventContext: SuperEventContext{
		AnvandareUID: "6209f0e8-27c2-11ec-b742-49fcffce49ad",
		Anvandarnamn: "feedevent@ladokintern.se",
		LarosateID:   "27",
	},
	HandelseUID: "62127c6a-27c2-11ec-b742-49fcffce49ad",
	Handelsetyp: "UPPDATERAD",
	Email:       "testMail@example.com",
	Postadresser: []SuperPostadress{
		{
			PostadressTyp:    "POSTADRESS",
			Postnummer:       "10010",
			Postort:          "CITY",
			Utdelningsadress: "TESTGATAN 2 LGH 1000",
			CareOf:           "NAME",
		},
		{
			Land:             "Sverige",
			PostadressTyp:    "FOLKBOKFORINGSADRESS",
			Postnummer:       "10020",
			Postort:          "CITY",
			Utdelningsadress: "TESTGATAN 1 LGH 1000",
		},
	},
	StudentUID:    "041e8b44-b593-11e7-96e6-896ca17746d1",
	Telefonnummer: "0701234567",
}

var mockResultatPaModulAttesteratEvent = &SuperEvent{
	EntryID:       "684731cb-276c-11ec-a5df-22713cb94088",
	EventTypeName: "ResultatPaModulAttesteratEvent",
	EventContext: SuperEventContext{
		AnvandareUID: "ba1ca180-7ad2-11e9-8e63-5fd9b2d24100",
		Anvandarnamn: "testName@example.com",
		LarosateID:   "27",
	},
	HandelseUID:   "67a12d1a-276c-11ec-a60e-c0f64d1847cf",
	StudentUID:    "a32402ed-52be-11e8-9ac9-7d414daf4d27",
	Telefonnummer: "",
	Beslut: SuperBeslut{
		BeslutUID:         "ba1ca180-7ad2-11e9-8e63-5fd9b2d24100",
		Beslutsdatum:      "2021-10-07",
		Beslutsfattare:    "TestFornamn TestEfterNamn",
		BeslutsfattareUID: "ba1ca180-7ad2-11e9-8e63-5fd9b2d24100",
	},
	Resultat: SuperResultat{
		BetygsgradID:       "2302",
		BetygsskalaID:      "2",
		Examinationsdatum:  "2021-10-01",
		GiltigSomSlutbetyg: "true",
		OmfattningsPoang:   "1.5",
		PrestationsPoang:   "1.5",
		ResultatUID:        "fb770d5e-276b-11ec-a60e-c0f64d1847cf",
	},
	UtbildningsinstansUID: "bd07fd89-be5e-11e7-a688-df865af0497f",
	KursUID:               "bd391f51-be5e-11e7-a74b-fbb589e24dac",
	KursinstansUID:        "bcf84602-be5e-11e7-a688-df865af0497f",
	KurstillfalleUID:      "1aac3ee2-ae07-11e8-8034-bd68ea484fc7",
}

var mockResultatPaHelKursAttesteratEvent = &SuperEvent{
	EntryID:       "A2D30F0A-2CD6-4EBF-B814-426646030252",
	EventTypeName: "ResultatPaHelKursAttesteratEvent",
	EventContext: SuperEventContext{
		AnvandareUID: "b0289ab3-5186-11ea-8091-b70ab71540fa",
		Anvandarnamn: "TestNamn@konstfack.se",
		LarosateID:   "27",
	},
	HandelseUID: "0e627df9-3279-11ec-871f-f5b046564fb2",
	StudentUID:  "ebac93d8-0b38-11e8-8b82-013496834cc0",
	Beslut: SuperBeslut{
		BeslutUID:         "b0289ab3-5186-11ea-8091-b70ab71540fa",
		Beslutsdatum:      "2021-10-21",
		Beslutsfattare:    "TestForOchEfternamn",
		BeslutsfattareUID: "b0289ab3-5186-11ea-8091-b70ab71540fa",
	},
	Resultat: SuperResultat{
		BetygsgradID:       "101313",
		BetygsskalaID:      "101312",
		Examinationsdatum:  "2021-10-21",
		GiltigSomSlutbetyg: "true",
		OmfattningsPoang:   "15.0",
		PrestationsPoang:   "0.0",
		ResultatUID:        "0e627df6-3279-11ec-871f-f5b046564fb2",
	},
	UtbildningsinstansUID: "c9ef5dc4-ca2c-11e9-89dc-9348f6ec4783",
	KursUID:               "bf010dbe-be5e-11e7-a74b-fbb589e24dac",
	KursinstansUID:        "c9ef5dc4-ca2c-11e9-89dc-9348f6ec4783",
	KurstillfalleUID:      "b4294f9e-5438-11eb-bec3-d5a2938f4dea",
}

var mockLokalStudentEvent = &SuperEvent{
	EntryID:       "36E561D5-88D4-42E0-953B-6C86FA47E299",
	EventTypeName: "LokalStudentEvent",
	EventContext: SuperEventContext{
		AnvandareUID: "799b04af-32be-11ec-aeeb-67874d294267",
		Anvandarnamn: "feedevent@ladokintern.se",
		LarosateID:   "27",
	},
	HandelseUID:       "79a2cce2-32be-11ec-aeeb-67874d294267",
	Efternamn:         "TestEfternamn",
	Fornamn:           "TestFornamn",
	Handelsetyp:       "UPPDATERAD",
	StudentUID:        "54871756-790b-11e7-807b-490425ec48ab",
	ExterntStudentUID: "1e32b258-2ad3-4804-b288-11338efe6e44",
	Fodelsedata:       "1970-01-01",
	Kon:               "1",
	Personnummer:      "197001014622",
}

func TestParse(t *testing.T) {
	tts := []struct {
		name    string
		event   interface{}
		entryID string
		payload []byte
		want    *SuperEvent
	}{
		{
			name:    "AnvandareAndradEvent",
			event:   &anvandareEvent{},
			entryID: "e01ec574-2815-11ec-989a-cc769fd346b3",
			want:    mockAnvandareAndradEvent,
			payload: testinginfra.XMLAnvandareAndraEvent,
		},
		{
			name:    "AnvandareSkapadEvent",
			event:   &anvandareEvent{},
			entryID: "df7ae52e-2815-11ec-989a-cc769fd346b3",
			want:    mockAnvandareSkapadEventSuperEvent,
			payload: testinginfra.XMLAnvandareSkapadEvent,
		},
		{
			name:    "ExternPartEvent",
			event:   &externPartEvent{},
			entryID: "4994B8E2-B4E9-41CB-B73D-F9A26D454294",
			want:    mockExternPartEvent,
			payload: testinginfra.XMLExternPartEvent,
		},
		{
			name:    "KontaktuppgifterEvent",
			event:   &kontaktuppgifterEvent{},
			entryID: "63073d13-27c2-11ec-a5df-22713cb94088",
			want:    mockKontaktuppgifterEvent,
			payload: testinginfra.XMLKontaktuppgifterEvent,
		},
		{
			name:    "ResultatPaModulAttesteratEvent",
			event:   &resultatEvent{},
			entryID: "684731cb-276c-11ec-a5df-22713cb94088",
			want:    mockResultatPaModulAttesteratEvent,
			payload: testinginfra.XMLResultatPaModulAttesteratEvent,
		},
		{
			name:    "ResultatPaHelKursAttesteratEvent",
			event:   &resultatEvent{},
			entryID: "A2D30F0A-2CD6-4EBF-B814-426646030252",
			want:    mockResultatPaHelKursAttesteratEvent,
			payload: testinginfra.XMLResultatPaHelKursAttesteratEvent,
		},
		{
			name:    "LokalStudentEvent",
			event:   &lokalStudentEvent{},
			entryID: "36E561D5-88D4-42E0-953B-6C86FA47E299",
			want:    mockLokalStudentEvent,
			payload: testinginfra.XMLLokalStudentEvent,
		},
	}
	// ExternPartEvent.parse()
	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			var got = &SuperEvent{}

			err := xml.Unmarshal(tt.payload, tt.event)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			switch tt.event.(type) {
			case *anvandareEvent:
				got = tt.event.(*anvandareEvent).parse(tt.name, tt.entryID)
			case *externPartEvent:
				got = tt.event.(*externPartEvent).parse(tt.entryID)
			case *kontaktuppgifterEvent:
				got = tt.event.(*kontaktuppgifterEvent).parse(tt.entryID)
			case *resultatEvent:
				got = tt.event.(*resultatEvent).parse(tt.name, tt.entryID)
			case *lokalStudentEvent:
				got = tt.event.(*lokalStudentEvent).parse(tt.entryID)
			default:
				t.Fatalf("ERROR: type: %T not found", tt.event)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

var mockFeedRecent = &SuperFeed{
	ID: "4856",
	SuperEvents: []*SuperEvent{
		mockAnvandareAndradEvent,
		mockAnvandareSkapadEventSuperEvent,
		mockKontaktuppgifterEvent,
		mockResultatPaModulAttesteratEvent,
		mockExternPartEvent,
		mockLokalStudentEvent,
		mockResultatPaHelKursAttesteratEvent,
	},
}

func TestMotherParser(t *testing.T) {
	tts := []struct {
		name    string
		payload []byte
		event   *feedRecent
		want    interface{}
	}{
		{
			name:    "OK",
			payload: testinginfra.XMLFeedRecent,
			event:   &feedRecent{},
			want:    mockFeedRecent,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			err := xml.Unmarshal(tt.payload, tt.event)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			got, err := tt.event.parse()

			switch tt.want.(type) {
			case *SuperFeed:
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

func jsonSuperFeed(t *testing.T) []byte {
	superFeed := &SuperFeed{
		ID: "4856",
		SuperEvents: []*SuperEvent{
			mockAnvandareAndradEvent,
			mockAnvandareSkapadEventSuperEvent,
			mockKontaktuppgifterEvent,
			mockResultatPaModulAttesteratEvent,
			mockExternPartEvent,
			mockLokalStudentEvent,
			mockResultatPaHelKursAttesteratEvent,
		},
	}
	b, err := json.Marshal(superFeed)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	return b
}
