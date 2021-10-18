package goladok3

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLadokPermissionsSufficent(t *testing.T) {
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
			w.Write(payloadEgna)
		},
	)

	uid := "11111111-2222-0000-0000-000000000000"

	mux.HandleFunc(fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", uid),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentTypeKataloginformationJSON)
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", uid))
			w.Write(payloadProfil)
		},
	)

	for _, tt := range tts {
		ctx := context.TODO()
		got, err := client.IsLadokPermissionsSufficent(ctx, tt.have)
		if !assert.NoError(t, err) {
			t.Fail()
		}

		assert.Equal(t, tt.want, got)
	}
}
