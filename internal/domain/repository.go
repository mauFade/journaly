package domain

type UserRepository interface {
	Save(u *UserModel) error
	FindByEmail(e string) *UserModel
	FindByPhone(p string) *UserModel
}
