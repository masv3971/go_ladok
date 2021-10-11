package goladokrest

func TestGetAnvandarbehorighetEgna(t *testing.T){
	json := '{
		"Anvandarbehorighet" : [ {
		  "AnvandareRef" : {
			"Anvandarnamn" : "kallep103@mit.com",
			"Efternamn" : "Pettersson",
			"Fornamn" : "Kalle",
			"Uid" : "11111111-2222-0000-0000-000000000000",
			"link" : {
			  "method" : "POST",
			  "uri" : "https://api.mit.ladok.se:443/test",
			  "mediaType" : "application/vnd.ladok+xml",
			  "rel" : "http://schemas.ladok.se"
			}
		  },
		  "BehorighetsprofilRef" : {
			"Benamning" : [ {
			  "Sprakkod" : "sv",
			  "Text" : "Svenska",
			  "link" : [ ]
			}, {
			  "Sprakkod" : "en",
			  "Text" : "English",
			  "link" : [ ]
			} ],
			"Uid" : "11111111-2222-0000-0000-000000000000",
			"link" : {
			  "method" : "POST",
			  "uri" : "https://api.mit.ladok.se:443/test",
			  "mediaType" : "application/vnd.ladok+xml",
			  "rel" : "http://schemas.ladok.se"
			}
		  },
		  "BestalldTidpunkt" : "2013-10-14T12:45:45",
		  "LarosateID" : 96,
		  "OrganisationRef" : {
			"Benamning" : [ {
			  "Sprakkod" : "sv",
			  "Text" : "Svenska",
			  "link" : [ ]
			}, {
			  "Sprakkod" : "en",
			  "Text" : "English",
			  "link" : [ ]
			} ],
			"Uid" : "11111111-2222-0000-0000-000000000000",
			"link" : {
			  "method" : "POST",
			  "uri" : "https://api.mit.ladok.se:443/test",
			  "mediaType" : "application/vnd.ladok+xml",
			  "rel" : "http://schemas.ladok.se"
			}
		  },
		  "SenastAndradAv" : "eva@ladok3.ladok.se",
		  "SenastSparad" : "2012-01-11T12:45:45",
		  "Status" : "AKTIV",
		  "Uid" : "11111111-2222-0000-0000-000000000000",
		  "link" : [ {
			"method" : "POST",
			"uri" : "https://api.mit.ladok.se:443/test",
			"mediaType" : "application/vnd.ladok+xml",
			"rel" : "http://schemas.ladok.se"
		  } ]
		} ],
		"LarosateID" : 96,
		"SenastAndradAv" : "eva@ladok3.ladok.se",
		"SenastSparad" : "2012-01-11T12:45:45",
		"Uid" : "11111111-2222-0000-0000-000000000000",
		"link" : [ {
		  "method" : "POST",
		  "uri" : "https://api.mit.ladok.se:443/test",
		  "mediaType" : "application/vnd.ladok+xml",
		  "rel" : "http://schemas.ladok.se"
		} ]
	  }'


}