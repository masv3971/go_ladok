package goladok3

var jsonProfil = []byte(`{
	"Behorighetsprofiler": [{
		"Benamning": {
			"sv": "Svensk benämning"
		},
		"Dataavgransningar": {
			"LarosateID": 96,
			"Lista": [{
				"DataDimension": "ORGANISATION",
				"DataId": "01234567-1234-abcd-ef01-1234567890abcd",
				"LarosateID": 96,
				"SenastAndradAv": "testMail@example.com",
				"SenastSparad": "2012-01-11T12:45:45",
				"Uid": "11111111-2222-0000-0000-000000000000",
				"link": [{
					"method": "POST",
					"uri": "https://api.mit.ladok.se:443/test",
					"mediaType": "application/vnd.ladok+xml",
					"rel": "http://schemas.ladok.se"
				}]
			}],
			"SenastAndradAv": "testMail@example.com",
			"SenastSparad": "2012-01-11T12:45:45",
			"Uid": "11111111-2222-0000-0000-000000000000",
			"link": [{
				"method": "POST",
				"uri": "https://api.mit.ladok.se:443/test",
				"mediaType": "application/vnd.ladok+xml",
				"rel": "http://schemas.ladok.se"
			}]
		},
		"LarosateID": 96,
		"Rattighetsniva": "rattighetsniva.support",
		"SenastAndradAv": "testMail@example.com",
		"SenastSparad": "2012-01-11T12:45:45",
		"Systemaktiviteter": [{
			"Betafunktion": false,
			"I18nNyckel": "systemaktivitet.resultatrapportering",
			"Id": 2147483647,
			"KlarForProduktion": false,
			"LarosateID": 96,
			"Rattighetsniva": "rattighetsniva.support",
			"link": [{
				"method": "POST",
				"uri": "https://api.mit.ladok.se:443/test",
				"mediaType": "application/vnd.ladok+xml",
				"rel": "http://schemas.ladok.se"
			}]
		}],
		"Uid": "11111111-2222-0000-0000-000000000000",
		"link": [{
			"method": "POST",
			"uri": "https://api.mit.ladok.se:443/test",
			"mediaType": "application/vnd.ladok+xml",
			"rel": "http://schemas.ladok.se"
		}]
	}],
	"LarosateID": 96,
	"SenastAndradAv": "testMail@example.com",
	"SenastSparad": "2012-01-11T12:45:45",
	"Uid": "11111111-2222-0000-0000-000000000000",
	"link": [{
		"method": "POST",
		"uri": "https://api.mit.ladok.se:443/test",
		"mediaType": "application/vnd.ladok+xml",
		"rel": "http://schemas.ladok.se"
	}]
}`)

var jsonKataloginformationBehorighetsprofilNoPermissions = []byte(`

{
	"Benamning": {
	  "sv": "Intergration-Sunet",
	  "en": "Intergration-Sunet"
	},
	"Dataavgransningar": {
	  "Lista": [],
	  "link": []
	},
	"LarosateID": 27,
	"Rattighetsniva": "rattighetsniva.las",
	"Systemaktiviteter": [],
	"Uid": "09E52B69-5D50-4A62-B65C-636BCA68FAE5",
	"link": [
	  {
		"method": "GET",
		"uri": "https://api.integrationstest.ladok.se:443/kataloginformation/behorighetsprofil/09E52B69-5D50-4A62-B65C-636BCA68FAE5",
		"mediaType": "application/vnd.ladok+xml,application/vnd.ladok-kataloginformation+xml,application/vnd.ladok-kataloginformation+json",
		"rel": "self"
	  }
	]
  }
`)

var jsonKataloginformationBehorighetsprofil = []byte(`
{
	"Benamning": {
	  "sv": "Intergration-Sunet",
	  "en": "Intergration-Sunet"
	},
	"Dataavgransningar": {
	  "Lista": [],
	  "link": []
	},
	"LarosateID": 27,
	"Rattighetsniva": "rattighetsniva.las",
	"Systemaktiviteter": [
	  {
		"Betafunktion": false,
		"I18nNyckel": "systemaktivitet.uppfoljning.feeds",
		"Id": 90019,
		"KlarForProduktion": true,
		"Rattighetsniva": "rattighetsniva.las",
		"link": []
	  },
	  {
		"Betafunktion": false,
		"I18nNyckel": "systemaktivitet.studiedeltagande.las",
		"Id": 51001,
		"KlarForProduktion": true,
		"Rattighetsniva": "rattighetsniva.las",
		"link": []
	  },
	  {
		"Betafunktion": false,
		"I18nNyckel": "systemaktivitet.studentinformation.lasa",
		"Id": 61001,
		"KlarForProduktion": true,
		"Rattighetsniva": "rattighetsniva.las",
		"link": []
	  },
	  {
		"Betafunktion": false,
		"I18nNyckel": "systemaktivitet.kataloginformation.las",
		"Id": 11004,
		"KlarForProduktion": true,
		"Rattighetsniva": "rattighetsniva.las",
		"link": []
	  },
	  {
		"Betafunktion": false,
		"I18nNyckel": "systemaktivitet.extintegration.lasa",
		"Id": 860131,
		"KlarForProduktion": true,
		"Rattighetsniva": "rattighetsniva.las",
		"link": []
	  }
	],
	"Uid": "09E52B69-5D50-4A62-B65C-636BCA68FAE5",
	"link": [
	  {
		"method": "GET",
		"uri": "https://api.integrationstest.ladok.se:443/kataloginformation/behorighetsprofil/09E52B69-5D50-4A62-B65C-636BCA68FAE5",
		"mediaType": "application/vnd.ladok+xml,application/vnd.ladok-kataloginformation+xml,application/vnd.ladok-kataloginformation+json",
		"rel": "self"
	  }
	]
  }
`)

var jsonAutentiserad = []byte(`
{
		"Anvandarnamn": "mail@school.se",
		"Efternamn": "testEfternamn",
		"Fornamn": "testFornamn",
		"LarosateID": 96,
		"SenastAndradAv": "name@ladok3.ladok.se",
		"SenastSparad": "2012-01-11T12:45:45",
		"Uid": "11111111-2222-0000-0000-000000000000",
		"link": [ {
		  "method": "POST",
		  "uri": "https://api.mit.ladok.se:443/test",
		  "mediaType": "application/vnd.ladok+xml",
		  "rel": "http://schemas.ladok.se"
		} ]
	  }
`)

var jsonEgna = []byte(`{
		"Anvandarbehorighet": [{
		  "AnvandareRef": {
			"Anvandarnamn": "testEppn@example.com",
			"Efternamn": "TestEfternamn",
			"Fornamn": "TestFornamn",
			"Uid": "11111111-2222-0000-0000-000000000000",
			"link": {
			  "method": "POST",
			  "uri": "https://api.mit.ladok.se:443/test",
			  "mediaType": "application/vnd.ladok+xml",
			  "rel": "http://schemas.ladok.se"
			}
		  },
		  "BehorighetsprofilRef":{
			"Benamning":[{
			  "Sprakkod":"sv",
			  "Text": "Svenska",
			  "link": [ ]
			}, {
			  "Sprakkod": "en",
			  "Text": "English",
			  "link": [ ]
			}],
			"Uid": "11111111-2222-0000-0000-000000000000",
			"link": {
			  "method": "POST",
			  "uri": "https://api.mit.ladok.se:443/test",
			  "mediaType": "application/vnd.ladok+xml",
			  "rel": "http://schemas.ladok.se"
			}
		  },
		  "BestalldTidpunkt": "2013-10-14T12:45:45",
		  "LarosateID": 96,
		  "OrganisationRef": {
			"Benamning": [{
			  "Sprakkod": "sv",
			  "Text": "Svenska",
			  "link": [ ]
			}, {
			  "Sprakkod": "en",
			  "Text": "English",
			  "link": [ ]
			}],
			"Uid": "11111111-2222-0000-0000-000000000000",
			"link": {
			  "method": "POST",
			  "uri": "https://api.mit.ladok.se:443/test",
			  "mediaType": "application/vnd.ladok+xml",
			  "rel": "http://schemas.ladok.se"
			}
		  },
		  "SenastAndradAv": "testEppn@example.com",
		  "SenastSparad": "2012-01-11T12:45:45",
		  "Status": "AKTIV",
		  "Uid": "11111111-2222-0000-0000-000000000000",
		  "link": [{
			"method": "POST",
			"uri": "https://api.mit.ladok.se:443/test",
			"mediaType": "application/vnd.ladok+xml",
			"rel": "http://schemas.ladok.se"
		  }]
		}],
		"LarosateID": 96,
		"SenastAndradAv": "testEppn@example.com",
		"SenastSparad": "2012-01-11T12:45:45",
		"Uid": "11111111-2222-0000-0000-000000000000",
		"link": [{
		  "method": "POST",
		  "uri": "https://api.mit.ladok.se:443/test",
		  "mediaType": "application/vnd.ladok+xml",
		  "rel": "http://schemas.ladok.se"
		}]
	  }`)

var jsonErrors500 = []byte(`
	  {
		"FelUID": "c0f52d2c-3a5f-11ec-aa00-acd34b504da7",
		"Felkategori": "commons.fel.kategori.applikationsfel",
		"FelkategoriText": "Generellt fel i applikationen",
		"Meddelande": "java.lang.NullPointerException null",
		"link": []
	  } 
	  `)

var jsonErrorValideringsFel = []byte(`
	  {
		"Detaljkod": "commons.domain.uid",
		"DetaljkodText": "Unik identifierare",
		"FelUID": "14c837fd-3a60-11ec-aa00-acd34b504da7",
		"Felgrupp": "commons.fel.grupp.felaktigt_format",
		"FelgruppText": "Felaktigt format",
		"Felkategori": "commons.fel.kategori.valideringsfel",
		"FelkategoriText": "Valideringsfel",
		"Meddelande": "Uid: 6daf0d1e-114f-11ec-95ca-f52940734df",
		"link": []
	  } 
	  `)
