package domain

type UserRepository interface {
	Create(u *UserModel) error
	FindByEmail(e string) *UserModel
	FindByPhone(p string) *UserModel
}
