package studentinformation

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"software.sslmate.com/src/go-pkcs12"
)

func newUUID() string { return uuid.New().String() }

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
§
func mockNew(t *testing.T, env, url string) *Client {
	pkcs12 := mockNewCertificateBundle(t, env, "test")

	cfg := Config{
		Password:     "test",
		LadokRestURL: url,
		Pkck12:       pkcs12,
	}
	client, err := New(cfg)
	assert.NoError(t, err)
	return client
}

func mockSetup(t *testing.T, env string) (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()

	//	server := httptest.NewTLSServer(mux)
	server := httptest.NewServer(mux)

	client := mockNew(t, env, server.URL)

	return mux, server, client
}

func takeDown(server *httptest.Server) { server.Close() }

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
