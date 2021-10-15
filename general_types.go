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

var (
	// ErrNoValidContentType if no valid content-type is found
	ErrNoValidContentType = errors.New("No valid content-type found")
)

const (
	// ContentTypeStudiedeltagandeJSON ladok specefic type
	ContentTypeStudiedeltagandeJSON = "application/vnd.ladok-studiedeltagande+json;charset=UTF-8"
	// ContentTypeKataloginformationJSON ladok specefic type
	ContentTypeKataloginformationJSON = "application/vnd.ladok-kataloginformation+json;charset=UTF-8"
	// ContentTypeStudentinformationJSON ladok specefic type
	ContentTypeStudentinformationJSON = "application/vnd.ladok-studentinformation+json;charset=UTF-8"
	// ContentTypeAtomXML ladok specefic type
	ContentTypeAtomXML = "application/atom+xml;charset=UTF-8"
)
