package domain

import "time"

type JournalModel struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	WordCount int
	Tags      []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewJournalModel(id, userId, title, content string, wordCount int, tags []string, ca, ua time.Time) *JournalModel {
	return &JournalModel{
		ID:        id,
		UserID:    userId,
		Title:     title,
		Content:   content,
		WordCount: wordCount,
		Tags:      tags,
		CreatedAt: ca,
		UpdatedAt: ua,
	}
}
