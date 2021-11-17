package goladok3

import (
	"context"
	"fmt"
	"net/http"

	"github.com/masv3971/goladok3/ladoktypes"
)

// studentinformationService handles studentinformation
type studentinformationService struct {
	client  *Client
	service string
}

func (s *studentinformationService) acceptHeader() string {
	return ladokAcceptHeader[s.service][s.client.format]
}

// GetStudentReq config for GetStudent
type GetStudentReq struct {
	UID          string `validate:"required_without=Personnummer ExterntUID,uuid"`
	ExterntUID   string `validate:"required_without=Personnummer UID, uuid"`
	Personnummer string `validate:"required_without=UID ExterntUID"`
}

// GetStudent return student
func (s *studentinformationService) GetStudent(ctx context.Context, req *GetStudentReq) (*ladoktypes.Student, *http.Response, error) {
	reply := &ladoktypes.Student{}
	var url string

	if req.UID != "" {
		url = fmt.Sprintf("%s/%s/%s", s.service, "student", req.UID)
	} else if req.Personnummer != "" {
		url = fmt.Sprintf("%s/%s/%s/%s", s.service, "student", "personnummer", req.Personnummer)
	} else if req.ExterntUID != "" {
		url = fmt.Sprintf("%s/%s/%s/%s", s.service, "student", "externtuuid", req.ExterntUID)
	}

	resp, err := s.client.call(ctx, s.acceptHeader(), "GET", url, nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// GetAktivPaLarosateReq config for GetAktivPaLarosate
type GetAktivPaLarosateReq struct {
	UID string `validate:"required,uuid"`
}

func (s *studentinformationService) GetAktivPaLarosate(ctx context.Context, req *GetAktivPaLarosateReq) (*ladoktypes.AktivPaLarosate, *http.Response, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", s.service, "student", req.UID, "aktivpalarosaten")
	reply := &ladoktypes.AktivPaLarosate{}
	resp, err := s.client.call(ctx, s.acceptHeader(), "GET", url, nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}
