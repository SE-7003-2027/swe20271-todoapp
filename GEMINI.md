# Configuración del Agente — StudyNotes / TaskSync

Este archivo configura el comportamiento del agente de IA para este repositorio.

## LID

- Mode: Full
- Version: linked-intent-dev@latest

El modo Full significa que cualquier solicitud que pueda resultar en un cambio de código
activa el flujo de seis fases de LID:

```
HLD → LLD → Especificaciones EARS → Auditoría de bordes → Pruebas primero → Código
```

Cada límite de fase es una parada obligatoria para revisión del usuario.

## LID Tooling

- Árbol de intención: `docs/intent/`
- HLD: `docs/high-level-design.md`
- Formato de IDs de especificación: `<SEGMENTO>-<NNN>` (ej. `NOTES-API-001`)

## Idioma

Todos los documentos, comentarios de código, texto de interfaz y mensajes de error
deben estar en **español de México (es-MX)**. Los identificadores de código
(nombres de variables, funciones, tipos) permanecen en inglés.

## Contexto del Proyecto

- **Sistema:** StudyNotes / TaskSync — gestor minimalista de tareas y notas
- **Versión actual:** v0.1.0 (MVP)
- **Frontend:** React 18 + Vite + TypeScript + Tailwind CSS v3
- **Backend:** Go + chi v5 + almacén en memoria (sync.RWMutex)
- **Persistencia futura (v0.2.0):** MySQL
- **Despliegue estático:** GitHub Pages (LocalStorageAdapter como modo sin conexión)
- **Contenedores:** Docker + Docker Compose

## Segmentos LID

| Segmento | LLD | Alcance |
|---|---|---|
| `notes-api` | `docs/intent/notes-api/notes-api-design.md` | Handler, Service, Repository en Go |
| `storage-adapter` | `docs/intent/storage-adapter/storage-adapter-design.md` | HttpAdapter, LocalAdapter, lógica de selección |
| `notes-ui` | `docs/intent/notes-ui/notes-ui-design.md` | Componentes React: NoteForm, NoteList, NoteItem |
