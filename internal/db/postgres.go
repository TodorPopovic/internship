package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func NewPostgres(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, os.Getenv("DB_URI"))
}
