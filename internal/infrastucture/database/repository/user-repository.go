package repository

import (
	"database/sql"
	"time"

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
	var u domain.UserModel
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, password, phone, created_at, updated_at
		FROM users WHERE email = $1
	`, e).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Phone,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil
	}

	u.CreatedAt = createdAt
	u.UpdatedAt = updatedAt
	return &u
}

func (r *UserRepository) FindByPhone(p string) *domain.UserModel {
	var u domain.UserModel
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
		SELECT id, name, email, password, phone, created_at, updated_at
		FROM users WHERE phone = $1
	`, p).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Phone,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil
	}

	u.CreatedAt = createdAt
	u.UpdatedAt = updatedAt
	return &u
}
