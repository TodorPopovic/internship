//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"intership/internal/app"
)

func InitializeEvent() *app.App {
	panic(wire.Build(app.Providers))
}
