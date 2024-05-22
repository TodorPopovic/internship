package main

import (
	"github.com/google/wire"
	"intership/internal/app"
)

func initApp() *app.App {
	panic(wire.Build(app.Providers))
}
