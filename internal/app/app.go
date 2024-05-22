package app

import (
	"context"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"intership/internal/db"
	"intership/internal/handlers"
	"intership/internal/router"
	"intership/internal/service"
)

type App struct {
	context.Context
	*pgxpool.Pool
	Router *router.MyRouter
}

var Providers = wire.NewSet(
	context.Background,
	db.NewPostgres,
	service.Providers,
	handlers.NewUserHandler,
	router.NewRouter,
	wire.Struct(new(App), "*"),
)
