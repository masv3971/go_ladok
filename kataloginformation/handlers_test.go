package kataloginformation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAnvandareAutentiserad(t *testing.T) {
	d := &GetAnvandareAutentiseradReply{}
	if err := json.Unmarshal(jsonAutentiserad, d); err != nil {
		assert.NoError(t, err)
	}

	got, err := json.Marshal(d)
	assert.NoError(t, err)

	require.JSONEq(t, string(jsonAutentiserad), string(got))

	mux, server, client := mockSetup(t, envIntTestAPI)
	defer takeDown(server)

	mux.HandleFunc("/kataloginformation/anvandare/autentiserad",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentTypeKataloginformationJSON)
			testMethod(t, r, "GET")
			testURL(t, r, "/kataloginformation/anvandare/autentiserad")
			w.Write(jsonAutentiserad)
		},
	)
	reply, _, err := client.KataloginformationService.GetAnvandareAutentiserad(context.TODO())
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")
}

func TestGetAnvandarbehorighetEgna(t *testing.T) {
	d := &GetAnvandarbehorighetEgnaReply{}
	if err := json.Unmarshal(jsonEgna, d); err != nil {
		assert.NoError(t, err)
	}

	got, err := json.Marshal(d)
	assert.NoError(t, err)

	require.JSONEq(t, string(jsonEgna), string(got))

	mux, server, client := mockSetup(t, envIntTestAPI)
	defer takeDown(server)

	mux.HandleFunc("/kataloginformation/anvandarbehorighet/egna",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentTypeKataloginformationJSON)
			testMethod(t, r, "GET")
			testURL(t, r, "/kataloginformation/anvandarbehorighet/egna")
			w.Write(jsonEgna)
		},
	)
	reply, _, err := client.KataloginformationService.GetAnvandarbehorighetEgna(context.TODO())
	assert.NoError(t, err)

	assert.Equal(t, d, reply, "Should be equal")
}

func TestGetBehorighetsprofil(t *testing.T) {
	d := &GetBehorighetsprofilReply{}
	if err := json.Unmarshal(jsonProfil, d); err != nil {
		assert.NoError(t, err)
	}

	got, err := json.Marshal(d)
	assert.NoError(t, err)

	require.JSONEq(t, string(jsonProfil), string(got))

	mux, server, client := mockSetup(t, envIntTestAPI)
	defer takeDown(server)

	cfg := &GetBehorighetsprofilerCfg{
		UID: newUUID(),
	}

	mux.HandleFunc(fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", cfg.UID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentTypeKataloginformationJSON)
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/kataloginformation/behorighetsprofil/%s", cfg.UID))
			w.Write(jsonProfil)
		},
	)
	reply, _, err := client.KataloginformationService.GetBehorighetsprofil(context.TODO(), cfg)
	assert.NoError(t, err)

	assert.Equal(t, d, reply, "Should be equal")
}
