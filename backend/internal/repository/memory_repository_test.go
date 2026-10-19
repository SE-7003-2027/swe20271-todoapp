package repository

import (
	"testing"
	"time"
	"todoapp/backend/internal/domain"
)

func TestMemoryRepository_Create(t *testing.T) {
	// @spec NOTES-API-001
	repo := NewMemoryRepository()
	note := domain.Note{
		ID:        "123",
		Title:     "Comprar leche",
		Status:    domain.StatusPending,
		CreatedAt: time.Now(),
	}

	created, err := repo.Create(note)
	if err != nil {
		t.Fatalf("Error inesperado al crear nota: %v", err)
	}

	if created.ID != note.ID || created.Title != note.Title {
		t.Errorf("La nota creada no coincide con la original. Obtenido: %+v", created)
	}
}

func TestMemoryRepository_ListEmpty(t *testing.T) {
	// @spec NOTES-API-004
	repo := NewMemoryRepository()
	list, err := repo.List()
	if err != nil {
		t.Fatalf("Error inesperado al listar notas: %v", err)
	}

	if len(list) != 0 {
		t.Errorf("Se esperaba una lista vacía, se obtuvo longitud %d", len(list))
	}
}

func TestMemoryRepository_ListSortOrder(t *testing.T) {
	// @spec NOTES-API-003
	repo := NewMemoryRepository()

	now := time.Now()
	repo.Create(domain.Note{ID: "1", Title: "Primera", CreatedAt: now.Add(-2 * time.Hour)})
	repo.Create(domain.Note{ID: "2", Title: "Tercera", CreatedAt: now})
	repo.Create(domain.Note{ID: "3", Title: "Segunda", CreatedAt: now.Add(-1 * time.Hour)})

	list, err := repo.List()
	if err != nil {
		t.Fatalf("Error inesperado: %v", err)
	}

	if len(list) != 3 {
		t.Fatalf("Se esperaban 3 notas, se obtuvieron %d", len(list))
	}

	if list[0].ID != "2" || list[1].ID != "3" || list[2].ID != "1" {
		t.Errorf("El orden de las notas es incorrecto. Obtenido: %v", list)
	}
}

func TestMemoryRepository_ToggleStatus(t *testing.T) {
	// @spec NOTES-API-005
	repo := NewMemoryRepository()
	note := domain.Note{
		ID:        "123",
		Title:     "Comprar leche",
		Status:    domain.StatusPending,
		CreatedAt: time.Now(),
	}
	repo.Create(note)

	// Cambiar a COMPLETED
	toggled, err := repo.ToggleStatus("123")
	if err != nil {
		t.Fatalf("Error inesperado al cambiar estado: %v", err)
	}
	if toggled.Status != domain.StatusCompleted {
		t.Errorf("Se esperaba estado COMPLETED, se obtuvo %s", toggled.Status)
	}

	// Cambiar de vuelta a PENDING
	toggled, err = repo.ToggleStatus("123")
	if err != nil {
		t.Fatalf("Error inesperado al cambiar estado: %v", err)
	}
	if toggled.Status != domain.StatusPending {
		t.Errorf("Se esperaba estado PENDING, se obtuvo %s", toggled.Status)
	}
}

func TestMemoryRepository_ToggleStatusNotFound(t *testing.T) {
	// @spec NOTES-API-006
	repo := NewMemoryRepository()
	_, err := repo.ToggleStatus("no-existe")
	if err != ErrNotFound {
		t.Errorf("Se esperaba ErrNotFound, se obtuvo %v", err)
	}
}
