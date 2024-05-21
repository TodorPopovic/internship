package app

import (
	"intership/internal/router"
)

type App struct {
	router *router.MyRouter
}

func NewApp(router *router.MyRouter) *App {
	return &App{router: router}
}
