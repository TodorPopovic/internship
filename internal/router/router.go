package router

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"intership/internal/handlers"
)

type MyRouter struct {
	*mux.Router
	*handlers.UserHandler
}

func NewRouter(handler *handlers.UserHandler) *MyRouter {
	r := mux.NewRouter()
	return &MyRouter{r, handler}
}

func (r *MyRouter) HandleRequests() {
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	r.HandleFunc("/users/{id}", r.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", r.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", r.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", r.UpdateUserHandler).Methods("PATCH")
	r.HandleFunc("/users/{id}", r.DeleteUserHandler).Methods("DELETE")
}
