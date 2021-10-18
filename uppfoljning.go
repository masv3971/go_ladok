package goladok3

import (
	"context"
	"fmt"
	"net/http"
)

// UppfoljningService holds uppfoljning feed object
type UppfoljningService struct {
	client      *Client
	contentType string
}

// FeedRecent atom feed /uppfoljning/feed/recent
func (s *UppfoljningService) FeedRecent(ctx context.Context) (*FeedRecent, *http.Response, error) {
	env, err := s.client.environment()
	if err != nil {
		// TODO(masv): handle error
	}

	var url string
	switch env {
	case envIntTestAPI:
		url = "/handelse/feed/recent"
	default:
		url = "/uppfoljning/feed/recent"
	}

	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s", url),
		ladokAcceptHeader[s.contentType]["xml"],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &FeedRecent{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
