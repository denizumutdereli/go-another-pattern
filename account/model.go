package account

import "context"

type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Repository interface {
	Create(ctx context.Context, user User) error
	Get(ctx context.Context, id string) (string, error)
}
