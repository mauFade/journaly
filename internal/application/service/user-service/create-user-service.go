package userservice

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mauFade/journaly/internal/application/dto"
	"github.com/mauFade/journaly/internal/domain"
)

func (s *UserService) CreateUser(req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	u := s.repository.FindByEmail(req.Email)

	fmt.Println(u)
	if u != nil {
		return nil, domain.ErrEmailAlreadyExists
	}

	u = s.repository.FindByPhone(req.Phone)

	if u != nil {
		return nil, domain.ErrPhoneAlreadyExists
	}

	u = domain.NewUserModel(uuid.NewString(), req.Name, req.Email, u.GenerateHashPassword(req.Password), req.Phone, time.Now(), time.Now())
	err := s.repository.Save(u)

	if err != nil {
		return nil, err
	}

	return dto.ToDto(u), nil
}
