# Diagramas de Arquitectura — StudyNotes / TaskSync

A continuación se presentan los diagramas de arquitectura esenciales para comprender la estructura conceptual del sistema en sus interacciones, de contenedores y los procesos secuenciales de casos de uso principales.

### Diagrama de Casos de Uso UML

```mermaid
flowchart TD
  ActorUsuario((Usuario))
  
  ActorUsuario --> UC1(Crear Nota)
  ActorUsuario --> UC2(Ver Lista de Notas)
  ActorUsuario --> UC3(Alternar Estado)
```

### Diagrama de Contenedores C4 Nivel 2

```mermaid
flowchart LR
  subgraph "Navegador del Usuario"
    SPA["React SPA\nVite + TypeScript"]
    ADAPTER{"Storage Adapter\n(Selección automática)"}
    LOCAL["LocalStorageApiClient\n(Modo sin conexión)"]
    HTTP["HttpApiClient\n(Modo con backend)"]
  end
  subgraph "Servidor Local"
    API["API Go\nchi + net/http"]
    STORE["Almacén en Memoria\nsync.RWMutex"]
  end
  subgraph "GitHub Pages (Estático)"
    PAGES["Dist Vite\n(solo LocalStorage)"]
  end
  
  SPA --> ADAPTER
  ADAPTER -->|"VITE_USE_MOCK=true o sin conexión"| LOCAL
  ADAPTER -->|"backend accesible"| HTTP
  HTTP -->|"REST JSON"| API
  API --> STORE
  SPA -.->|"deploy"| PAGES
```

### Diagrama de Secuencia UML (Crear Nota)

```mermaid
sequenceDiagram
  participant U as Usuario
  participant F as NoteForm
  participant A as StorageAdapter
  participant H as HTTP Handler
  participant S as NoteService
  participant R as MemoryRepository
  
  U->>F: Ingresa título y envía
  F->>F: Validación cliente (trim, no vacío, ≤100 chars)
  F->>A: createNote(título)
  A->>H: POST /api/v1/notes
  H->>S: CreateNote(título)
  S->>S: Validación servidor + generar UUID
  S->>R: Create(note)
  R-->>S: note guardada
  S-->>H: note
  H-->>A: 201 Created + JSON
  A-->>F: Note
  F->>F: Limpiar input + refrescar lista
```
