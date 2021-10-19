package goladok3

import (
	"context"
	"fmt"
	"net/http"
)

// KataloginformationService handles kataloginformation
type KataloginformationService struct {
	client      *Client
	contentType string
}

// GetAnvandareAutentiserad gets kataloginformation/anvandare/autentiserad
func (s *KataloginformationService) GetAnvandareAutentiserad(ctx context.Context) (*GetAnvandareAutentiseradReply, *http.Response, error) {
	req, err := s.client.newRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s", "kataloginformation/anvandare/autentiserad"),
		ladokAcceptHeader[s.contentType][s.client.format],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GetAnvandareAutentiseradReply{}
	resp, err := s.client.do(req, reply)
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
func (s *KataloginformationService) GetBehorighetsprofil(ctx context.Context, cfg *GetBehorighetsprofilerCfg) (*GetBehorighetsprofilReply, *http.Response, error) {
	req, err := s.client.newRequest(
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
func (s *KataloginformationService) GetAnvandarbehorighetEgna(ctx context.Context) (*GetAnvandarbehorighetEgnaReply, *http.Response, error) {
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
