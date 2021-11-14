package goladok3

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/masv3971/goladok3/ladokmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mockGenericEndpointServer(t *testing.T, mux *http.ServeMux, contentType, method, url, param string, payload []byte, statusCode int) {
	mux.HandleFunc(fmt.Sprintf("%s/%s", url, param),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(statusCode)
			testMethod(t, r, method)
			testURL(t, r, fmt.Sprintf("%s/%s", url, param))
			w.Write(payload)
		},
	)
}

func mockNewClient(t *testing.T, env, url string) *Client {
	certPEM, cert, privateKeyPEM, privateKey := ladokmocks.MockCertificateAndKey(t, env, 0, 100)
	cfg := Config{
		URL:            url,
		Certificate:    cert,
		CertificatePEM: certPEM,
		PrivateKey:     privateKey,
		PrivateKeyPEM:  privateKeyPEM,
	}
	client, err := New(cfg)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	return client
}

func mockSetup(t *testing.T, env string) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)

	client := mockNewClient(t, env, server.URL)

	return mux, server, client
}

func testMethod(t *testing.T, r *http.Request, want string) {
	got := r.Method
	assert.Equal(t, want, got)
}

func testURL(t *testing.T, r *http.Request, want string) {
	got := r.RequestURI
	assert.Equal(t, want, got)
}

func testBody(t *testing.T, r *http.Request, want string) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	assert.NoError(t, err)

	got := buffer.String()
	require.JSONEq(t, want, got)
}
