package httpclient

import "errors"

var (
	contentTypeStudiedeltagandeJSON   = "application/vnd.ladok-studiedeltagande+json;charset=UTF-8"
	contentTypeKataloginformationJSON = "application/vnd.ladok-kataloginformation+json;charset=UTF-8"
	contentTypeStudentinformationJSON = "application/vnd.ladok-studentinformation+json;charset=UTF-8"
	contentTypeAtomXML                = "application/atom+xml;charset=UTF-8"
)

var (
	// ErrNoValidContentType error no valid content type
	ErrNoValidContentType = errors.New("No valid content type")
)
