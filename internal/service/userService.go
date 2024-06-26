package service

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
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
	*pgxpool.Pool
}

func NewUserService(p *pgxpool.Pool) *UserService {
	return &UserService{p}
}

func (u *UserService) GetUser(ctx context.Context, id int64) (sqlcdb.User, error) {
	q := sqlcdb.New(u)
	return q.GetUser(ctx, id)
}

func (u *UserService) GetAllUsers(ctx context.Context) ([]sqlcdb.User, error) {
	q := sqlcdb.New(u)
	return q.ListUsers(ctx)
}

func (u *UserService) CreateUser(ctx context.Context, user sqlcdb.CreateUserParams) error {
	q := sqlcdb.New(u)
	return q.CreateUser(ctx, user)
}

func (u *UserService) UpdateUser(ctx context.Context, user sqlcdb.UpdateUserParams) error {
	q := sqlcdb.New(u)
	userToUpdate, _ := q.GetUser(ctx, user.ID)
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
	return q.UpdateUser(ctx, sqlcdb.UpdateUserParams(userToUpdate))
}

func (u *UserService) DeleteUser(ctx context.Context, id int64) error {
	q := sqlcdb.New(u)
	return q.DeleteUser(ctx, id)
}
