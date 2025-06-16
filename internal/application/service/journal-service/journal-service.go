package journalservice

import "github.com/mauFade/journaly/internal/domain"

type JournalService struct {
	repository domain.JournalRepository
}

func NewJournalService(r domain.JournalRepository) *JournalService {
	return &JournalService{
		repository: r,
	}
}
