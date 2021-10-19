package goladok3

import "errors"

// Permissions is a simplfy permissions object
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

var (
	// ErrNoValidContentType if no valid content-type is found
	ErrNoValidContentType = errors.New("No valid content-type found")
	// ErrNoEnvFound if no valid environment is found in certificate (ou)
	ErrNoEnvFound = errors.New("No valid environment (ou) found")
	// ErrNotSufficentPermissions if not all provided permissions are met
	ErrNotSufficentPermissions = errors.New("Not sufficent permissions")
)

const (
	contentTypeStudiedeltagandeJSON   = "application/vnd.ladok-studiedeltagande+json;charset=UTF-8"
	contentTypeKataloginformationJSON = "application/vnd.ladok-kataloginformation+json;charset=UTF-8"
	contentTypeStudentinformationJSON = "application/vnd.ladok-studentinformation+json;charset=UTF-8"
	contentTypeAtomXML                = "application/atom+xml;charset=UTF-8"

	envIntTestAPI = "Int-test-API"
	envProdAPI    = "Prod-API"
	envTestAPI    = "Test-API"
)
