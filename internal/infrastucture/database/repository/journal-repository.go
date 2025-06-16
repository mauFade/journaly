package repository

import (
	"database/sql"

	"github.com/mauFade/journaly/internal/domain"
)

type JournalRepository struct {
	db *sql.DB
}

func NewJournalRepository(d *sql.DB) *JournalRepository {
	return &JournalRepository{
		db: d,
	}
}

func (r *JournalRepository) Save(j *domain.JournalModel) error {
	return nil
}
func (r *JournalRepository) GetByID(id string) (*domain.JournalModel, error) {
	return nil, nil
}
func (r *JournalRepository) GetByUser(userId string) ([]*domain.JournalModel, error) {
	return nil, nil
}
func (r *JournalRepository) Update(j *domain.JournalModel) error {
	return nil
}
func (r *JournalRepository) Delete(id string) error {
	return nil
}
