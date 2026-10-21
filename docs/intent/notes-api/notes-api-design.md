# LLD — Segmento: notes-api

## Propósito
Definir la estructura, contratos y comportamiento del backend escrito en Go responsable de la gestión de notas, así como su verificación tanto a nivel unitario como de integración y extremo a extremo (E2E) sobre HTTP real.

## Interfaz de API
- `GET /healthz` — Verificación de salud del servidor (`200 OK`, `{"status":"ok"}`).
- `GET /api/v1/notes` — Lista todas las notas existentes ordenadas por `createdAt` descendente (`200 OK`). Devuelve `[]` cuando no hay notas.
- `POST /api/v1/notes` — Crea una nueva nota a partir de `{"title": "..."}` (`201 Created`).
- `PATCH /api/v1/notes/{id}/toggle` — Alterna atómicamente el estado de una nota entre `PENDING` y `COMPLETED` (`200 OK`).

## Modelo de Dominio
El modelo principal se representa mediante el struct `Note` en `internal/domain/note.go`:
- `id` (`string`): Identificador único UUID v4.
- `title` (`string`): Título de la nota recortado (`strings.TrimSpace`), entre 1 y 100 caracteres.
- `status` (`Status`): Estado actual de la nota (`"PENDING"` o `"COMPLETED"`).
- `createdAt` (`time.Time`): Marca de tiempo en UTC de creación de la nota.

## Flujo de Creación y Ciclo de Vida
1. El cliente envía un payload JSON a `POST /api/v1/notes` con el campo `title`.
2. `NoteHandler.CreateNote` decodifica el cuerpo y delega a `NoteService.CreateNote`.
3. `NoteService` aplica `strings.TrimSpace(title)` y valida que no sea vacío ni supere los 100 caracteres.
4. Superada la validación, genera un UUID v4 mediante `crypto/rand` e inicializa la nota con estado `PENDING` y `CreatedAt` en UTC.
5. `MemoryRepository.Create` almacena la nota de forma segura para concurrencia usando `sync.RWMutex`.
6. Para validar el contrato completo sin *mocks*, la suite de integración/E2E levanta el servidor HTTP real (o apunta a `E2E_BASE_URL` en el entorno Staging desplegado) y ejecuta la secuencia completa de extremo a extremo (`GET /healthz` $\rightarrow$ `GET /api/v1/notes` $\rightarrow$ `POST` inválido $\rightarrow$ `POST` válido $\rightarrow$ `PATCH /toggle` $\rightarrow$ verificación de persistencia).

## Manejo de Errores
- **400 Bad Request:** Cuerpo JSON malformado, título vacío tras recortar espacios o título mayor a 100 caracteres.
- **404 Not Found:** El identificador `id` en `PATCH /api/v1/notes/{id}/toggle` no existe en el repositorio (`repository.ErrNotFound`).
- **500 Internal Server Error:** Error inesperado no controlado en el servidor.
Todas las respuestas de error devuelven la estructura JSON uniforme: `{ "error": "mensaje en español" }`.

## Tabla de Decisiones y Alternativas
| Decisión | Opción elegida | Alternativas consideradas | Justificación |
|----------|----------------|---------------------------|---------------|
| Persistencia de datos (`v0.1.0`) | Almacén en memoria con `sync.RWMutex` detrás de `NoteRepository` | SQLite, MySQL directo en `v0.1.0` | Minimiza dependencias externas en el MVP mientras deja lista la costura (`NoteRepository`) para intercambiar por MySQL en `v0.2.0`. |
| Generación de UUID | Paquete estándar `crypto/rand` | `github.com/google/uuid` | Evita dependencias externas innecesarias cumpliendo el formato estándar UUID v4. |
| Enrutamiento HTTP | `github.com/go-chi/chi/v5` | `net/http` puro, `gorilla/mux` | Ligero, 100% compatible con `http.Handler` estándar y limpio para parámetros de ruta como `{id}`. |
| Pruebas de Integración y E2E | Suite Go `httptest.NewServer` + soporte para `E2E_BASE_URL` contra Staging vivo | Scripts sueltos en Bash sin tipado | Permite ejecutar exactamente la misma prueba de contrato tanto en `go test` local (levantando un servidor efímero real) como en la Etapa 7 de Jenkins contra `http://staging-web` desplegado. |

## Preguntas Abiertas
- Para la versión `v0.2.0`, definir el esquema DDL de migración inicial para la implementación `MySQLNoteRepository`.
