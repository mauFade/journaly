package journalservice

import (
	"time"

	"github.com/google/uuid"
	"github.com/mauFade/journaly/internal/application/dto"
	"github.com/mauFade/journaly/internal/domain"
)

func (s *JournalService) CreateJournal(req *dto.CreateJournalRequest, userId string) (*dto.JournalResponse, error) {
	j := domain.NewJournalModel(uuid.NewString(), userId, req.Title, req.Content, req.WordCount, req.Tags, time.Now(), time.Now())
	err := s.repository.Save(j)
	if err != nil {
		return nil, err
	}

	return dto.ToJournalDto(j), nil
}
