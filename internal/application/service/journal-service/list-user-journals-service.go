package journalservice

import (
	"github.com/mauFade/journaly/internal/application/dto"
)

func (s *JournalService) ListUserJournals(userId string) ([]*dto.JournalResponse, error) {
	js, err := s.repository.GetByUser(userId)

	if err != nil {
		return nil, err
	}

	journals := []*dto.JournalResponse{}
	for _, j := range js {
		journals = append(journals, dto.ToJournalDto(j))
	}
	return journals, nil
}
