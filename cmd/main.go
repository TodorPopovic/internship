package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

// @title Internship API
// @version 1.0
// @description This is the API documentation for the Internship project
// @host localhost:8080
// @BasePath /
func main() {
	if app, err := InitApp(); err == nil {
		app.Router.HandleRequests()
		//scheduler := app.Scheduler.Start()
		//<-scheduler

		err = http.ListenAndServe(":8080", app.Router)
		if err != nil {
			logrus.Fatal(err)
		}

	}
}
