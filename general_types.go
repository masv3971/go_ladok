package goladok3

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
