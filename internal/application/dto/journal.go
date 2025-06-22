package dto

import (
	"time"

	"github.com/mauFade/journaly/internal/domain"
)

type CreateJournalRequest struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	WordCount int      `json:"word_count"`
	Tags      []string `json:"tags"`
}

type UpdateJournalRequest struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	WordCount int      `json:"word_count"`
	Tags      []string `json:"tags"`
}

type JournalResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	WordCount int       `json:"word_count"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToJournalDto(j *domain.JournalModel) *JournalResponse {
	return &JournalResponse{
		ID:        j.ID,
		UserID:    j.UserID,
		Title:     j.Title,
		Content:   j.Content,
		WordCount: j.WordCount,
		Tags:      j.Tags,
		CreatedAt: j.CreatedAt,
		UpdatedAt: j.UpdatedAt,
	}
}
