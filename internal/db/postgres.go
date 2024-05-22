package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, "postgresql://postgres:password@localhost:5432/internship")
}
