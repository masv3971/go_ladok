package goladok3

import "errors"

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

var (
	// ErrNoValidContentType if no valid content-type is found
	ErrNoValidContentType = errors.New("No valid content-type found")
	// ErrNoEnvFound if no valid environment is found in certificate (ou)
	ErrNoEnvFound = errors.New("No valid environment (ou) found")
	// ErrNotSufficientPermissions if not all provided permissions are met
	ErrNotSufficientPermissions = errors.New("Not sufficient permissions")
)

const (
	envIntTestAPI = "Int-test-API"
	envProdAPI    = "Prod-API"
	envTestAPI    = "Test-API"
)
