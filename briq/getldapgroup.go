package briq

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type GetGroupRequest struct {
	Name string `json:"name,omitempty"`
}

type Group struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Platform    string    `json:"platform,omitempty"`
	ExternalRef string    `json:"externalRef,omitempty"`
	CreatedAt   string    `json:"created_at,omitempty"`
	Users       []User    `json:"users,omitempty"`
}

var cachedGetGroupsResponse *ListGroupsResponse

func (client *Client) GetGroup(ctx context.Context, req *GetGroupRequest) (*Group, error) {
	if cachedGetGroupsResponse == nil {
		res, err := client.ListGroups(ctx, &ListGroupsRequest{})
		if err != nil {
			return nil, err
		}
		cachedGetGroupsResponse = res
	}
	for _, group := range cachedGetGroupsResponse.Groups {
		if group.Name == req.Name {
			return &group, nil
		}
	}
	return nil, fmt.Errorf("Group not found: %w", ErrNotFound)
}
