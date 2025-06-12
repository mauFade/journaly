package userservice

import (
	"github.com/google/uuid"
	"github.com/mauFade/journaly/internal/domain"
	"github.com/mauFade/journaly/internal/dto"
)

func (s *UserService) CreateUser(req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	u := s.repository.FindByEmail(req.Email)

	if u != nil {
		return nil, domain.ErrEmailAlreadyExists
	}

	u = s.repository.FindByPhone(req.Phone)

	if u != nil {
		return nil, domain.ErrPhoneAlreadyExists
	}

	u = domain.NewUserModel(uuid.NewString(), req.Name, req.Email, u.GenerateHashPassword(req.Password), req.Phone)
	err := s.repository.Create(u)

	if err != nil {
		return nil, err
	}

	return dto.ToDto(u), nil
}
