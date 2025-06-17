package journalservice

import "github.com/mauFade/journaly/internal/domain"

func (s *JournalService) ListUserJournals(userId string) ([]*domain.JournalModel, error) {
	return s.repository.GetByUser(userId)
}
