package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"strings"
	"time"
	"todoapp/backend/internal/domain"
)

var (
	ErrEmptyTitle    = errors.New("el título de la nota no puede estar vacío")
	ErrTitleTooLong  = errors.New("el título no puede exceder los 100 caracteres")
)

// NoteService proporciona la lógica de negocio para las notas.
type NoteService struct {
	repo domain.NoteRepository
}

// NewNoteService crea una nueva instancia de NoteService.
func NewNoteService(repo domain.NoteRepository) *NoteService {
	return &NoteService{repo: repo}
}

// generateUUIDv4 genera un identificador UUID v4 usando crypto/rand.
func generateUUIDv4() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	
	// Establecer la versión 4 y la variante
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16]), nil
}

// CreateNote crea una nueva nota tras validarla.
// @spec NOTES-API-001
// @spec NOTES-API-002
func (s *NoteService) CreateNote(title string) (domain.Note, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return domain.Note{}, ErrEmptyTitle
	}

	if len(title) > 100 {
		return domain.Note{}, ErrTitleTooLong
	}

	id, err := generateUUIDv4()
	if err != nil {
		return domain.Note{}, fmt.Errorf("error al generar UUID: %w", err)
	}

	note := domain.Note{
		ID:        id,
		Title:     title,
		Status:    domain.StatusPending,
		CreatedAt: time.Now().UTC(),
	}

	return s.repo.Create(note)
}

// ListNotes devuelve la lista de notas.
// @spec NOTES-API-003
// @spec NOTES-API-004
func (s *NoteService) ListNotes() ([]domain.Note, error) {
	return s.repo.List()
}

// ToggleStatus cambia el estado de una nota.
// @spec NOTES-API-005
// @spec NOTES-API-006
func (s *NoteService) ToggleStatus(id string) (domain.Note, error) {
	return s.repo.ToggleStatus(id)
}
