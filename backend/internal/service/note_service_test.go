package service

import (
	"errors"
	"strings"
	"testing"
	"todoapp/backend/internal/domain"
)

// mockRepository es un repositorio simulado para las pruebas.
type mockRepository struct {
	createFunc       func(note domain.Note) (domain.Note, error)
	listFunc         func() ([]domain.Note, error)
	toggleStatusFunc func(id string) (domain.Note, error)
}

func (m *mockRepository) Create(note domain.Note) (domain.Note, error) {
	if m.createFunc != nil {
		return m.createFunc(note)
	}
	return note, nil
}

func (m *mockRepository) List() ([]domain.Note, error) {
	if m.listFunc != nil {
		return m.listFunc()
	}
	return []domain.Note{}, nil
}

func (m *mockRepository) ToggleStatus(id string) (domain.Note, error) {
	if m.toggleStatusFunc != nil {
		return m.toggleStatusFunc(id)
	}
	return domain.Note{}, nil
}

func TestNoteService_CreateNote(t *testing.T) {
	repo := &mockRepository{}
	svc := NewNoteService(repo)

	// @spec NOTES-API-002
	t.Run("Título vacío es rechazado", func(t *testing.T) {
		_, err := svc.CreateNote("")
		if err != ErrEmptyTitle {
			t.Errorf("Se esperaba ErrEmptyTitle, se obtuvo %v", err)
		}
	})

	// @spec NOTES-API-002
	t.Run("Título de solo espacios es rechazado", func(t *testing.T) {
		_, err := svc.CreateNote("   \t  ")
		if err != ErrEmptyTitle {
			t.Errorf("Se esperaba ErrEmptyTitle para solo espacios, se obtuvo %v", err)
		}
	})

	// @spec NOTES-API-002
	t.Run("Título demasiado largo es rechazado", func(t *testing.T) {
		longTitle := strings.Repeat("A", 101)
		_, err := svc.CreateNote(longTitle)
		if err != ErrTitleTooLong {
			t.Errorf("Se esperaba ErrTitleTooLong, se obtuvo %v", err)
		}
	})

	// @spec NOTES-API-001
	t.Run("Título válido crea nota con estado PENDING", func(t *testing.T) {
		title := "  Aprender Go  "
		
		var savedNote domain.Note
		repo.createFunc = func(note domain.Note) (domain.Note, error) {
			savedNote = note
			return note, nil
		}

		note, err := svc.CreateNote(title)
		if err != nil {
			t.Fatalf("Error inesperado: %v", err)
		}

		if savedNote.Title != "Aprender Go" {
			t.Errorf("Se esperaba que el título fuera recortado, se obtuvo '%s'", savedNote.Title)
		}

		if savedNote.Status != domain.StatusPending {
			t.Errorf("El estado inicial debería ser PENDING, es %s", savedNote.Status)
		}

		if savedNote.ID == "" || len(savedNote.ID) != 36 {
			t.Errorf("Se esperaba un UUID v4, se obtuvo %s", savedNote.ID)
		}

		if note.ID != savedNote.ID {
			t.Errorf("La nota devuelta difiere de la guardada")
		}
	})
}

func TestNoteService_ToggleStatus(t *testing.T) {
	// @spec NOTES-API-005
	var toggledID string
	repo := &mockRepository{
		toggleStatusFunc: func(id string) (domain.Note, error) {
			toggledID = id
			return domain.Note{ID: id, Status: domain.StatusCompleted}, nil
		},
	}
	svc := NewNoteService(repo)

	note, err := svc.ToggleStatus("mi-id")
	if err != nil {
		t.Fatalf("Error inesperado: %v", err)
	}

	if toggledID != "mi-id" {
		t.Errorf("Se esperaba delegar el id 'mi-id', se delegó '%s'", toggledID)
	}

	if note.Status != domain.StatusCompleted {
		t.Errorf("Se esperaba COMPLETED, se obtuvo %s", note.Status)
	}
}

func TestNoteService_ToggleStatusNotFound(t *testing.T) {
	// @spec NOTES-API-006
	repoError := errors.New("no encontrado")
	repo := &mockRepository{
		toggleStatusFunc: func(id string) (domain.Note, error) {
			return domain.Note{}, repoError
		},
	}
	svc := NewNoteService(repo)

	_, err := svc.ToggleStatus("no-existe")
	if err != repoError {
		t.Errorf("Se esperaba que el error del repositorio se propagara, se obtuvo %v", err)
	}
}
