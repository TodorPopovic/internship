// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"intership/internal/app"
	"intership/internal/db"
	"intership/internal/handlers"
	"intership/internal/router"
	"intership/internal/service"
)

// Injectors from init.go:

func InitApp() (*app.App, error) {
	contextContext := context.Background()
	pool, err := db.NewPostgres(contextContext)
	if err != nil {
		return nil, err
	}
	scheduler := service.NewScheduler()
	myRepo := db.NewMyRepo(pool)
	userService := service.NewUserService(myRepo)
	userHandler := handlers.NewUserHandler(userService)
	myRouter := router.NewRouter(userHandler)
	appApp := &app.App{
		Context:   contextContext,
		Pool:      pool,
		Scheduler: scheduler,
		Router:    myRouter,
	}
	return appApp, nil
}
