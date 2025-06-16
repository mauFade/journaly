package repository

import (
	"database/sql"
	"time"

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
	stmt, err := r.db.Prepare(`INSERT INTO journals (id, title, content, tags, word_count, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		j.ID,
		j.Title,
		j.Content,
		j.Tags,
		j.WordCount,
		j.CreatedAt,
		j.UpdatedAt,
	)

	return err
}

func (r *JournalRepository) GetByID(id string) (*domain.JournalModel, error) {
	var j domain.JournalModel
	var createdAt, updatedAt time.Time

	err := r.db.QueryRow(`
	SELECT id, title, content, tags, word_count, created_at, updated_at
	FROM journals WHERE id = $1`, id,
	).Scan(
		&j.ID,
		&j.Title,
		&j.Content,
		&j.Tags,
		&j.WordCount,
		&createdAt,
		&updatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, domain.ErrJournalNotFound
	}
	if err != nil {
		return nil, err
	}
	j.CreatedAt = createdAt
	j.UpdatedAt = updatedAt

	return &j, nil
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
