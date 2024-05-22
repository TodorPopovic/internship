package app

import (
	"github.com/google/wire"
	"intership/internal/router"
	"intership/internal/service"
)

type App struct {
	router *router.MyRouter
}

var Providers = wire.NewSet(
	service.Providers,
	wire.Struct(new(App), "*"),
)
