package testinginfra

// XMLAnvandareAndraEvent type
var XMLAnvandareAndraEvent = []byte(`
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
