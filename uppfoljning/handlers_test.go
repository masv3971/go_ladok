package uppfoljning

import (
	"context"
	"encoding/xml"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
