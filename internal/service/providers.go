package service

import (
	"context"
	"github.com/google/wire"
	"intership/internal/db"
	"intership/internal/handlers"
	"intership/internal/router"
)

var Providers = wire.NewSet(
	context.Background,
	db.NewPostgres,
	NewUserService,
	wire.Bind(new(IUserService), new(*UserService)),
	handlers.NewUserHandler,
	router.NewRouter,
)
