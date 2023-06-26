package account

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) Create(ctx context.Context, email, password string) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user := User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	if err := s.repository.Create(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id)

	return "Success", nil
}

func (s *service) Get(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "Get")

	email, err := s.repository.Get(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
	}

	logger.Log("Get User", id)

	return email, nil
}
