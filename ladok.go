package goladok3

import (
	"crypto/rsa"
	"crypto/x509"

	"github.com/masv3971/goladok3/kataloginformation"
	"github.com/masv3971/goladok3/studentinformation"
)

// Config configures new function
type Config struct {
	Password string
	//Format       string
	LadokRestURL string
	Pkck12       []byte
}

// Client holds the ladok object
type Client struct {
	HTTClient *httpClient.Client

	Kataloginformation *kataloginformation.Client
	Studentinformation *Studentinformation.Client
	Studentdeltagande *Studentdeltagande.Client
	Uppfoljning       *Uppfoljning.Service
}

// New create a new instanace of ladok
func New(config Config) (*Client, error) {
	c := &Client{
		password:     config.Password,
		format:       "json",
		ladokRestURL: config.LadokRestURL,
		pkcs12:       config.Pkck12,
	}

	c.httpClient = httpserver.New()

	c.Kataloginformation = &KataloginformationService{client: c, contentType: "kataloginformation"}
	c.StudentinformationService = &StudentinformationService{client: c, contentType: "studentinformation"}
	//c.Studentinformation = studentinformation.Service{}
	c.Studentdeltagande = &StudentdeltagandeService{client: c, contentType: "studentdeltagande"}
	c.Uppfoljning = &UppfoljningService{client: c, contentType: "uppfoljning-feed"}

	c.Mura = &studentinformation.Client{
		HttpClient: c.httpClient
	}

	return c, nil
}
