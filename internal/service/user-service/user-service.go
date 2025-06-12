package userservice

import "github.com/mauFade/journaly/internal/domain"

type UserService struct {
	repository domain.UserRepository
}

func NewUserService(r domain.UserRepository) *UserService {
	return &UserService{
		repository: r,
	}
}
