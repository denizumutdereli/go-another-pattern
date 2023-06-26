package account

import "context"

type Service interface {
	Create(ctx context.Context, email, password string) (string, error)
	Get(ctx context.Context, id string) (string, error)
}
