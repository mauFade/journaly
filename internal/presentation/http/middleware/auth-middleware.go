package middleware

import (
	"encoding/json"
	"net/http"

	userservice "github.com/mauFade/journaly/internal/application/service/user-service"
)

type AuthMiddleware struct {
	s *userservice.UserService
}

func NewAuthMiddleware(s *userservice.UserService) *AuthMiddleware {
	return &AuthMiddleware{
		s: s,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")

		if apiKey == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "api key is required for this endpoint"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
