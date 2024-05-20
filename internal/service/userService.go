package service

import (
	"context"
	"intership/internal/sqlcdb"
)

type UserService struct {
	*sqlcdb.Queries
}

func NewUserService(q *sqlcdb.Queries) *UserService {
	return &UserService{q}
}

func (u *UserService) GetAllUsers(ctx context.Context) ([]sqlcdb.User, error) {
	return u.ListUsers(ctx)
}

func (u *UserService) CreateUser(ctx context.Context, user sqlcdb.User) (sqlcdb.User, error) {
	return u.CreateUser(ctx, user)
}
