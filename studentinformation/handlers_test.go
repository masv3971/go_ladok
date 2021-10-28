package studentinformation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetStudent(t *testing.T) {
	d := &GetStudentReply{}
	if err := json.Unmarshal(jsonStudent, d); err != nil {
		assert.NoError(t, err)
		t.Fatal()
	}

	got, err := json.Marshal(d)
	assert.NoError(t, err)

	require.JSONEq(t, string(jsonStudent), string(got))

	mux, server, client := mockSetup(t, envIntTestAPI)
	defer server.Close()

	cfg := &GetStudentCfg{
		UID: newUUID(),
	}

	mux.HandleFunc(fmt.Sprintf("/studentinformation/student/%s", cfg.UID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentTypeStudentinformationJSON)
			testMethod(t, r, "GET")
			testURL(t, r, fmt.Sprintf("/studentinformation/student/%s", cfg.UID))
			w.Write(jsonStudent)
		},
	)

	reply, _, err := client.StudentinformationService.GetStudent(context.TODO(), cfg)
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	assert.Equal(t, d, reply, "Should be equal")

}
