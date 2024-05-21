package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"intership/internal/handlers"
	"intership/internal/router"
	"intership/internal/service"
	"intership/internal/sqlcdb"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://postgres:password@localhost:5432/internship")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close(ctx)

	db := sqlcdb.New(conn)

	//err = db.CreateUser(ctx, sqlcdb.CreateUserParams{
	//	Name:     "Todor",
	//	Surname:  "Popovic",
	//	Email:    "todor.popovic@nimbus-tech.io",
	//	Password: "password",
	//})

	iservice := service.NewUserService(db)
	handler := handlers.NewUserHandler(iservice)
	r := router.NewRouter(handler)
	r.HandleRequests()
	go func() {
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			fmt.Println(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	fmt.Println("Graceful shutdown", sig)

}
