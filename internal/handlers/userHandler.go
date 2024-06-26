package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	_ "github.com/swaggo/http-swagger"
	"intership/internal/service"
	"intership/internal/sqlcdb"
	"net/http"
	"os"
	"strconv"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.WarnLevel)
}

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{userService}
}

func (handler *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error("Unable to parse id")
		json.NewEncoder(w).Encode("Unable to parse id")
	}
	user, err2 := handler.service.GetUser(r.Context(), int64(id))
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error("Unable to retrieve user")
		json.NewEncoder(w).Encode("Unable to retrieve user")
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func (handler *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := handler.service.GetAllUsers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error("Unable to retrieve users")
		json.NewEncoder(w).Encode("Unable to retrieve users")
	} else {
		json.NewEncoder(w).Encode(users)
	}
}

// CreateUserHandler
// @Summary Create a new user
// @Description Create a new user
// @Accept  json
// @Param user body sqlcdb.CreateUserParams true "Create User"
// @Success 200 string "User created successfully"
// @Router /users [post]
func (handler *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var createUser sqlcdb.CreateUserParams
	err := decoder.Decode(&createUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error("Invalid json")
		json.NewEncoder(w).Encode("Invalid json")
	}
	err2 := handler.service.CreateUser(r.Context(), createUser)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error("Unable to create user")
		json.NewEncoder(w).Encode("Unable to create user")
	} else {
		json.NewEncoder(w).Encode("User created successfully")
	}
}

func (handler *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error("Unable to parse id")
		json.NewEncoder(w).Encode("Unable to parse id")
	}
	var updateUser sqlcdb.UpdateUserParams
	decoder := json.NewDecoder(r.Body)
	err2 := decoder.Decode(&updateUser)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error("Invalid json")
		json.NewEncoder(w).Encode("Invalid json")
	}
	updateUser.ID = int64(id)
	err3 := handler.service.UpdateUser(r.Context(), updateUser)
	if err3 != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error("Unable to update user")
		json.NewEncoder(w).Encode("Unable to update user")
	} else {
		json.NewEncoder(w).Encode("User updated successfully")
	}
}

func (handler *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error("Unable to parse id")
		json.NewEncoder(w).Encode("Unable to parse id")
	}
	err1 := handler.service.DeleteUser(r.Context(), int64(id))
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Error("Unable to delete user")
		json.NewEncoder(w).Encode("Unable to delete user")
	} else {
		json.NewEncoder(w).Encode("User deleted successfully")
	}
}
