package studentinformation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/masv3971/goladok3"
	"github.com/masv3971/goladok3/internal/httpclient"
	"github.com/masv3971/goladok3/internal/validate"
)

// Client handles studentinformation
type Client struct {
	contentType string
	client      *goladok3.Client
}

// GetStudentCfg config for GetStudent
type GetStudentCfg struct {
	UID string `validate:"required,uuid"`
}

// GetStudent return student
func (c *Client) GetStudent(ctx context.Context, cfg *GetStudentCfg) (*GetStudentReply, *http.Response, error) {
	if err := validate.Check(cfg); err != nil {
		return nil, nil, err
	}

	req, err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s/%s", "studentinformation/student", cfg.UID),
		httpclient.LadokAcceptHeader[c.contentType][c.httpClient.Format],
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	reply := &GetStudentReply{}
	resp, err := c.httpClient.Do(req, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
