package service

import (
	"context"
	"intership/internal/sqlcdb"
)

type IUserService interface {
	GetUser(ctx context.Context, id int64) (sqlcdb.User, error)
	GetAllUsers(ctx context.Context) ([]sqlcdb.User, error)
	CreateUser(ctx context.Context, user sqlcdb.CreateUserParams) error
	UpdateUser(ctx context.Context, user sqlcdb.UpdateUserParams) error
	DeleteUser(ctx context.Context, id int64) error
}

type UserService struct {
	q *sqlcdb.Queries
}

func NewUserService(q *sqlcdb.Queries) *UserService {
	return &UserService{q}
}

func (u *UserService) GetUser(ctx context.Context, id int64) (sqlcdb.User, error) {
	return u.q.GetUser(ctx, id)
}

func (u *UserService) GetAllUsers(ctx context.Context) ([]sqlcdb.User, error) {
	return u.q.ListUsers(ctx)
}

func (u *UserService) CreateUser(ctx context.Context, user sqlcdb.CreateUserParams) error {
	return u.q.CreateUser(ctx, user)
}

func (u *UserService) UpdateUser(ctx context.Context, user sqlcdb.UpdateUserParams) error {
	userToUpdate, _ := u.q.GetUser(ctx, user.ID)
	if user.Name != "" {
		userToUpdate.Name = user.Name
	}
	if user.Surname != "" {
		userToUpdate.Surname = user.Surname
	}
	if user.Email != "" {
		userToUpdate.Email = user.Email
	}
	if user.Password != "" {
		userToUpdate.Password = user.Password
	}
	return u.q.UpdateUser(ctx, sqlcdb.UpdateUserParams(userToUpdate))
}

func (u *UserService) DeleteUser(ctx context.Context, id int64) error {
	return u.q.DeleteUser(ctx, id)
}
