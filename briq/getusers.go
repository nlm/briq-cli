package briq

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type ListUsersRequest struct {
}

type GetUserRequest struct {
	Username string `json:"username,omitempty"`
}

type User struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Username    string    `json:"username,omitempty"`
	DisplayName string    `json:"displayName,omitempty"`
}

type ListUsersResponse struct {
	Users []User `json:"users,omitempty"`
}

func (client *Client) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	resp := make([]User, 0)
	if err := client.do(ctx, http.MethodGet, UrlUsers, req, &resp); err != nil {
		return nil, err
	}
	return &ListUsersResponse{Users: resp}, nil
}

var cachedGetUsersResponse *ListUsersResponse

func (client *Client) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	if cachedGetUsersResponse == nil {
		res, err := client.ListUsers(ctx, &ListUsersRequest{})
		if err != nil {
			return nil, err
		}
		cachedGetUsersResponse = res
	}
	for _, user := range cachedGetUsersResponse.Users {
		if user.Username == req.Username {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("User not found: %w", ErrNotFound)
}
