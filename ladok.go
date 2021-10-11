package goladokrest

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Config configures new function
type Config struct {
	Certificate    *x509.Certificate
	Chain          []*x509.Certificate
	CertificatePEM []byte
	PrivateKey     *rsa.PrivateKey
	PrivateKeyPEM  []byte
	Password       string
	Format         string
}

// Client holds the ladok object
type Client struct {
	Certificate    *x509.Certificate
	Chain          []*x509.Certificate
	CertificatePEM []byte
	PrivateKey     *rsa.PrivateKey
	PrivateKeyPEM  []byte
	Password       string
	Format         string
	httpClient     *http.Client

	KataloginformationService *KataloginformationService
	StudentinformationService *StudentinformationService
	StudentdeltagandeService  *StudentdeltagandeService
}

// New creats a new instanace of ladok
func New(config Config) (*Client, error) {
	c := &Client{
		Certificate: config.Certificate,
		PrivateKey:  config.PrivateKey,
		Password:    config.Password,
		Format:      config.Format,
	}

	if err := c.httpConfigure(); err != nil {
		return nil, err
	}

	c.KataloginformationService = KataloginformationService{client: c, contentType: "kataloginformation"}
	c.StudentinformationService = StudentinformationService{client: c, contentType: "studentinformation"}
	c.StudentdeltagandeService = StudentdeltagandeService{client: c, contentType: "studentdeltagande"}

	return c, nil
}

func (c *Client) httpConfigure() error {
	keyPair, err := tls.X509KeyPair(c.CertificatePEM, c.CertifiatePEM)
	if err != nil {
		return err
	}

	tlsCfg := &tls.Config{
		Rand:               rand.Reader,
		Certificates:       []tls.Certificate{keyPair},
		NextProtos:         []string{},
		ClientAuth:         tls.RequireAndVerifyClientCert,
		ClientCAs:          c.Chain,
		InsecureSkipVerify: false,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		},
		PreferServerCipherSuites:    true,
		SessionTicketsDisabled:      false,
		DynamicRecordSizingDisabled: false,
		Renegotiation:               0,
		KeyLogWriter:                nil,
	}

	tlsCfg.BuildNameToCertificate()

	c.httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:     tlsCfg,
			DialContext:         nil,
			TLSHandshakeTimeout: 30 * time.Second,
		},
	}

	return nil
}

// LadokAcceptHeader map consists of different accept headers for ladok
var LadokAcceptHeader = map[string]map[string]string{
	"studentinformation": {
		"json": "application/vnd.ladok-studentinformation+json",
		"xml":  "application/vnd.ladok-studentinformation+xml",
	},
	"studentdeltagande": {
		"json": "application/vnd.ladok-studiedeltagande+json",
		"xml":  "application/vnd.ladok-studiedeltagande+xml",
	},
	"resultat": {
		"json": "application/vnd.ladok-resultat+json",
		"xml":  "application/vnd.ladok-resultat+xml",
	},
	"uppfoljning": {
		"json": "application/vnd.ladok-uppfoljning+json",
		"xml":  "application/vnd.ladok-uppfoljning+xml",
	},
	"examen": {
		"json": "application/vnd.ladok-examen+json",
		"xml":  "application/vnd.ladok-examen+xml",
	},
	"utbildningsinformation": {
		"json": "application/vnd.ladok-utbildningsinformation+json",
		"xml":  "application/vnd.ladok-utbildningsinformation+xml",
	},
	"kataloginformation": {
		"json": "application/vnd.ladok-kataloginformation+json",
		"xml":  "application/vnd.ladok-kataloginformation+xml",
	},
}

func (c *Client) newRequest(ctx context.Context, method, path, acceptHeader string, body interface{}) (*http.Request, error) {
	c.logger.Info("newRequest", "method", method, "path", path)

	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(c.Service.config.LadokRestURL)
	if err != nil {
		return nil, err
	}
	url := u.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		payload := struct {
			Data interface{} `json:"data"`
		}{
			Data: body,
		}
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", acceptHeader)

	return req, nil
}

func (c *Client) do(req *http.Request, value interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err := checkResponse(resp); err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(value); err != nil {
		return nil, err
	}

	return resp, nil
}

func checkResponse(r *http.Response) error {
	serviceName := "ladok"

	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	case 500:
		return fmt.Errorf("%s: not allowed", serviceName)
	}
	return fmt.Errorf("%s: invalid request", serviceName)
}
