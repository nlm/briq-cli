package briq

import (
	"context"
	"net/http"
)

type ListGroupsRequest struct {
}

type ListGroupsResponse struct {
	Groups []Group `json:"groups,omitempty"`
}

func (client *Client) ListGroups(ctx context.Context, req *ListGroupsRequest) (*ListGroupsResponse, error) {
	resp := make([]Group, 0)
	if err := client.do(ctx, http.MethodGet, UrlGroups, req, &resp); err != nil {
		return nil, err
	}
	return &ListGroupsResponse{Groups: resp}, nil
}
