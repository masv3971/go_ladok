package goladok3

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"testing"

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
			payload:    payload{jsonSuperFeed(t), xmlFeedRecent},
			reply:      &SuperFeed{},
			statusCode: 200,
			env:        envProdAPI,
		},
		{
			name:    "Prod_GET:/uppfoljning/feed/recent 500",
			url:     "/uppfoljning/feed/recent",
			payload: payload{jsonSuperFeed(t), jsonErrors500},
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
			name:       "IntTest_GET:/handelse/feed/recent 200",
			url:        "/handelse/feed/recent",
			payload:    payload{jsonSuperFeed(t), xmlFeedRecent},
			reply:      &SuperFeed{},
			statusCode: 200,
			env:        envIntTestAPI,
		},
		{
			name:       "Test_GET:/uppfoljning/feed/recent 200",
			url:        "/uppfoljning/feed/recent",
			payload:    payload{jsonSuperFeed(t), xmlFeedRecent},
			reply:      &SuperFeed{},
			statusCode: 200,
			env:        envTestAPI,
		},
		{
			name:       "Invalid ladok-environment",
			url:        "/uppfoljning/feed/recent",
			payload:    payload{jsonSuperFeed(t), xmlFeedRecent},
			reply:      &Errors{Internal: []InternalError{{Msg: "No valid ladok-environment (OU) found in certificate"}}},
			statusCode: 200,
			env:        "test",
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			mux, server, client := mockSetup(t, tt.env)

			mockGenericEndpointServer(t, mux, contentTypeAtomXML, "GET", tt.url, "", tt.payload.server, tt.statusCode)

			err := json.Unmarshal(tt.payload.client, tt.reply)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			switch tt.reply.(type) {
			case *SuperFeed:
				got, _, _ := client.Feed.FeedRecent(context.TODO())

				assert.Equal(t, tt.reply, got, "Should be equal")
			case *Errors:
				_, _, err = client.Feed.FeedRecent(context.TODO())
				assert.Equal(t, tt.reply.(*Errors), err)
			}

			server.Close() // Close server after each run
		})
	}
}

var xmlAnvandareAndraEvent = []byte(`
  <ki:AnvandareAndradEvent
    xmlns:ki="http://schemas.ladok.se/kataloginformation"
    xmlns:base="http://schemas.ladok.se"
    xmlns:dap="http://schemas.ladok.se/dap"
    xmlns:events="http://schemas.ladok.se/events">
    <events:HandelseUID>df3ca2cd-2815-11ec-b525-441c04d24542</events:HandelseUID>
    <events:EventContext>
      <events:AnvandareUID>de624944-2815-11ec-b525-441c04d24542</events:AnvandareUID>
      <events:Anvandarnamn>system@ladokintern.se</events:Anvandarnamn>
      <events:LarosateID>27</events:LarosateID>
    </events:EventContext>
    <ki:AnvandareUID>db20a822-2814-11ec-b525-441c04d24542</ki:AnvandareUID>
    <ki:Anvandarnamnet>konsortiesupport-mape5338@konstfack.se</ki:Anvandarnamnet>
    <ki:Efternamn>Konsortiesupport TestEfternamn</ki:Efternamn>
    <ki:Email>testFornamn.testEfternamn@example.com</ki:Email>
    <ki:Fornamn>testFornamn</ki:Fornamn>
  </ki:AnvandareAndradEvent>
`)

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

var xmlAnvandareSkapadEvent = []byte(`
      <ki:AnvandareSkapadEvent
        xmlns:ki="http://schemas.ladok.se/kataloginformation"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>deeef7f0-2815-11ec-b525-441c04d24542</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>de624944-2815-11ec-b525-441c04d24542</events:AnvandareUID>
          <events:Anvandarnamn>system@ladokintern.se</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <ki:AnvandareUID>db17f56c-2814-11ec-b525-441c04d24542</ki:AnvandareUID>
        <ki:Anvandarnamnet>sunet@kf</ki:Anvandarnamnet>
        <ki:Fornamn>sunet@KF</ki:Fornamn>
      </ki:AnvandareSkapadEvent>
`)

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

var xmlExternPartEvent = []byte(`
      <ki:ExternPartEvent
        xmlns:ki="http://schemas.ladok.se/kataloginformation"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>e849148a-276b-11ec-a912-d80914c94ada</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>3d284b5a-8dc6-11e5-923c-c49715df4966</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>-1</events:LarosateID>
        </events:EventContext>
        <events:Benamningar>
          <base:Benamning>
            <base:Sprakkod>sv</base:Sprakkod>
            <base:Text>Belarusian State Technological University</base:Text>
          </base:Benamning>
          <base:Benamning>
            <base:Sprakkod>en</base:Sprakkod>
            <base:Text>Belarusian State Technological University</base:Text>
          </base:Benamning>
        </events:Benamningar>
        <events:Beskrivningar>
          <base:Benamning>
            <base:Sprakkod>sv</base:Sprakkod>
            <base:Text>Ryska: Belorusskij gosudarstvennyj technologiceskij universitet</base:Text>
          </base:Benamning>
        </events:Beskrivningar>
        <events:EventTyp>SKAPAD</events:EventTyp>
        <events:Giltighetsperiod />
        <events:Id>152447</events:Id>
        <events:Kod>MINSK10</events:Kod>
        <ki:LandID>25</ki:LandID>
        <ki:TypAvExternPartID>1</ki:TypAvExternPartID>
      </ki:ExternPartEvent>
`)

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

var xmlKontaktuppgifterEvent = []byte(`
      <si:KontaktuppgifterEvent
        xmlns:si="http://schemas.ladok.se/studentinformation"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>62127c6a-27c2-11ec-b742-49fcffce49ad</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>6209f0e8-27c2-11ec-b742-49fcffce49ad</events:AnvandareUID>
          <events:Anvandarnamn>feedevent@ladokintern.se</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <events:Handelsetyp>UPPDATERAD</events:Handelsetyp>
        <si:Epostadress>testMail@example.com</si:Epostadress>
        <si:Postadresser>
          <si:CareOf>NAME</si:CareOf>
          <si:Land />
          <si:PostadressTyp>POSTADRESS</si:PostadressTyp>
          <si:Postnummer>10010</si:Postnummer>
          <si:Postort>CITY</si:Postort>
          <si:Utdelningsadress>TESTGATAN 2 LGH 1000</si:Utdelningsadress>
        </si:Postadresser>
        <si:Postadresser>
          <si:Land>Sverige</si:Land>
          <si:PostadressTyp>FOLKBOKFORINGSADRESS</si:PostadressTyp>
          <si:Postnummer>10020</si:Postnummer>
          <si:Postort>CITY</si:Postort>
          <si:Utdelningsadress>TESTGATAN 1 LGH 1000</si:Utdelningsadress>
        </si:Postadresser>
        <si:StudentUID>041e8b44-b593-11e7-96e6-896ca17746d1</si:StudentUID>
        <si:Telefonnummer>0701234567</si:Telefonnummer>
      </si:KontaktuppgifterEvent>
`)

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

var xmlResultatPaModulAttesteratEvent = []byte(`
      <rr:ResultatPaModulAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>67a12d1a-276c-11ec-a60e-c0f64d1847cf</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-07</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestFornamn TestEfterNamn</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>bd391f51-be5e-11e7-a74b-fbb589e24dac</rr:KursUID>
        <rr:KursinstansUID>bcf84602-be5e-11e7-a688-df865af0497f</rr:KursinstansUID>
        <rr:KurstillfalleUID>1aac3ee2-ae07-11e8-8034-bd68ea484fc7</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>2302</rr:BetygsgradID>
          <rr:BetygsskalaID>2</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-01</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>1.5</rr:OmfattningsPoang>
          <rr:PrestationsPoang>1.5</rr:PrestationsPoang>
          <rr:ResultatUID>fb770d5e-276b-11ec-a60e-c0f64d1847cf</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>a32402ed-52be-11e8-9ac9-7d414daf4d27</rr:StudentUID>
        <rr:UtbildningsinstansUID>bd07fd89-be5e-11e7-a688-df865af0497f</rr:UtbildningsinstansUID>
      </rr:ResultatPaModulAttesteratEvent>
`)

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

var xmlResultatPaHelKursAttesteratEvent = []byte(`
      <rr:ResultatPaHelKursAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>0e627df9-3279-11ec-871f-f5b046564fb2</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>b0289ab3-5186-11ea-8091-b70ab71540fa</events:AnvandareUID>
          <events:Anvandarnamn>TestNamn@konstfack.se</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>b0289ab3-5186-11ea-8091-b70ab71540fa</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-21</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestForOchEfternamn</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>b0289ab3-5186-11ea-8091-b70ab71540fa</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>bf010dbe-be5e-11e7-a74b-fbb589e24dac</rr:KursUID>
        <rr:KursinstansUID>c9ef5dc4-ca2c-11e9-89dc-9348f6ec4783</rr:KursinstansUID>
        <rr:KurstillfalleUID>b4294f9e-5438-11eb-bec3-d5a2938f4dea</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>101313</rr:BetygsgradID>
          <rr:BetygsskalaID>101312</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-21</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>15.0</rr:OmfattningsPoang>
          <rr:PrestationsPoang>0.0</rr:PrestationsPoang>
          <rr:ResultatUID>0e627df6-3279-11ec-871f-f5b046564fb2</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>ebac93d8-0b38-11e8-8b82-013496834cc0</rr:StudentUID>
        <rr:UtbildningsinstansUID>c9ef5dc4-ca2c-11e9-89dc-9348f6ec4783</rr:UtbildningsinstansUID>
      </rr:ResultatPaHelKursAttesteratEvent>
`)

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

var xmlLokalStudentEvent = []byte(`
<si:LokalStudentEvent
xmlns:si="http://schemas.ladok.se/studentinformation"
xmlns:base="http://schemas.ladok.se"
xmlns:dap="http://schemas.ladok.se/dap"
xmlns:events="http://schemas.ladok.se/events">
<events:HandelseUID>79a2cce2-32be-11ec-aeeb-67874d294267</events:HandelseUID>
<events:EventContext>
  <events:AnvandareUID>799b04af-32be-11ec-aeeb-67874d294267</events:AnvandareUID>
  <events:Anvandarnamn>feedevent@ladokintern.se</events:Anvandarnamn>
  <events:LarosateID>27</events:LarosateID>
</events:EventContext>
<events:Handelsetyp>UPPDATERAD</events:Handelsetyp>
<si:Efternamn>TestEfternamn</si:Efternamn>
<si:ExterntStudentUID>1e32b258-2ad3-4804-b288-11338efe6e44</si:ExterntStudentUID>
<si:Fodelsedata>1970-01-01</si:Fodelsedata>
<si:Fornamn>TestFornamn</si:Fornamn>
<si:Kon>1</si:Kon>
<si:Personnummer>197001014622</si:Personnummer>
<si:StudentUID>54871756-790b-11e7-807b-490425ec48ab</si:StudentUID>
</si:LokalStudentEvent>
`)

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
			payload: xmlAnvandareAndraEvent,
		},
		{
			name:    "AnvandareSkapadEvent",
			event:   &anvandareEvent{},
			entryID: "df7ae52e-2815-11ec-989a-cc769fd346b3",
			want:    mockAnvandareSkapadEventSuperEvent,
			payload: xmlAnvandareSkapadEvent,
		},
		{
			name:    "ExternPartEvent",
			event:   &externPartEvent{},
			entryID: "4994B8E2-B4E9-41CB-B73D-F9A26D454294",
			want:    mockExternPartEvent,
			payload: xmlExternPartEvent,
		},
		{
			name:    "KontaktuppgifterEvent",
			event:   &kontaktuppgifterEvent{},
			entryID: "63073d13-27c2-11ec-a5df-22713cb94088",
			want:    mockKontaktuppgifterEvent,
			payload: xmlKontaktuppgifterEvent,
		},
		{
			name:    "ResultatPaModulAttesteratEvent",
			event:   &resultatEvent{},
			entryID: "684731cb-276c-11ec-a5df-22713cb94088",
			want:    mockResultatPaModulAttesteratEvent,
			payload: xmlResultatPaModulAttesteratEvent,
		},
		{
			name:    "ResultatPaHelKursAttesteratEvent",
			event:   &resultatEvent{},
			entryID: "A2D30F0A-2CD6-4EBF-B814-426646030252",
			want:    mockResultatPaHelKursAttesteratEvent,
			payload: xmlResultatPaHelKursAttesteratEvent,
		},
		{
			name:    "LokalStudentEvent",
			event:   &lokalStudentEvent{},
			entryID: "36E561D5-88D4-42E0-953B-6C86FA47E299",
			want:    mockLokalStudentEvent,
			payload: xmlLokalStudentEvent,
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
	ID: 4856,
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
			payload: xmlFeedRecent,
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

var xmlFeedRecent = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<feed
  xmlns="http://www.w3.org/2005/Atom">
  <title type="text">Events for Ladok3.Uppfoljning</title>
  <link rel="self" type="application/atom+xml" href="https://api.integrationstest.ladok.se:443/uppfoljning/feed/recent" />
  <link rel="via" type="application/atom+xml" href="https://api.integrationstest.ladok.se:443/uppfoljning/feed/4856" />
  <link rel="prev-archive" type="application/atom+xml" href="https://api.integrationstest.ladok.se:443/uppfoljning/feed/4855" />
  <id>urn:id:4856</id>
  <generator uri="http://ladok.se/uppfoljning">Uppfoljning</generator>
  <updated>2021-10-14T10:22:31.994Z</updated>
  <entry>
    <category term="http://schemas.ladok.se/kataloginformation/AnvandareAndradEvent" label="Event-typ" />
    <id>e01ec574-2815-11ec-989a-cc769fd346b3</id>
    <updated>2021-10-08T08:58:14.636Z</updated>
    <content type="application/vnd.ladok+xml">
      <ki:AnvandareAndradEvent
        xmlns:ki="http://schemas.ladok.se/kataloginformation"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>df3ca2cd-2815-11ec-b525-441c04d24542</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>de624944-2815-11ec-b525-441c04d24542</events:AnvandareUID>
          <events:Anvandarnamn>system@ladokintern.se</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <ki:AnvandareUID>db20a822-2814-11ec-b525-441c04d24542</ki:AnvandareUID>
        <ki:Anvandarnamnet>konsortiesupport-mape5338@konstfack.se</ki:Anvandarnamnet>
        <ki:Efternamn>Konsortiesupport TestEfternamn</ki:Efternamn>
        <ki:Email>testFornamn.testEfternamn@example.com</ki:Email>
        <ki:Fornamn>testFornamn</ki:Fornamn>
      </ki:AnvandareAndradEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/kataloginformation/AnvandareSkapadEvent" label="Event-typ" />
    <id>df7ae52e-2815-11ec-989a-cc769fd346b3</id>
    <updated>2021-10-08T08:58:14.127Z</updated>
    <content type="application/vnd.ladok+xml">
      <ki:AnvandareSkapadEvent
        xmlns:ki="http://schemas.ladok.se/kataloginformation"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>deeef7f0-2815-11ec-b525-441c04d24542</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>de624944-2815-11ec-b525-441c04d24542</events:AnvandareUID>
          <events:Anvandarnamn>system@ladokintern.se</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <ki:AnvandareUID>db17f56c-2814-11ec-b525-441c04d24542</ki:AnvandareUID>
        <ki:Anvandarnamnet>sunet@kf</ki:Anvandarnamnet>
        <ki:Fornamn>sunet@KF</ki:Fornamn>
      </ki:AnvandareSkapadEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/studentinformation/KontaktuppgifterEvent" label="Event-typ" />
    <id>63073d13-27c2-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T23:00:36.416Z</updated>
    <content type="application/vnd.ladok+xml">
      <si:KontaktuppgifterEvent
        xmlns:si="http://schemas.ladok.se/studentinformation"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>62127c6a-27c2-11ec-b742-49fcffce49ad</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>6209f0e8-27c2-11ec-b742-49fcffce49ad</events:AnvandareUID>
          <events:Anvandarnamn>feedevent@ladokintern.se</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <events:Handelsetyp>UPPDATERAD</events:Handelsetyp>
        <si:Epostadress>testMail@example.com</si:Epostadress>
        <si:Postadresser>
          <si:CareOf>NAME</si:CareOf>
          <si:Land />
          <si:PostadressTyp>POSTADRESS</si:PostadressTyp>
          <si:Postnummer>10010</si:Postnummer>
          <si:Postort>CITY</si:Postort>
          <si:Utdelningsadress>TESTGATAN 2 LGH 1000</si:Utdelningsadress>
        </si:Postadresser>
        <si:Postadresser>
          <si:Land>Sverige</si:Land>
          <si:PostadressTyp>FOLKBOKFORINGSADRESS</si:PostadressTyp>
          <si:Postnummer>10020</si:Postnummer>
          <si:Postort>CITY</si:Postort>
          <si:Utdelningsadress>TESTGATAN 1 LGH 1000</si:Utdelningsadress>
        </si:Postadresser>
        <si:StudentUID>041e8b44-b593-11e7-96e6-896ca17746d1</si:StudentUID>
        <si:Telefonnummer>0701234567</si:Telefonnummer>
      </si:KontaktuppgifterEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/resultat/ResultatPaModulAttesteratEvent" label="Event-typ" />
    <id>684731cb-276c-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T12:45:09.021Z</updated>
    <content type="application/vnd.ladok+xml">
      <rr:ResultatPaModulAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>67a12d1a-276c-11ec-a60e-c0f64d1847cf</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-07</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestFornamn TestEfterNamn</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>bd391f51-be5e-11e7-a74b-fbb589e24dac</rr:KursUID>
        <rr:KursinstansUID>bcf84602-be5e-11e7-a688-df865af0497f</rr:KursinstansUID>
        <rr:KurstillfalleUID>1aac3ee2-ae07-11e8-8034-bd68ea484fc7</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>2302</rr:BetygsgradID>
          <rr:BetygsskalaID>2</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-01</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>1.5</rr:OmfattningsPoang>
          <rr:PrestationsPoang>1.5</rr:PrestationsPoang>
          <rr:ResultatUID>fb770d5e-276b-11ec-a60e-c0f64d1847cf</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>a32402ed-52be-11e8-9ac9-7d414daf4d27</rr:StudentUID>
        <rr:UtbildningsinstansUID>bd07fd89-be5e-11e7-a688-df865af0497f</rr:UtbildningsinstansUID>
      </rr:ResultatPaModulAttesteratEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/kataloginformation/ExternPartEvent" label="Event-typ" />
    <id>4994B8E2-B4E9-41CB-B73D-F9A26D454294</id>
    <updated>2021-10-07T12:41:35.373Z</updated>
    <content type="application/vnd.ladok+xml">
      <ki:ExternPartEvent
        xmlns:ki="http://schemas.ladok.se/kataloginformation"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>e849148a-276b-11ec-a912-d80914c94ada</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>3d284b5a-8dc6-11e5-923c-c49715df4966</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>-1</events:LarosateID>
        </events:EventContext>
        <events:Benamningar>
          <base:Benamning>
            <base:Sprakkod>sv</base:Sprakkod>
            <base:Text>Belarusian State Technological University</base:Text>
          </base:Benamning>
          <base:Benamning>
            <base:Sprakkod>en</base:Sprakkod>
            <base:Text>Belarusian State Technological University</base:Text>
          </base:Benamning>
        </events:Benamningar>
        <events:Beskrivningar>
          <base:Benamning>
            <base:Sprakkod>sv</base:Sprakkod>
            <base:Text>Ryska: Belorusskij gosudarstvennyj technologiceskij universitet</base:Text>
          </base:Benamning>
        </events:Beskrivningar>
        <events:EventTyp>SKAPAD</events:EventTyp>
        <events:Giltighetsperiod />
        <events:Id>152447</events:Id>
        <events:Kod>MINSK10</events:Kod>
        <ki:LandID>25</ki:LandID>
        <ki:TypAvExternPartID>1</ki:TypAvExternPartID>
      </ki:ExternPartEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/studentinformation/LokalStudentEvent" label="Event-typ" />
    <id>36E561D5-88D4-42E0-953B-6C86FA47E299</id>
    <updated>2021-10-07T12:41:35.373Z</updated>
    <content type="application/vnd.ladok+xml">
      <si:LokalStudentEvent
        xmlns:si="http://schemas.ladok.se/studentinformation"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>79a2cce2-32be-11ec-aeeb-67874d294267</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>799b04af-32be-11ec-aeeb-67874d294267</events:AnvandareUID>
          <events:Anvandarnamn>feedevent@ladokintern.se</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <events:Handelsetyp>UPPDATERAD</events:Handelsetyp>
        <si:Efternamn>TestEfternamn</si:Efternamn>
        <si:ExterntStudentUID>1e32b258-2ad3-4804-b288-11338efe6e44</si:ExterntStudentUID>
        <si:Fodelsedata>1970-01-01</si:Fodelsedata>
        <si:Fornamn>TestFornamn</si:Fornamn>
        <si:Kon>1</si:Kon>
        <si:Personnummer>197001014622</si:Personnummer>
        <si:StudentUID>54871756-790b-11e7-807b-490425ec48ab</si:StudentUID>
      </si:LokalStudentEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/kataloginformation/ResultatPaHelKursAttesteratEvent" label="Event-typ" />
    <id>A2D30F0A-2CD6-4EBF-B814-426646030252</id>
    <updated>2021-10-07T12:41:35.373Z</updated>
    <content type="application/vnd.ladok+xml">
      <rr:ResultatPaHelKursAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>0e627df9-3279-11ec-871f-f5b046564fb2</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>b0289ab3-5186-11ea-8091-b70ab71540fa</events:AnvandareUID>
          <events:Anvandarnamn>TestNamn@konstfack.se</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>b0289ab3-5186-11ea-8091-b70ab71540fa</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-21</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestForOchEfternamn</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>b0289ab3-5186-11ea-8091-b70ab71540fa</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>bf010dbe-be5e-11e7-a74b-fbb589e24dac</rr:KursUID>
        <rr:KursinstansUID>c9ef5dc4-ca2c-11e9-89dc-9348f6ec4783</rr:KursinstansUID>
        <rr:KurstillfalleUID>b4294f9e-5438-11eb-bec3-d5a2938f4dea</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>101313</rr:BetygsgradID>
          <rr:BetygsskalaID>101312</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-21</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>15.0</rr:OmfattningsPoang>
          <rr:PrestationsPoang>0.0</rr:PrestationsPoang>
          <rr:ResultatUID>0e627df6-3279-11ec-871f-f5b046564fb2</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>ebac93d8-0b38-11e8-8b82-013496834cc0</rr:StudentUID>
        <rr:UtbildningsinstansUID>c9ef5dc4-ca2c-11e9-89dc-9348f6ec4783</rr:UtbildningsinstansUID>
      </rr:ResultatPaHelKursAttesteratEvent>
    </content>
  </entry>
</feed>
`)

func jsonSuperFeed(t *testing.T) []byte {
	superFeed := &SuperFeed{
		ID: 4856,
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
