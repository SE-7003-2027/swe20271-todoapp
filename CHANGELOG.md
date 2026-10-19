# Registro de Cambios

Todos los cambios notables de este proyecto serán documentados en este archivo.

El formato está basado en [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
y este proyecto sigue el [Versionado Semántico](https://semver.org/spec/v2.0.0.html).

## [Sin liberar]

### Planificado para v0.2.0
- [FR-04] Endpoint y acción de interfaz para eliminar tareas.
- [FR-05] Edición de título en línea.
- [FR-06] Etiquetado por categorías y filtrado del lado del cliente.
- [INFRA] Migración del almacén en memoria a base de datos MySQL.

---

## [0.1.0] - 2026-09-21

### Agregado
- **[FR-01] Creación de Notas:** Endpoint `POST /api/v1/notes` y formulario frontend con validación cliente/servidor (cadena no vacía, máximo 100 caracteres).
- **[FR-02] Listado de Notas:** Endpoint `GET /api/v1/notes` y lista de tareas ordenada cronológicamente por fecha de creación (descendente).
- **[FR-03] Alternancia de Estado:** Endpoint `PATCH /api/v1/notes/{id}/toggle` con transición inmediata entre `PENDING` y `COMPLETED`.
- **Cliente de Datos Dual:** Interfaz de almacenamiento unificada con respaldo automático:
  - `HttpApiClient`: Se comunica con el backend Go mediante REST.
  - `LocalStorageApiClient`: Respaldo en navegador para despliegue estático sin backend en GitHub Pages.
- **Documentación de Arquitectura:** Diagramas C4 Nivel 2, Casos de Uso UML y Diagrama de Secuencia en formato Mermaid bajo `docs/02-architecture-diagrams.md`.
- **Especificación de Requisitos:** Requisitos formales, historias de usuario INVEST con criterios de aceptación Gherkin, y catálogo de casos límite bajo `docs/01-requirements-brief.md`.
- **LID (Desarrollo de Intención Vinculada):** Árbol de intención completo con HLD, LLDs y especificaciones EARS bajo `docs/intent/`.
- **Automatización CI/CD:**
  - Flujo de trabajo de GitHub Pages para despliegue automático del frontend.
  - Flujo de trabajo de liberación para compilación cruzada de binarios Go (`linux/amd64`, `darwin/arm64`, `windows/amd64`) y generación de Releases en GitHub.
- **Contenedores Docker:** `Dockerfile` para backend y frontend, `docker-compose.yml` para entorno de desarrollo sin configuración.
- **Gancho de Pre-confirmación:** `.githooks/pre-commit` para ejecutar pruebas unitarias antes de cada confirmación Git.

### Decisiones Arquitectónicas
- **ADR-001 (Monolito en capas):** Se implementó un monolito limpio por capas (Handler → Service → Repository) en Go con almacén en memoria thread-safe (`sync.RWMutex`) para minimizar la complejidad de despliegue en demos educativas.
- **ADR-002 (Portabilidad estática):** Se desacopló el acceso a datos de la UI mediante el patrón Strategy/Adapter para permitir ejecución autónoma en GitHub Pages sin cómputo del lado del servidor.
