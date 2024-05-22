package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func NewPostgres(ctx context.Context) (*pgxpool.Pool, error) {
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Error loading .env file")
	}
	return pgxpool.New(ctx, os.Getenv("DB_URI"))
}
