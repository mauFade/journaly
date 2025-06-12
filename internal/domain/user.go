package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserModel(id, name, email, pass, phone string, createdAt, updatedAt time.Time) *UserModel {
	return &UserModel{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  pass,
		Phone:     phone,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func (u *UserModel) GenerateHashPassword(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 6)

	if err != nil {
		return pass
	}

	return string(hash)
}
