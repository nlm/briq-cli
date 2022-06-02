package briq

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type GetUsersRequest struct {
}

type User struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Username    string    `json:"username,omitempty"`
	DisplayName string    `json:"displayName,omitempty"`
}

type GetUsersResponse struct {
	Users []User
}

func (client *Client) GetUsers(ctx context.Context, req *GetUsersRequest) (*GetUsersResponse, error) {
	resp := make([]User, 0)
	if err := client.do(ctx, http.MethodGet, UrlUsers, req, &resp); err != nil {
		return nil, err
	}
	return &GetUsersResponse{Users: resp}, nil
}
