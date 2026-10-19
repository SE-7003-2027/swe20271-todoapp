package domain

import "time"

// Status representa el estado de una nota
type Status string

const (
	StatusPending   Status = "PENDING"
	StatusCompleted Status = "COMPLETED"
)

// Note es la entidad central del dominio
type Note struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

// NoteRepository define el contrato de acceso a datos (seam para intercambio de BD)
type NoteRepository interface {
	Create(note Note) (Note, error)
	List() ([]Note, error)
	ToggleStatus(id string) (Note, error)
}
