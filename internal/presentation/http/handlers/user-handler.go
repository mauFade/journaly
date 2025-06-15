package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mauFade/journaly/internal/application/dto"
	userservice "github.com/mauFade/journaly/internal/application/service/user-service"
)

type UserHandler struct {
	userService *userservice.UserService
}

func NewUserHandler(crs *userservice.UserService) *UserHandler {
	return &UserHandler{
		userService: crs,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	res, err := h.userService.CreateUser(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	var input dto.AuthUserRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	res, err := h.userService.Authenticate(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
