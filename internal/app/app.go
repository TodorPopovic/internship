package app

import (
	"intership/internal/handlers"
	"intership/internal/router"
	"intership/internal/service"
)

type App struct {
	router  *router.MyRouter
	handler *handlers.UserHandler
	service *service.UserService
}
