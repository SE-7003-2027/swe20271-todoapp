package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"

	"todoapp/backend/internal/domain"
	"todoapp/backend/internal/repository"
	"todoapp/backend/internal/service"
)

func setupTestApp() (*chi.Mux, *repository.MemoryRepository) {
	repo := repository.NewMemoryRepository()
	svc := service.NewNoteService(repo)
	h := NewNoteHandler(svc)

	r := chi.NewRouter()
	r.Get("/healthz", h.HealthCheck)
	r.Get("/api/v1/notes", h.GetNotes)
	r.Post("/api/v1/notes", h.CreateNote)
	r.Patch("/api/v1/notes/{id}/toggle", h.ToggleStatus)

	return r, repo
}

func TestHealthCheck(t *testing.T) {
	// @spec NOTES-API-007
	r, _ := setupTestApp()

	req := httptest.NewRequest("GET", "/healthz", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado %d, obtenido %d", http.StatusOK, rr.Code)
	}

	var res map[string]string
	json.NewDecoder(rr.Body).Decode(&res)
	if res["status"] != "ok" {
		t.Errorf("esperado status=ok, obtenido %s", res["status"])
	}
}

func TestCreateNote(t *testing.T) {
	// @spec NOTES-API-001
	// @spec NOTES-API-002
	r, _ := setupTestApp()

	tests := []struct {
		name         string
		body         map[string]string
		expectedCode int
	}{
		{
			name:         "Válido",
			body:         map[string]string{"title": "Hacer la tarea"},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "Título vacío",
			body:         map[string]string{"title": "   "},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Título largo",
			body:         map[string]string{"title": string(make([]byte, 105))},
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b, _ := json.Marshal(tc.body)
			req := httptest.NewRequest("POST", "/api/v1/notes", bytes.NewReader(b))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			r.ServeHTTP(rr, req)

			if rr.Code != tc.expectedCode {
				t.Errorf("esperado %d, obtenido %d", tc.expectedCode, rr.Code)
			}
		})
	}
}

func TestGetNotesEmpty(t *testing.T) {
	// @spec NOTES-API-004
	r, _ := setupTestApp()

	req := httptest.NewRequest("GET", "/api/v1/notes", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado %d, obtenido %d", http.StatusOK, rr.Code)
	}

	body := rr.Body.String()
	if body != "[]\n" {
		t.Errorf("esperado []\\n, obtenido %s", body)
	}
}

func TestGetNotesSort(t *testing.T) {
	// @spec NOTES-API-003
	r, repo := setupTestApp()

	now := time.Now()
	repo.Create(domain.Note{ID: "1", Title: "Nota Vieja", CreatedAt: now.Add(-1 * time.Hour)})
	repo.Create(domain.Note{ID: "2", Title: "Nota Nueva", CreatedAt: now})

	req := httptest.NewRequest("GET", "/api/v1/notes", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado %d, obtenido %d", http.StatusOK, rr.Code)
	}

	var notes []domain.Note
	json.NewDecoder(rr.Body).Decode(&notes)

	if len(notes) != 2 {
		t.Fatalf("esperado 2 notas, obtenido %d", len(notes))
	}

	if notes[0].ID != "2" || notes[1].ID != "1" {
		t.Errorf("orden incorrecto, primero debería ser el más reciente")
	}
}

func TestToggleStatus(t *testing.T) {
	// @spec NOTES-API-005
	// @spec NOTES-API-006
	r, repo := setupTestApp()

	repo.Create(domain.Note{ID: "mi-nota", Title: "Prueba", Status: domain.StatusPending, CreatedAt: time.Now()})

	tests := []struct {
		name         string
		id           string
		expectedCode int
	}{
		{
			name:         "Nota existe",
			id:           "mi-nota",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Nota no existe",
			id:           "no-existe",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Usando httptest puro ya que el router maneja el path y los params
			url := "/api/v1/notes/" + tc.id + "/toggle"
			req := httptest.NewRequest("PATCH", url, nil)
			rr := httptest.NewRecorder()

			// Simulando el chi router ctx si usamos handler directo, pero estamos usando ServeHTTP de router chi
			r.ServeHTTP(rr, req)

			if rr.Code != tc.expectedCode {
				t.Errorf("esperado %d, obtenido %d en %s", tc.expectedCode, rr.Code, tc.name)
			}
		})
	}
}
