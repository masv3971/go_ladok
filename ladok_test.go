package goladok3

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/masv3971/goladok3/ladokmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mockGenericEndpointServer(t *testing.T, mux *http.ServeMux, contentType, method, url string, reply []byte, statusCode int) {
	mux.HandleFunc(url,
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(statusCode)
			testMethod(t, r, method)
			testURL(t, r, url)
			w.Write(reply)
		},
	)
}

func mockNewClient(t *testing.T, env, url string) *Client {
	certPEM, _, privateKeyPEM, _ := ladokmocks.MockCertificateAndKey(t, env, 0, 100)
	cfg := X509Config{
		URL: url,
		//ProxyURL:       url,
		//Certificate:    cert,
		CertificatePEM: certPEM,
		PrivateKeyPEM:  privateKeyPEM,
	}
	client, err := NewX509(cfg)
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
	assert.Equal(t, want, r.Method)
}

func testURL(t *testing.T, r *http.Request, want string) {
	assert.Equal(t, want, r.RequestURI)
}

func testBody(t *testing.T, r *http.Request, want string) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	assert.NoError(t, err)

	got := buffer.String()
	require.JSONEq(t, want, got)
}
