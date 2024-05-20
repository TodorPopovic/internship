package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"intership/internal/service"
	"net/http"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (handler *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte("Get user with id " + vars["id"]))
}

func (handler *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := handler.service.GetAllUsers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	fmt.Println(users)
}

func (handler *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
}

func (handler *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte("Update user with id " + vars["id"]))
}
func (handler *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte("Delete user with id " + vars["id"]))
}
