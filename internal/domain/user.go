package domain

import "golang.org/x/crypto/bcrypt"

type UserModel struct {
	ID       string
	Name     string
	Email    string
	Password string
	Phone    string
}

func NewUserModel(id, name, email, pass, phone string) *UserModel {
	return &UserModel{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: pass,
		Phone:    phone,
	}
}

func (u *UserModel) GenerateHashPassword(pass string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 6)

	if err != nil {
		return pass
	}

	return string(hash)
}
