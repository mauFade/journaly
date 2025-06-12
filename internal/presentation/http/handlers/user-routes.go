package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mauFade/journaly/internal/dto"
	userservice "github.com/mauFade/journaly/internal/service/user-service"
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
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.userService.CreateUser(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
