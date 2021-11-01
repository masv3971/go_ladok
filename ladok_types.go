package goladok3

// Permissions is a simplify permissions object
type Permissions map[int64]string

// Link is a general ladok link structure
type Link struct {
	Method    string `json:"method"`
	URI       string `json:"uri"`
	MediaType string `json:"mediaType"`
	Rel       string `json:"rel"`
}

// Benamning is a general ladok Benamning structure
type Benamning struct {
	Sprakkod string     `json:"Sprakkod"`
	Text     string     `json:"Text"`
	Link     []struct{} `json:"link"`
}

type (
	// FeedID ladok id is an int
	FeedID string
)

type serviceTypes struct {
	client       *Client
	service      string
	acceptHeader string
}

const (
	envIntTestAPI = "Int-test-API"
	envProdAPI    = "Prod-API"
	envTestAPI    = "Test-API"
)

var (
	contentTypeStudiedeltagandeJSON   = "application/vnd.ladok-studiedeltagande+json;charset=UTF-8"
	contentTypeKataloginformationJSON = "application/vnd.ladok-kataloginformation+json;charset=UTF-8"
	contentTypeStudentinformationJSON = "application/vnd.ladok-studentinformation+json;charset=UTF-8"
	contentTypeAtomXML                = "application/atom+xml;charset=UTF-8"
)
