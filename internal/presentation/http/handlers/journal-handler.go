package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mauFade/journaly/internal/application/dto"
	journalservice "github.com/mauFade/journaly/internal/application/service/journal-service"
	"github.com/mauFade/journaly/internal/domain"
)

type JournalHandler struct {
	journalService *journalservice.JournalService
}

func NewJournalhandler(s *journalservice.JournalService) *JournalHandler {
	return &JournalHandler{
		journalService: s,
	}
}

func (h *JournalHandler) CreateJournal(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateJournalRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	userId := r.Context().Value(domain.UserKey).(string)

	res, err := h.journalService.CreateJournal(&input, userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *JournalHandler) ListUserJournals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := r.Context().Value(domain.UserKey).(string)
	res, err := h.journalService.ListUserJournals(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (h *JournalHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input dto.UpdateJournalRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	res, err := h.journalService.UpdateJournal(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	json.NewEncoder(w).Encode(res)
}
