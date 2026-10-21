package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"todoapp/backend/internal/domain"
)

// handlerRoundTripper permite que un *http.Client estándar ejecute peticiones HTTP
// completas directamente contra el enrutador chi en memoria cuando no se define E2E_BASE_URL.
type handlerRoundTripper struct {
	handler http.Handler
}

func (rt *handlerRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	rr := httptest.NewRecorder()
	rt.handler.ServeHTTP(rr, req)
	return rr.Result(), nil
}

// TestNoteLifecycleIntegrationAndE2E ejecuta el flujo completo de extremo a extremo (E2E)
// atravesando todas las capas reales (Router -> Handler -> Service -> Repository).
// Si la variable de entorno E2E_BASE_URL está definida (ej. en la Etapa 7 de Jenkins CD),
// ejecuta las peticiones TCP reales contra el entorno Staging desplegado en vivo; de lo contrario,
// ejecuta el cliente HTTP contra el stack completo en memoria.
// @spec NOTES-API-008
func TestNoteLifecycleIntegrationAndE2E(t *testing.T) {
	baseURL := strings.TrimRight(os.Getenv("E2E_BASE_URL"), "/")
	isLiveTarget := baseURL != ""

	client := http.DefaultClient
	if !isLiveTarget {
		r, _ := setupTestApp()
		client = &http.Client{Transport: &handlerRoundTripper{handler: r}}
		baseURL = "http://local.integration"
	} else {
		// En modo E2E contra el entorno Staging desplegado (Nginx + API Go),
		// verificar también que la raíz "/" sirva el HTML de la SPA de React.
		respRoot, err := client.Get(baseURL + "/")
		if err != nil {
			t.Fatalf("error al conectar con el frontend desplegado en %s/: %v", baseURL, err)
		}
		defer respRoot.Body.Close()
		if respRoot.StatusCode != http.StatusOK {
			t.Fatalf("esperado 200 OK en el frontend desplegado, obtenido %d", respRoot.StatusCode)
		}
		htmlBytes, _ := io.ReadAll(respRoot.Body)
		if !strings.Contains(string(htmlBytes), "StudyNotes") {
			t.Fatalf("el HTML servido en %s/ no contiene el título esperado 'StudyNotes'", baseURL)
		}
	}

	// 1. Verificar salud del servicio (GET /healthz)
	respHealth, err := client.Get(baseURL + "/healthz")
	if err != nil {
		t.Fatalf("fallo al consultar /healthz: %v", err)
	}
	defer respHealth.Body.Close()
	if respHealth.StatusCode != http.StatusOK {
		t.Fatalf("esperado 200 OK en /healthz, obtenido %d", respHealth.StatusCode)
	}

	// 2. Validar rechazo de título inválido (POST /api/v1/notes -> 400 Bad Request)
	invalidPayload, _ := json.Marshal(map[string]string{"title": "    "})
	respInvalid, err := client.Post(baseURL+"/api/v1/notes", "application/json", bytes.NewReader(invalidPayload))
	if err != nil {
		t.Fatalf("fallo en POST inválido: %v", err)
	}
	defer respInvalid.Body.Close()
	if respInvalid.StatusCode != http.StatusBadRequest {
		t.Fatalf("esperado 400 Bad Request para título vacío, obtenido %d", respInvalid.StatusCode)
	}

	// 3. Crear una nota válida (FR-01: POST /api/v1/notes -> 201 Created)
	noteTitle := "Nota de integración E2E"
	validPayload, _ := json.Marshal(map[string]string{"title": "  " + noteTitle + "  "})
	respCreate, err := client.Post(baseURL+"/api/v1/notes", "application/json", bytes.NewReader(validPayload))
	if err != nil {
		t.Fatalf("fallo al crear nota en E2E: %v", err)
	}
	defer respCreate.Body.Close()
	if respCreate.StatusCode != http.StatusCreated {
		t.Fatalf("esperado 201 Created al crear nota, obtenido %d", respCreate.StatusCode)
	}

	var created domain.Note
	if err := json.NewDecoder(respCreate.Body).Decode(&created); err != nil {
		t.Fatalf("error al decodificar nota creada: %v", err)
	}
	if created.ID == "" {
		t.Fatal("se esperaba un UUID no vacío en la nota creada")
	}
	if created.Title != noteTitle {
		t.Fatalf("esperado título recortado %q, obtenido %q", noteTitle, created.Title)
	}
	if created.Status != domain.StatusPending {
		t.Fatalf("esperado estado inicial PENDING, obtenido %s", created.Status)
	}

	// 4. Consultar el listado y verificar que la nota recién creada esté presente (FR-02: GET /api/v1/notes)
	respList1, err := client.Get(baseURL + "/api/v1/notes")
	if err != nil {
		t.Fatalf("fallo al listar notas: %v", err)
	}
	defer respList1.Body.Close()
	var notesAfterCreate []domain.Note
	if err := json.NewDecoder(respList1.Body).Decode(&notesAfterCreate); err != nil {
		t.Fatalf("error al decodificar lista de notas: %v", err)
	}

	foundPending := false
	for _, n := range notesAfterCreate {
		if n.ID == created.ID && n.Status == domain.StatusPending {
			foundPending = true
			break
		}
	}
	if !foundPending {
		t.Fatalf("la nota creada (%s) no se encontró con estado PENDING en GET /api/v1/notes", created.ID)
	}

	// 5. Alternar el estado de PENDING a COMPLETED (FR-03: PATCH /api/v1/notes/{id}/toggle)
	reqToggle1, _ := http.NewRequest(http.MethodPatch, baseURL+"/api/v1/notes/"+created.ID+"/toggle", nil)
	respToggle1, err := client.Do(reqToggle1)
	if err != nil {
		t.Fatalf("fallo al alternar estado de la nota: %v", err)
	}
	defer respToggle1.Body.Close()
	if respToggle1.StatusCode != http.StatusOK {
		t.Fatalf("esperado 200 OK al alternar estado, obtenido %d", respToggle1.StatusCode)
	}

	var toggled domain.Note
	if err := json.NewDecoder(respToggle1.Body).Decode(&toggled); err != nil {
		t.Fatalf("error al decodificar nota alternada: %v", err)
	}
	if toggled.Status != domain.StatusCompleted {
		t.Fatalf("esperado estado COMPLETED tras primer toggle, obtenido %s", toggled.Status)
	}

	// 6. Confirmar persistencia del estado COMPLETED en un nuevo GET /api/v1/notes
	respList2, err := client.Get(baseURL + "/api/v1/notes")
	if err != nil {
		t.Fatalf("fallo al re-consultar lista de notas: %v", err)
	}
	defer respList2.Body.Close()
	var notesAfterToggle []domain.Note
	if err := json.NewDecoder(respList2.Body).Decode(&notesAfterToggle); err != nil {
		t.Fatalf("error al decodificar segunda lista de notas: %v", err)
	}

	foundCompleted := false
	for _, n := range notesAfterToggle {
		if n.ID == created.ID && n.Status == domain.StatusCompleted {
			foundCompleted = true
			break
		}
	}
	if !foundCompleted {
		t.Fatalf("la nota (%s) no persistió con estado COMPLETED en GET /api/v1/notes", created.ID)
	}

	// 7. Verificar error 404 al intentar alternar un ID inexistente
	reqNotFound, _ := http.NewRequest(http.MethodPatch, baseURL+"/api/v1/notes/id-inexistente-999/toggle", nil)
	respNotFound, err := client.Do(reqNotFound)
	if err != nil {
		t.Fatalf("fallo en petición PATCH inexistente: %v", err)
	}
	defer respNotFound.Body.Close()
	if respNotFound.StatusCode != http.StatusNotFound {
		t.Fatalf("esperado 404 Not Found para ID inexistente, obtenido %d", respNotFound.StatusCode)
	}
}
