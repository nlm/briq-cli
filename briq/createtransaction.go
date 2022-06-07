package briq

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CreateTransactionRequest struct {
	App     string    `json:"app,omitempty"`
	Comment string    `json:"comment,omitempty"`
	To      uuid.UUID `json:"to,omitempty"`
}

type CreateTransactionResponse struct {
	Amount      int        `json:"amount,omitempty"`
	App         string     `json:"app,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	From        string     `json:"from,omitempty"`
	Id          uuid.UUID  `json:"id,omitempty"`
	IsAnonymous bool       `json:"isAnonymous,omitempty"`
	ReactedAt   *time.Time `json:"reactedAt,omitempty"`
	Reaction    *string    `json:"reaction,omitempty"`
	To          string     `json:"to,omitempty"`
	UserFromId  *uuid.UUID `json:"user_from_id,omitempty"`
	UserToId    *uuid.UUID `json:"user_to_id,omitempty"`
}

// CreateTransaction creates a transaction, which is the way to give briqs.
func (client *Client) CreateTransaction(ctx context.Context, req *CreateTransactionRequest) (*CreateTransactionResponse, error) {
	resp := &CreateTransactionResponse{}
	if err := client.do(ctx, http.MethodPost, UrlTransactions, req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
