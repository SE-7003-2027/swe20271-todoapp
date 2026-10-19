package repository

import (
	"errors"
	"sort"
	"sync"
	"todoapp/backend/internal/domain"
)

// ErrNotFound es el error devuelto cuando no se encuentra una nota.
var ErrNotFound = errors.New("nota no encontrada")

// MemoryRepository es una implementación en memoria de NoteRepository.
type MemoryRepository struct {
	mu    sync.RWMutex
	notes map[string]domain.Note
}

// NewMemoryRepository crea una nueva instancia de MemoryRepository.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		notes: make(map[string]domain.Note),
	}
}

// Create almacena una nueva nota.
// @spec NOTES-API-001
func (r *MemoryRepository) Create(note domain.Note) (domain.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.notes[note.ID] = note
	return note, nil
}

// List devuelve todas las notas ordenadas por fecha de creación descendente.
// @spec NOTES-API-003
// @spec NOTES-API-004
func (r *MemoryRepository) List() ([]domain.Note, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var list []domain.Note
	for _, note := range r.notes {
		list = append(list, note)
	}

	// Ordenar por createdAt de manera descendente
	sort.Slice(list, func(i, j int) bool {
		return list[i].CreatedAt.After(list[j].CreatedAt)
	})

	return list, nil
}

// ToggleStatus cambia atómicamente el estado de una nota entre PENDING y COMPLETED.
// Retorna ErrNotFound si la nota no existe.
// @spec NOTES-API-005
// @spec NOTES-API-006
func (r *MemoryRepository) ToggleStatus(id string) (domain.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	note, exists := r.notes[id]
	if !exists {
		return domain.Note{}, ErrNotFound
	}

	if note.Status == domain.StatusPending {
		note.Status = domain.StatusCompleted
	} else {
		note.Status = domain.StatusPending
	}

	r.notes[id] = note
	return note, nil
}
