//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"intership/internal/app"
)

func InitApp() (*app.App, error) {
	panic(wire.Build(app.Providers))
}
