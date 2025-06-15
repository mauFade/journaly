package userservice

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mauFade/journaly/internal/application/dto"
	"github.com/mauFade/journaly/internal/domain"
)

type authResponse struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

func (s *UserService) Authenticate(data *dto.AuthUserRequest) (*authResponse, error) {
	u := s.repository.FindByEmail(data.Email)

	if u == nil {
		return nil, domain.ErrAccountNotFoundByEmail
	}

	if err := u.ComparePasswords(data.Password); err != nil {
		return nil, domain.ErrInvalidPassword
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return nil, err
	}

	return &authResponse{
		Name:  u.Name,
		Token: tokenString,
	}, nil
}
