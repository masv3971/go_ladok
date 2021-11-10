package ladokmocks

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// MockCertificatePassword mock password for certificate
var MockCertificatePassword = "testPassword"

func mockCACertificateAndKey(t *testing.T) (*x509.Certificate, *rsa.PrivateKey, []*x509.Certificate) {
	ca := &x509.Certificate{
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		Version:            3,
		SerialNumber:       big.NewInt(2300),
		Issuer: pkix.Name{
			Organization:       []string{"Ladok3", "LED", "ca01", "ROOT", "CA"},
			OrganizationalUnit: []string{"pki-tomcat"},
			CommonName:         "CA Signing Certificate",
		},
		Subject: pkix.Name{
			Country:            []string{"SE"},
			Organization:       []string{"Ladok"},
			OrganizationalUnit: []string{"LED"},
			Locality:           []string{"Stockholm"},
			SerialNumber:       "",
			CommonName:         "sunet@KF",
		},
		NotBefore:  time.Now(),
		NotAfter:   time.Now().AddDate(10, 0, 0),
		IsCA:       true,
		KeyUsage:   x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		OCSPServer: []string{"URI:http://ca01.rome08.led.ladok.se:8080/ca/ocsp"},
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &privateKey.PublicKey, privateKey)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	signRootCA, err := x509.ParseCertificates(caBytes)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	//	caPEM := new(bytes.Buffer)
	//	pem.Encode(caPEM, &pem.Block{
	//		Type:  "CERTIFICATE",
	//		Bytes: caBytes,
	//	})

	return ca, privateKey, signRootCA
}

// MockCertificateAndKey return mock certificate template
func MockCertificateAndKey(t *testing.T, env string, notBefore, notAfter int) (*x509.Certificate, *rsa.PrivateKey, []*x509.Certificate) {
	ca, caPrivateKey, signRootCA := mockCACertificateAndKey(t)

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
		NotBefore:   time.Now().AddDate(0, 0, notBefore),
		NotAfter:    time.Now().AddDate(0, 0, notAfter),
		KeyUsage:    x509.KeyUsageDataEncipherment | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		OCSPServer:  []string{"URI:http://ca01.rome08.led.ladok.se:8080/ca/ocsp"},
	}

	certPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	certByte, err := x509.CreateCertificate(rand.Reader, certTemplate, ca, &certPrivateKey.PublicKey, caPrivateKey)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	clientCert, err := x509.ParseCertificate(certByte)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	return clientCert, certPrivateKey, signRootCA
}
