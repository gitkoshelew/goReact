package auth

import (
	"context"
	"fmt"
	"gateway/internal/apperror"
	"gateway/internal/client"
	"gateway/pkg/response"
	"io"
)

// Logout ...
func Logout(ctx context.Context, c *client.Client, body io.Reader) (*response.Service, error) {
	url, err := c.Base.BuildURL(c.Resource, nil)
	if err != nil {
		return nil, err
	}

	req, err := c.Base.CreateRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("%v", ctx.Value(client.AccessTokenCtxKey)))

	resp, err := c.Base.SendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.IsOk {
		resp, err := c.Base.ReadResponse(resp, []string{"Access-Token"})
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	c.Base.Logger.Errorf("%v", apperror.APIError(resp.Error.ErrorCode, resp.Error.Message, resp.Error.DeveloperMessage))

	return nil, apperror.APIError(resp.Error.ErrorCode, resp.Error.Message, resp.Error.DeveloperMessage)

}
