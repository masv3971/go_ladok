package goladok3

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLadokPermissionsSufficient(t *testing.T) {
	t.SkipNow()
	tts := []struct {
		name string
		have Permissions
		want Permissions
	}{
		{
			name: "OK",
			have: Permissions{
				61001: "rattighetsniva.las",
			},
			want: Permissions{
				61001: "rattighetsniva.las",
			},
		},
	}

	mux, server, client := mockSetup(t, envIntTestAPI)
	takeDown(server)

	mux.HandleFunc("/kataloginformation/anvandarbehorighet/egna",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentTypeKataloginformationJSON)
			testMethod(t, r, "GET")
			testURL(t, r, "/kataloginformation/anvandarbehorighet/egna")
			w.Write(jsonEgna)
		},
	)

	uid := "11111111-2222-0000-0000-000000000000"

	mux.HandleFunc(fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", uid),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentTypeKataloginformationJSON)
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", uid))
			w.Write(jsonProfil)
		},
	)

	for _, tt := range tts {
		ctx := context.TODO()
		got, err := client.IsLadokPermissionsSufficient(ctx, tt.have)
		if !assert.NoError(t, err) {
			t.Fail()
		}

		assert.Equal(t, tt.want, got)
	}
}

func TestSane(t *testing.T) {
	got := FeedID("urn:id:4856").sane()

	assert.Equal(t, FeedID("4856"), got, "sane")

	gotInt, err := got.int()
	assert.NoError(t, err)

	assert.Equal(t, 4856, gotInt, "int")
}
