package service

import (
	"context"
	"intership/internal/sqlc/repo"
)

type UserService struct {
	*repo.Queries
}

func NewUserService(q *repo.Queries) *UserService {
	return &UserService{q}
}

func (u *UserService) GetAllUsers(ctx context.Context) ([]repo.User, error) {
	return u.ListUsers(ctx)
}
