package goladok3

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"software.sslmate.com/src/go-pkcs12"
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

func mockNewCertificateBundle(t *testing.T, env, password string) []byte {
	certTemplate := mockCertificateTemplate(t, env)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}

	certByte, err := x509.CreateCertificate(rand.Reader, certTemplate, certTemplate, &privateKey.PublicKey, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	cert, err := x509.ParseCertificate(certByte)
	if err != nil {
		t.Fatal(err)
	}

	data, err := pkcs12.Encode(rand.Reader, privateKey, cert, []*x509.Certificate{}, password)
	if err != nil {
		t.Fatal(err)
	}

	return data
}
func mockCertificateTemplate(t *testing.T, env string) *x509.Certificate {
	certTemplate := &x509.Certificate{
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		Version:            3,
		SerialNumber:       big.NewInt(2300),
		Issuer: pkix.Name{
			SerialNumber: "",
			CommonName:   "Ladok3 LED MIT API CA",
		},
		Subject: pkix.Name{
			Country:            []string{"SE"},
			Organization:       []string{"Ladok"},
			OrganizationalUnit: []string{"LED", env},
			Locality:           []string{"Stockholm"},
			SerialNumber:       "",
			CommonName:         "sunet@KF",
		},
		NotBefore:   time.Now().AddDate(0, 0, 0),
		NotAfter:    time.Now().AddDate(0, 0, 100),
		KeyUsage:    x509.KeyUsageDataEncipherment | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		OCSPServer:  []string{"URI:http://ca01.rome08.led.ladok.se:8080/ca/ocsp"},
	}

	return certTemplate
}

func mockNewClient(t *testing.T, env, url string) *Client {
	cfg := Config{
		Password: "test",
		URL:      url,
		Pkcs12:   mockNewCertificateBundle(t, env, "test"),
	}
	client, err := New(cfg)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	return client
}

func mockSetup(t *testing.T, env string) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()

	//	server := httptest.NewTLSServer(mux)
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
