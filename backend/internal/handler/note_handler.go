package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"todoapp/backend/internal/domain"
	"todoapp/backend/internal/repository"
	"todoapp/backend/internal/service"
)

// NoteHandler maneja las peticiones HTTP para notas.
type NoteHandler struct {
	svc *service.NoteService
}

// NewNoteHandler crea un nuevo NoteHandler.
func NewNoteHandler(svc *service.NoteService) *NoteHandler {
	return &NoteHandler{svc: svc}
}

// writeJSON es una función de ayuda para escribir respuestas JSON.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// writeError es una función de ayuda para escribir errores en JSON.
func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// HealthCheck responde con un estado ok.
// @spec NOTES-API-007
func (h *NoteHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// GetNotes lista las notas actuales.
// @spec NOTES-API-003
// @spec NOTES-API-004
func (h *NoteHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.svc.ListNotes()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error interno del servidor al listar notas")
		return
	}

	// Asegurar que se devuelva [] en lugar de null cuando no hay notas
	if notes == nil {
		notes = []domain.Note{}
	}

	writeJSON(w, http.StatusOK, notes)
}

// CreateNoteRequest define el cuerpo esperado para crear una nota.
type CreateNoteRequest struct {
	Title string `json:"title"`
}

// CreateNote crea una nueva nota.
// @spec NOTES-API-001
// @spec NOTES-API-002
func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req CreateNoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "cuerpo de petición inválido")
		return
	}

	note, err := h.svc.CreateNote(req.Title)
	if err != nil {
		if err == service.ErrEmptyTitle || err == service.ErrTitleTooLong {
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}
		writeError(w, http.StatusInternalServerError, "error interno al crear la nota")
		return
	}

	writeJSON(w, http.StatusCreated, note)
}

// ToggleStatus cambia el estado de la nota dada por su ID.
// @spec NOTES-API-005
// @spec NOTES-API-006
func (h *NoteHandler) ToggleStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "se requiere el id de la nota")
		return
	}

	note, err := h.svc.ToggleStatus(id)
	if err != nil {
		if err == repository.ErrNotFound {
			writeError(w, http.StatusNotFound, "nota no encontrada")
			return
		}
		writeError(w, http.StatusInternalServerError, "error interno al actualizar la nota")
		return
	}

	writeJSON(w, http.StatusOK, note)
}
