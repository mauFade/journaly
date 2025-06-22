package journalservice

import (
	"github.com/mauFade/journaly/internal/application/dto"
	"github.com/mauFade/journaly/internal/domain"
)

func (s *JournalService) UpdateJournal(data *dto.UpdateJournalRequest) (*dto.JournalResponse, error) {
	j, err := s.repository.GetByID(data.ID)
	if err != nil {
		return nil, err
	}
	if j == nil {
		return nil, domain.ErrJournalNotFound
	}

	if data.Title != "" {
		j.Title = data.Title
	}
	if data.Content != "" {
		j.Content = data.Content
	}
	if data.WordCount > 0 {
		j.WordCount = data.WordCount
	}
	if len(data.Tags) > 0 {
		j.Tags = data.Tags
	}
	if err = s.repository.Update(j); err != nil {
		return nil, err
	}

	return dto.ToJournalDto(j), nil
}
