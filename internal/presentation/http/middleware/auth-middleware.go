package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
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

type ctxKey string

const UserKey ctxKey = "userID"

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")

		w.Header().Set("Content-Type", "application/json")

		if apiKey == "" {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "api key is required for this endpoint"})
			return
		}

		token, err := jwt.Parse(apiKey, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})

			return
		}

		claims, tokenClaimsOk := token.Claims.(jwt.MapClaims)

		if tokenClaimsOk && token.Valid {
			userID, ok := claims["userID"].(string)

			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"message": "Invalid userID type in token"})

				return
			}

			ctx := context.WithValue(r.Context(), UserKey, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Invalid token claims"})

			return
		}
	})
}
