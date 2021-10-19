package goladok3

import (
	"context"
	"encoding/xml"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var payloadFeedRecent = []byte(`
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
    <id>e01db400-2815-11ec-989a-cc769fd346b3</id>
    <updated>2021-10-08T08:58:14.634Z</updated>
    <content type="application/vnd.ladok+xml">
      <ki:AnvandareSkapadEvent
        xmlns:ki="http://schemas.ladok.se/kataloginformation"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>df3c54ac-2815-11ec-b525-441c04d24542</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>de624944-2815-11ec-b525-441c04d24542</events:AnvandareUID>
          <events:Anvandarnamn>system@ladokintern.se</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <ki:AnvandareUID>db20a822-2814-11ec-b525-441c04d24542</ki:AnvandareUID>
        <ki:Anvandarnamnet>konsortiesupport-test5338@example.com</ki:Anvandarnamnet>
        <ki:Efternamn>Konsortiesupport testEfternamn</ki:Efternamn>
        <ki:Fornamn>testFornamn</ki:Fornamn>
      </ki:AnvandareSkapadEvent>
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
    <category term="http://schemas.ladok.se/kataloginformation/AnvandareAndradEvent" label="Event-typ" />
    <id>df73e04a-2815-11ec-989a-cc769fd346b3</id>
    <updated>2021-10-08T08:58:13.787Z</updated>
    <content type="application/vnd.ladok+xml">
      <ki:AnvandareAndradEvent
        xmlns:ki="http://schemas.ladok.se/kataloginformation"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>debb17ad-2815-11ec-b525-441c04d24542</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>de624944-2815-11ec-b525-441c04d24542</events:AnvandareUID>
          <events:Anvandarnamn>system@ladokintern.se</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <ki:AnvandareUID>0c072e5d-f7d2-11e8-86c4-d6314f174dae</ki:AnvandareUID>
        <ki:Anvandarnamnet>TestName@example.com</ki:Anvandarnamnet>
        <ki:Efternamn>testEfternamn</ki:Efternamn>
        <ki:Fornamn>testFornamn</ki:Fornamn>
      </ki:AnvandareAndradEvent>
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
          <si:Land>Sverige</si:Land>
          <si:PostadressTyp>FOLKBOKFORINGSADRESS</si:PostadressTyp>
          <si:Postnummer>10020</si:Postnummer>
          <si:Postort>CIRY</si:Postort>
          <si:Utdelningsadress>TESTGATAN 1 LGH 1000</si:Utdelningsadress>
        </si:Postadresser>
        <si:StudentUID>041e8b44-b593-11e7-96e6-896ca17746d1</si:StudentUID>
        <si:Telefonnummer>0701234567</si:Telefonnummer>
      </si:KontaktuppgifterEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/studentinformation/KontaktuppgifterEvent" label="Event-typ" />
    <id>2639e281-27be-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T22:30:16.431Z</updated>
    <content type="application/vnd.ladok+xml">
      <si:KontaktuppgifterEvent
        xmlns:si="http://schemas.ladok.se/studentinformation"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>2546d065-27be-11ec-b742-49fcffce49ad</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>2540b5e3-27be-11ec-b742-49fcffce49ad</events:AnvandareUID>
          <events:Anvandarnamn>feedevent@ladokintern.se</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <events:Handelsetyp>UPPDATERAD</events:Handelsetyp>
        <si:Postadresser>
          <si:Land>Sverige</si:Land>
          <si:PostadressTyp>FOLKBOKFORINGSADRESS</si:PostadressTyp>
          <si:Postnummer>10030</si:Postnummer>
          <si:Postort>CITY</si:Postort>
          <si:Utdelningsadress>TESTVÄGEN 64 LGH 1000</si:Utdelningsadress>
        </si:Postadresser>
        <si:StudentUID>1665f4d8-b56e-11e7-96e6-896ca17746d1</si:StudentUID>
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
    <category term="http://schemas.ladok.se/resultat/ResultatPaModulAttesteratEvent" label="Event-typ" />
    <id>4c54312f-276c-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T12:44:22.044Z</updated>
    <content type="application/vnd.ladok+xml">
      <rr:ResultatPaModulAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>4ba10e0b-276c-11ec-a60e-c0f64d1847cf</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-07</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestFornamn TestEfternamn</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>1287b040-e79f-11ea-b014-b37d5f6645d4</rr:KursUID>
        <rr:KursinstansUID>1287b03a-e79f-11ea-b014-b37d5f6645d4</rr:KursinstansUID>
        <rr:KurstillfalleUID>2df36bfe-f91c-11eb-9a26-9cd428a04a8d</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>2302</rr:BetygsgradID>
          <rr:BetygsskalaID>2</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-01</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>1.5</rr:OmfattningsPoang>
          <rr:PrestationsPoang>1.5</rr:PrestationsPoang>
          <rr:ResultatUID>6f8b59a2-276b-11ec-a60e-c0f64d1847cf</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>79f12eb3-b5a5-11e7-96e6-896ca17746d1</rr:StudentUID>
        <rr:UtbildningsinstansUID>42e26f3c-089f-11eb-b7bb-9a753b344bfa</rr:UtbildningsinstansUID>
      </rr:ResultatPaModulAttesteratEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/resultat/ResultatPaModulAttesteratEvent" label="Event-typ" />
    <id>4c5346cc-276c-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T12:44:22.039Z</updated>
    <content type="application/vnd.ladok+xml">
      <rr:ResultatPaModulAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>4ba04ab8-276c-11ec-a60e-c0f64d1847cf</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-07</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestFornamn TestEfternamn</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>1287b040-e79f-11ea-b014-b37d5f6645d4</rr:KursUID>
        <rr:KursinstansUID>1287b03a-e79f-11ea-b014-b37d5f6645d4</rr:KursinstansUID>
        <rr:KurstillfalleUID>8b70e7aa-e79f-11ea-b014-b37d5f6645d4</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>2302</rr:BetygsgradID>
          <rr:BetygsskalaID>2</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-01</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>1.5</rr:OmfattningsPoang>
          <rr:PrestationsPoang>1.5</rr:PrestationsPoang>
          <rr:ResultatUID>cf9d2913-276b-11ec-a60e-c0f64d1847cf</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>474cc678-b439-11e9-8bbd-25378c5a4e4c</rr:StudentUID>
        <rr:UtbildningsinstansUID>42e26f3c-089f-11eb-b7bb-9a753b344bfa</rr:UtbildningsinstansUID>
      </rr:ResultatPaModulAttesteratEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/resultat/ResultatPaModulAttesteratEvent" label="Event-typ" />
    <id>4c525c69-276c-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T12:44:22.033Z</updated>
    <content type="application/vnd.ladok+xml">
      <rr:ResultatPaModulAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>4b9f6055-276c-11ec-a60e-c0f64d1847cf</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-07</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestFornamn TestEfternamn</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>1287b040-e79f-11ea-b014-b37d5f6645d4</rr:KursUID>
        <rr:KursinstansUID>1287b03a-e79f-11ea-b014-b37d5f6645d4</rr:KursinstansUID>
        <rr:KurstillfalleUID>2df36bfe-f91c-11eb-9a26-9cd428a04a8d</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>2302</rr:BetygsgradID>
          <rr:BetygsskalaID>2</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-01</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>1.5</rr:OmfattningsPoang>
          <rr:PrestationsPoang>1.5</rr:PrestationsPoang>
          <rr:ResultatUID>6f8b59a8-276b-11ec-a60e-c0f64d1847cf</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>498a74b0-b56e-11e7-96e6-896ca17746d1</rr:StudentUID>
        <rr:UtbildningsinstansUID>42e26f3c-089f-11eb-b7bb-9a753b344bfa</rr:UtbildningsinstansUID>
      </rr:ResultatPaModulAttesteratEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/resultat/ResultatPaModulAttesteratEvent" label="Event-typ" />
    <id>4c519916-276c-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T12:44:22.027Z</updated>
    <content type="application/vnd.ladok+xml">
      <rr:ResultatPaModulAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>4b9e75f2-276c-11ec-a60e-c0f64d1847cf</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-07</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestFornamn TestEfternamn</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>1287b040-e79f-11ea-b014-b37d5f6645d4</rr:KursUID>
        <rr:KursinstansUID>1287b03a-e79f-11ea-b014-b37d5f6645d4</rr:KursinstansUID>
        <rr:KurstillfalleUID>2df36bfe-f91c-11eb-9a26-9cd428a04a8d</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>2302</rr:BetygsgradID>
          <rr:BetygsskalaID>2</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-01</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>1.5</rr:OmfattningsPoang>
          <rr:PrestationsPoang>1.5</rr:PrestationsPoang>
          <rr:ResultatUID>6f8b59a5-276b-11ec-a60e-c0f64d1847cf</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>37b77c1b-c002-11e9-9d92-c3dc044b4ba2</rr:StudentUID>
        <rr:UtbildningsinstansUID>42e26f3c-089f-11eb-b7bb-9a753b344bfa</rr:UtbildningsinstansUID>
      </rr:ResultatPaModulAttesteratEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/resultat/ResultatPaModulAttesteratEvent" label="Event-typ" />
    <id>4c50fcd3-276c-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T12:44:22.019Z</updated>
    <content type="application/vnd.ladok+xml">
      <rr:ResultatPaModulAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>4b9d3d6f-276c-11ec-a60e-c0f64d1847cf</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-07</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestFornamn TestEfternamn</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>1287b040-e79f-11ea-b014-b37d5f6645d4</rr:KursUID>
        <rr:KursinstansUID>1287b03a-e79f-11ea-b014-b37d5f6645d4</rr:KursinstansUID>
        <rr:KurstillfalleUID>2df36bfe-f91c-11eb-9a26-9cd428a04a8d</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>2302</rr:BetygsgradID>
          <rr:BetygsskalaID>2</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-01</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>1.5</rr:OmfattningsPoang>
          <rr:PrestationsPoang>1.5</rr:PrestationsPoang>
          <rr:ResultatUID>6f8b599f-276b-11ec-a60e-c0f64d1847cf</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>d3033c69-ae60-11eb-9d84-4b68de2e4753</rr:StudentUID>
        <rr:UtbildningsinstansUID>42e26f3c-089f-11eb-b7bb-9a753b344bfa</rr:UtbildningsinstansUID>
      </rr:ResultatPaModulAttesteratEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/resultat/ResultatPaModulAttesteratEvent" label="Event-typ" />
    <id>4c503a80-276c-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T12:44:22.009Z</updated>
    <content type="application/vnd.ladok+xml">
      <rr:ResultatPaModulAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>4b9bb6cc-276c-11ec-a60e-c0f64d1847cf</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-07</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestFornamn TestEfternamn</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>1287b040-e79f-11ea-b014-b37d5f6645d4</rr:KursUID>
        <rr:KursinstansUID>1287b03a-e79f-11ea-b014-b37d5f6645d4</rr:KursinstansUID>
        <rr:KurstillfalleUID>2df36bfe-f91c-11eb-9a26-9cd428a04a8d</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>2302</rr:BetygsgradID>
          <rr:BetygsskalaID>2</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-01</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>1.5</rr:OmfattningsPoang>
          <rr:PrestationsPoang>1.5</rr:PrestationsPoang>
          <rr:ResultatUID>6f8b3289-276b-11ec-a60e-c0f64d1847cf</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>c9ce73e0-ba49-11e9-aed9-ed43d9c443a5</rr:StudentUID>
        <rr:UtbildningsinstansUID>42e26f3c-089f-11eb-b7bb-9a753b344bfa</rr:UtbildningsinstansUID>
      </rr:ResultatPaModulAttesteratEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/resultat/ResultatPaModulAttesteratEvent" label="Event-typ" />
    <id>4c4f501d-276c-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T12:44:22.002Z</updated>
    <content type="application/vnd.ladok+xml">
      <rr:ResultatPaModulAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>4b9aa559-276c-11ec-a60e-c0f64d1847cf</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-07</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestFornamn TestEfternamn</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>ba1ca180-7ad2-11e9-8e63-5fd9b2d24100</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>1287b040-e79f-11ea-b014-b37d5f6645d4</rr:KursUID>
        <rr:KursinstansUID>1287b03a-e79f-11ea-b014-b37d5f6645d4</rr:KursinstansUID>
        <rr:KurstillfalleUID>2df36bfe-f91c-11eb-9a26-9cd428a04a8d</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>2302</rr:BetygsgradID>
          <rr:BetygsskalaID>2</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-01</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>1.5</rr:OmfattningsPoang>
          <rr:PrestationsPoang>1.5</rr:PrestationsPoang>
          <rr:ResultatUID>6f8b599c-276b-11ec-a60e-c0f64d1847cf</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>d2a5169a-ae60-11eb-9d84-4b68de2e4753</rr:StudentUID>
        <rr:UtbildningsinstansUID>42e26f3c-089f-11eb-b7bb-9a753b344bfa</rr:UtbildningsinstansUID>
      </rr:ResultatPaModulAttesteratEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/kataloginformation/ExternPartEvent" label="Event-typ" />
    <id>e8eddf7f-276b-11ec-a5df-22713cb94088</id>
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
    <category term="http://schemas.ladok.se/resultat/ResultatPaModulAttesteratEvent" label="Event-typ" />
    <id>65944301-2768-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T12:16:26.759Z</updated>
    <content type="application/vnd.ladok+xml">
      <rr:ResultatPaModulAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>6514d971-2768-11ec-a60e-c0f64d1847cf</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>e7093468-6816-11e9-9c79-28b64e4f47fe</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>e7093468-6816-11e9-9c79-28b64e4f47fe</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-07</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestFornamn TestEftername</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>e7093468-6816-11e9-9c79-28b64e4f47fe</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>a199e4c6-be5e-11e7-a74b-fbb589e24dac</rr:KursUID>
        <rr:KursinstansUID>a18c9e5c-be5e-11e7-a688-df865af0497f</rr:KursinstansUID>
        <rr:KurstillfalleUID>73866c97-f8f4-11eb-9a26-9cd428a04a8d</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>2302</rr:BetygsgradID>
          <rr:BetygsskalaID>2</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-07</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>1.5</rr:OmfattningsPoang>
          <rr:PrestationsPoang>1.5</rr:PrestationsPoang>
          <rr:ResultatUID>3efbbcd8-2768-11ec-a60e-c0f64d1847cf</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>6cffaab5-b5a7-11e7-96e6-896ca17746d1</rr:StudentUID>
        <rr:UtbildningsinstansUID>a18c9e5d-be5e-11e7-a688-df865af0497f</rr:UtbildningsinstansUID>
      </rr:ResultatPaModulAttesteratEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/resultat/ResultatPaModulAttesteratEvent" label="Event-typ" />
    <id>6593318e-2768-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T12:16:26.753Z</updated>
    <content type="application/vnd.ladok+xml">
      <rr:ResultatPaModulAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>6513f00e-2768-11ec-a60e-c0f64d1847cf</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>e7093468-6816-11e9-9c79-28b64e4f47fe</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>e7093468-6816-11e9-9c79-28b64e4f47fe</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-07</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestFornamn TestEftername</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>e7093468-6816-11e9-9c79-28b64e4f47fe</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>a199e4c6-be5e-11e7-a74b-fbb589e24dac</rr:KursUID>
        <rr:KursinstansUID>a18c9e5c-be5e-11e7-a688-df865af0497f</rr:KursinstansUID>
        <rr:KurstillfalleUID>73866c97-f8f4-11eb-9a26-9cd428a04a8d</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>2302</rr:BetygsgradID>
          <rr:BetygsskalaID>2</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-07</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>1.5</rr:OmfattningsPoang>
          <rr:PrestationsPoang>1.5</rr:PrestationsPoang>
          <rr:ResultatUID>3f01894d-2768-11ec-a60e-c0f64d1847cf</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>78d7d53a-4dff-11e8-a562-6ec76bb54b9f</rr:StudentUID>
        <rr:UtbildningsinstansUID>a18c9e5d-be5e-11e7-a688-df865af0497f</rr:UtbildningsinstansUID>
      </rr:ResultatPaModulAttesteratEvent>
    </content>
  </entry>
  <entry>
    <category term="http://schemas.ladok.se/resultat/ResultatPaModulAttesteratEvent" label="Event-typ" />
    <id>6592472b-2768-11ec-a5df-22713cb94088</id>
    <updated>2021-10-07T12:16:26.747Z</updated>
    <content type="application/vnd.ladok+xml">
      <rr:ResultatPaModulAttesteratEvent
        xmlns:rr="http://schemas.ladok.se/resultat"
        xmlns:base="http://schemas.ladok.se"
        xmlns:dap="http://schemas.ladok.se/dap"
        xmlns:events="http://schemas.ladok.se/events">
        <events:HandelseUID>651305ab-2768-11ec-a60e-c0f64d1847cf</events:HandelseUID>
        <events:EventContext>
          <events:AnvandareUID>e7093468-6816-11e9-9c79-28b64e4f47fe</events:AnvandareUID>
          <events:Anvandarnamn>testName@example.com</events:Anvandarnamn>
          <events:LarosateID>27</events:LarosateID>
        </events:EventContext>
        <rr:Beslut>
          <rr:BeslutUID>e7093468-6816-11e9-9c79-28b64e4f47fe</rr:BeslutUID>
          <rr:Beslutsdatum>2021-10-07</rr:Beslutsdatum>
          <rr:Beslutsfattare>TestFornamn TestEftername</rr:Beslutsfattare>
          <rr:BeslutsfattareUID>e7093468-6816-11e9-9c79-28b64e4f47fe</rr:BeslutsfattareUID>
        </rr:Beslut>
        <rr:KursUID>a199e4c6-be5e-11e7-a74b-fbb589e24dac</rr:KursUID>
        <rr:KursinstansUID>a18c9e5c-be5e-11e7-a688-df865af0497f</rr:KursinstansUID>
        <rr:KurstillfalleUID>73866c97-f8f4-11eb-9a26-9cd428a04a8d</rr:KurstillfalleUID>
        <rr:Resultat>
          <rr:BetygsgradID>2302</rr:BetygsgradID>
          <rr:BetygsskalaID>2</rr:BetygsskalaID>
          <rr:Examinationsdatum>2021-10-07</rr:Examinationsdatum>
          <rr:GiltigSomSlutbetyg>true</rr:GiltigSomSlutbetyg>
          <rr:OmfattningsPoang>1.5</rr:OmfattningsPoang>
          <rr:PrestationsPoang>1.5</rr:PrestationsPoang>
          <rr:ResultatUID>3efbbcd5-2768-11ec-a60e-c0f64d1847cf</rr:ResultatUID>
        </rr:Resultat>
        <rr:StudentUID>bd8ecb49-2dd4-11ea-bc56-c5eaf8ca4667</rr:StudentUID>
        <rr:UtbildningsinstansUID>a18c9e5d-be5e-11e7-a688-df865af0497f</rr:UtbildningsinstansUID>
      </rr:ResultatPaModulAttesteratEvent>
    </content>
  </entry>
</feed>
`)

func TestFeedRecent(t *testing.T) {
	mux, server, client := mockSetup(t, envTestAPI)
	defer takeDown(server)

	mux.HandleFunc("/uppfoljning/feed/recent",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentTypeAtomXML)
			testMethod(t, r, "GET")
			testURL(t, r, "/uppfoljning/feed/recent")
			w.Write(payloadFeedRecent)
		},
	)
	_, _, err := client.UppfoljningService.FeedRecent(context.TODO())
	if !assert.NoError(t, err) {
		t.Fatal()
	}

}

func TestParse(t *testing.T) {
	d := &FeedRecent{}

	if err := xml.Unmarshal(payloadFeedRecent, d); err != nil {
		if !assert.NoError(t, err) {
			t.Fail()
		}
	}

	superFeed, err := d.parse()
	if !assert.NoError(t, err) {
		t.Fail()
	}

	assert.Equal(t, 4856, superFeed.ID)
}
