package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, "postgresql://postgres:tajna@db:5432/internship")
}
