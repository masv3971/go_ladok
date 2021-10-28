package kataloginformation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/masv3971/goladok3"
)

// Client handles kataloginformation
type Client struct {
	contentType string
	client      *goladok3.Client
}

// GetAnvandareAutentiserad gets kataloginformation/anvandare/autentiserad
func (c *Client) GetAnvandareAutentiserad(ctx context.Context) (*GetAnvandareAutentiseradReply, *http.Response, error) {
	req, err := c.client.HTTPClient.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s", "kataloginformation/anvandare/autentiserad"),
		ladokAcceptHeader[c.contentType][c.client.format],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GetAnvandareAutentiseradReply{}
	resp, err := c.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// GetBehorighetsprofilerCfg configuration for GetBehorighetsprofil
type GetBehorighetsprofilerCfg struct {
	UID string `validate:"required,uuid"`
}

// GetBehorighetsprofil return structure of rights for uid
func (c *Client) GetBehorighetsprofil(ctx context.Context, cfg *GetBehorighetsprofilerCfg) (*GetBehorighetsprofilReply, *http.Response, error) {
	req, err := c.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s/%s", "kataloginformation/behorighetsprofil", cfg.UID),
		ladokAcceptHeader[s.contentType][s.client.format],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GetBehorighetsprofilReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}

// GetAnvandarbehorighetEgna return structure of ladok permission
func (c *Client) GetAnvandarbehorighetEgna(ctx context.Context) (*GetAnvandarbehorighetEgnaReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s", "kataloginformation/anvandarbehorighet/egna"),
		ladokAcceptHeader[s.contentType][s.client.format],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	reply := &GetAnvandarbehorighetEgnaReply{}
	resp, err := s.client.do(req, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}
