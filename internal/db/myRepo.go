package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"intership/internal/sqlcdb"
)

type MyRepo struct {
	*pgxpool.Pool
}

type IPostgresRepo interface {
	GetUser(ctx context.Context, id int64) (sqlcdb.User, error)
	GetAllUsers(ctx context.Context) ([]sqlcdb.User, error)
	CreateUser(ctx context.Context, user sqlcdb.CreateUserParams) error
	UpdateUser(ctx context.Context, user sqlcdb.UpdateUserParams) error
	DeleteUser(ctx context.Context, id int64) error
}

func NewMyRepo(pool *pgxpool.Pool) *MyRepo {
	return &MyRepo{pool}
}

func (r *MyRepo) GetUser(ctx context.Context, id int64) (sqlcdb.User, error) {
	q := sqlcdb.New(r)
	return q.GetUser(ctx, id)
}

func (r *MyRepo) GetAllUsers(ctx context.Context) ([]sqlcdb.User, error) {
	q := sqlcdb.New(r)
	return q.ListUsers(ctx)
}

func (r *MyRepo) CreateUser(ctx context.Context, user sqlcdb.CreateUserParams) error {
	q := sqlcdb.New(r)
	return q.CreateUser(ctx, user)
}

func (r *MyRepo) UpdateUser(ctx context.Context, user sqlcdb.UpdateUserParams) error {
	q := sqlcdb.New(r)
	return q.UpdateUser(ctx, user)
}

func (r *MyRepo) DeleteUser(ctx context.Context, id int64) error {
	q := sqlcdb.New(r)
	return q.DeleteUser(ctx, id)
}
