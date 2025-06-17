package repository

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
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
	stmt, err := r.db.Prepare(`INSERT INTO journals (id, user_id, title, content, tags, word_count, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		j.ID,
		j.UserID,
		j.Title,
		j.Content,
		pq.StringArray(j.Tags),
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
	SELECT id, user_id, title, content, tags, word_count, created_at, updated_at
	FROM journals WHERE id = $1`, id,
	).Scan(
		&j.ID,
		&j.UserID,
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
	journals := []*domain.JournalModel{}

	rows, err := r.db.Query(`
	SELECT id, user_id, title, content, tags, word_count, created_at, updated_at
	FROM journals 
	WHERE user_id = $1`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var j domain.JournalModel
		var tags pq.StringArray

		err := rows.Scan(
			&j.ID,
			&j.UserID,
			&j.Title,
			&j.Content,
			&tags,
			&j.WordCount,
			&j.CreatedAt,
			&j.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		j.Tags = []string(tags)
		journals = append(journals, &j)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return journals, nil
}
func (r *JournalRepository) Update(j *domain.JournalModel) error {
	j.UpdatedAt = time.Now()
	rows, err := r.db.Exec(`UPDATE journals SET title = $1, content = $2, tags = $3, word_count = $4, updated_at = $5 WHERE id = $6`,
		j.Title, j.Content, j.Tags, j.WordCount, j.UpdatedAt, j.ID,
	)
	if err != nil {
		return err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrJournalNotFound
	}
	return nil
}
func (r *JournalRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM journals WHERE id = $1`, id)
	return err
}
