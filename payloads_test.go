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

var jsonAutentiserad = []byte(`{
		"Anvandarnamn": "mail@school.se",
		"Efternamn": "testEfternamn",
		"Fornamn": "testFornamn",
		"LarosateID": 96,
		"SenastAndradAv": "eva@ladok3.ladok.se",
		"SenastSparad": "2012-01-11T12:45:45",
		"Uid": "11111111-2222-0000-0000-000000000000",
		"link": [ {
		  "method": "POST",
		  "uri": "https://api.mit.ladok.se:443/test",
		  "mediaType": "application/vnd.ladok+xml",
		  "rel": "http://schemas.ladok.se"
		} ]
	  }`)

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
