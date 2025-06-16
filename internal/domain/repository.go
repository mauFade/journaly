package domain

type UserRepository interface {
	Save(u *UserModel) error
	FindByEmail(e string) *UserModel
	FindByPhone(p string) *UserModel
}

type JournalRepository interface {
	Save(j *JournalModel) error
	GetByID(id string) (*JournalModel, error)
	GetByUser(userId string) ([]*JournalModel, error)
	Update(j *JournalModel) error
	Delete(id string) error
}
