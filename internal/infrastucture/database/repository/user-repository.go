package repository

import (
	"database/sql"

	"github.com/mauFade/journaly/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(u *domain.UserModel) error {
	stmt, err := r.db.Prepare(`INSERT INTO users (id, name, email, password, phone, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		u.ID,
		u.Name,
		u.Email,
		u.Password,
		u.Phone,
		u.CreatedAt,
		u.UpdatedAt,
	)

	return err
}

func (r *UserRepository) FindByEmail(e string) *domain.UserModel {
	return nil
}
func (r *UserRepository) FindByPhone(p string) *domain.UserModel {
	return nil
}
