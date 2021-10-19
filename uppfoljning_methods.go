package goladok3

func (feedRecent *FeedRecent) parse() (*SuperFeed, error) {
	superFeed := &SuperFeed{}

	feedID, err := feedRecent.ID.sane().int()
	if err != nil {
		return nil, err
	}

	superFeed.ID = feedID

	for _, entry := range feedRecent.Entry {
		if entry.Content.AnvandareAndradEvent != nil {
			event := entry.Content.AnvandareAndradEvent.parse("AnvandareAndradEvent")
			superFeed.SuperEvents = append(superFeed.SuperEvents, event)
			continue
		}

		if entry.Content.AnvandareSkapadEvent != nil {
			event := entry.Content.AnvandareSkapadEvent.parse("AnvandareSkapadEven")
			superFeed.SuperEvents = append(superFeed.SuperEvents, event)
			continue
		}

		if entry.Content.ExternPartEvent != nil {
			event := entry.Content.ExternPartEvent.parse()
			superFeed.SuperEvents = append(superFeed.SuperEvents, event)
			continue
		}

		if entry.Content.KontaktuppgifterEvent != nil {
			event := entry.Content.KontaktuppgifterEvent.parse()
			superFeed.SuperEvents = append(superFeed.SuperEvents, event)
			continue
		}

		if entry.Content.ResultatPaModulAttesteratEvent != nil {
			event := entry.Content.ResultatPaModulAttesteratEvent.parse()
			superFeed.SuperEvents = append(superFeed.SuperEvents, event)
			continue
		}
	}

	return superFeed, nil
}

func (a *AnvandareEvent) parse(eventTypeName string) *SuperEvent {
	s := &SuperEvent{
		EventTypeName: eventTypeName,
		HandelseUID:   a.HandelseUID,
		EventContext: SuperEventContext{
			AnvandareUID: a.EventContext.AnvandareUID,
			Anvandarnamn: a.EventContext.Anvandarnamn,
			LarosateID:   a.EventContext.LarosateID,
		},
		AnvandareUID:   a.AnvandareUID,
		Anvandarnamnet: a.Anvandarnamnet,
		Efternamn:      a.Efternamn,
		Fornamn:        a.Fornamn,
	}
	return s
}

func (e *ExternPartEvent) parse() *SuperEvent {
	s := &SuperEvent{
		EventTypeName: "ExternPartEvent",
		HandelseUID:   e.HandelseUID,
		EventContext: SuperEventContext{
			AnvandareUID: e.EventContext.AnvandareUID,
			Anvandarnamn: e.EventContext.Anvandarnamn,
			LarosateID:   e.EventContext.LarosateID,
		},
		EventTyp:          e.EventTyp,
		Giltighetsperiod:  e.Giltighetsperiod,
		ID:                e.ID,
		Kod:               e.Kod,
		LandID:            e.LandID,
		TypAvExternPartID: e.TypAvExternPartID,
	}
	return s
}

func (k *KontaktuppgifterEvent) parse() *SuperEvent {
	s := &SuperEvent{
		EventTypeName: "KontaktuppgifterEvent",
		HandelseUID:   k.HandelseUID,
		EventContext: SuperEventContext{
			AnvandareUID: k.EventContext.AnvandareUID,
			Anvandarnamn: k.EventContext.Anvandarnamn,
			LarosateID:   k.EventContext.LarosateID,
		},
		Handelsetyp:   k.Handelsetyp,
		Epostadress:   k.Epostadress,
		StudentUID:    k.StudentUID,
		Telefonnummer: k.Telefonnummer,
	}
	return s
}

func (r *ResultatPaModulAttesteratEvent) parse() *SuperEvent {
	s := &SuperEvent{
		EventTypeName: "ResultatPaModulAttesteratEvent",
		HandelseUID:   r.HandelseUID,
		EventContext: SuperEventContext{
			AnvandareUID: r.EventContext.AnvandareUID,
			Anvandarnamn: r.EventContext.Anvandarnamn,
			LarosateID:   r.EventContext.LarosateID,
		},
		Beslut: SuperBeslut{
			BeslutUID:         r.Beslut.BeslutUID,
			Beslutsdatum:      r.Beslut.Beslutsdatum,
			Beslutsfattare:    r.Beslut.Beslutsfattare,
			BeslutsfattareUID: r.Beslut.BeslutsfattareUID,
		},
		KursUID:          r.KursUID,
		KursinstansUID:   r.KursinstansUID,
		KurstillfalleUID: r.KurstillfalleUID,
		Resultat: SuperResultat{

			BetygsgradID:       r.Resultat.BetygsgradID,
			BetygsskalaID:      r.Resultat.BetygsskalaID,
			Examinationsdatum:  r.Resultat.Examinationsdatum,
			GiltigSomSlutbetyg: r.Resultat.GiltigSomSlutbetyg,
			OmfattningsPoang:   r.Resultat.OmfattningsPoang,
			PrestationsPoang:   r.Resultat.PrestationsPoang,
			ResultatUID:        r.Resultat.ResultatUID,
		},
		StudentUID:            r.StudentUID,
		UtbildningsinstansUID: r.UtbildningsinstansUID,
	}
	return s
}
