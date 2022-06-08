package briq

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type GetMeRequest struct {
}

type GetMeResponse struct {
	Id              uuid.UUID `json:"id,omitempty"`
	Username        string    `json:"username,omitempty"`
	DisplayName     string    `json:"displayName,omitempty"`
	Email           string    `json:"email,omitempty"`
	FirstName       string    `json:"firstName,omitempty"`
	LastName        string    `json:"lastName,omitempty"`
	ExternalRef     string    `json:"externalRef,omitempty"`
	Image           string    `json:"image,omitempty"`
	Role            string    `json:"role,omitempty"`
	ActiveBalance   int       `json:"activeBalance,omitempty"`
	InactiveBalance int       `json:"inactiveBalance,omitempty"`
}

func (client *Client) GetMe(ctx context.Context, req *GetMeRequest) (*GetMeResponse, error) {
	res := &GetMeResponse{}
	if err := client.do(ctx, http.MethodGet, UrlMe, req, res); err != nil {
		return nil, err
	}
	return res, nil
}
