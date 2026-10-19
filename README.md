# StudyNotes / TaskSync

> Gestor minimalista de tareas y notas — repositorio educativo para el curso de Ingeniería de Software.

[![GitHub Pages](https://img.shields.io/badge/Demo-GitHub%20Pages-blue)](https://YOUR_USERNAME.github.io/todoapp)
[![Licencia MIT](https://img.shields.io/badge/Licencia-MIT-green)](LICENSE)

---

## 🎯 Propósito

Este repositorio es una demostración completa de grado productivo para el curso universitario **SWE20271**, cubriendo:

- **Ciclo de vida del desarrollo de software** — flujo LID de seis fases (HLD → LLD → EARS → pruebas → código)
- **Ingeniería de requisitos** — especificaciones EARS trazables, historias de usuario INVEST, criterios Gherkin
- **Arquitectura evolutiva** — monolito en capas con interfaces listas para intercambio de base de datos
- **Rebanado vertical** — cada funcionalidad cruza UI, lógica de aplicación y persistencia
- **Ingeniería de liberación** — CI/CD con GitHub Actions, etiquetas anotadas, CHANGELOG

---

## 🚀 Inicio Rápido

### Prerrequisitos

| Herramienta | Versión mínima |
|---|---|
| Go | 1.23+ |
| Node.js | 20+ |
| npm | 10+ |
| Docker | 24+ (opcional) |
| Docker Compose | v2+ (opcional) |

### Opción A — Docker Compose (recomendado, sin configuración)

```bash
git clone https://github.com/YOUR_USERNAME/todoapp.git
cd todoapp
docker compose up --build
```

- Frontend: [http://localhost:3000](http://localhost:3000)
- API: [http://localhost:8080](http://localhost:8080)
- Salud: [http://localhost:8080/healthz](http://localhost:8080/healthz)

### Opción B — Ejecución Local

**Backend:**
```bash
cd backend
go mod download
go run cmd/api/main.go
# → Escuchando en :8080
```

**Frontend:**
```bash
cd frontend
npm install
npm run dev
# → http://localhost:5173
```

### Opción C — Demo sin Backend (GitHub Pages)

El frontend detecta automáticamente si el backend es inaccesible y cambia al modo `LocalStorageAdapter`. Los datos se guardan en el navegador.

Para forzar el modo sin conexión:
```bash
VITE_USE_MOCK=true npm run dev
```

---

## 📁 Estructura del Repositorio

```
├── .github/workflows/        # Automatización CI/CD
│   ├── deploy-pages.yml      # Despliegue automático a GitHub Pages
│   └── release.yml           # Compilación cruzada y GitHub Releases
├── .githooks/
│   └── pre-commit            # Ejecuta pruebas antes de cada confirmación
├── docs/
│   ├── high-level-design.md  # HLD (Diseño de Alto Nivel) — LID Fase 1
│   ├── intent/               # Árbol de intención LID
│   │   ├── notes-api/        # LLD + especificaciones EARS del backend
│   │   ├── storage-adapter/  # LLD + especificaciones EARS del adaptador
│   │   └── notes-ui/         # LLD + especificaciones EARS de componentes React
│   ├── 01-requirements-brief.md   # Especificación de requisitos formal
│   ├── 02-architecture-diagrams.md # Diagramas Mermaid (C4, Casos de Uso, Secuencia)
│   └── 03-release-guide.md        # Guía de liberación y DoD
├── backend/                  # API Go (chi v5, almacén en memoria)
│   ├── cmd/api/main.go
│   ├── internal/
│   │   ├── domain/           # Entidades y contratos de repositorio
│   │   ├── handler/          # Enrutamiento HTTP y códigos de estado
│   │   ├── service/          # Lógica de dominio y validación
│   │   └── repository/       # Abstracción de acceso a datos
│   ├── Dockerfile
│   └── Makefile
├── frontend/                 # SPA React 18 + Vite + TypeScript + Tailwind
│   ├── src/
│   │   ├── components/       # NoteForm, NoteList, NoteItem
│   │   └── services/storage/ # StorageAdapter (Http + LocalStorage)
│   └── Dockerfile
├── docker-compose.yml
├── CHANGELOG.md
├── GEMINI.md                 # Configuración LID para el agente
└── README.md
```

---

## 🧪 Pruebas

```bash
# Backend
cd backend && go test ./...

# Frontend
cd frontend && npm test

# Ambos (gancho de pre-confirmación)
.githooks/pre-commit
```

---

## 🔧 Instalar Gancho Git

```bash
git config core.hooksPath .githooks
chmod +x .githooks/pre-commit
```

---

## 📚 Documentación

| Documento | Descripción |
|---|---|
| [Diseño de Alto Nivel](docs/high-level-design.md) | Arquitectura, tenets y decisiones |
| [Especificación de Requisitos](docs/01-requirements-brief.md) | FR, NFR, historias de usuario, Gherkin |
| [Diagramas de Arquitectura](docs/02-architecture-diagrams.md) | C4, Casos de Uso, Secuencia |
| [Guía de Liberación](docs/03-release-guide.md) | DoD, etiquetas Git, CHANGELOG |
| [LLD notes-api](docs/intent/notes-api/notes-api-design.md) | Diseño detallado del backend |
| [LLD storage-adapter](docs/intent/storage-adapter/storage-adapter-design.md) | Diseño detallado del adaptador |
| [LLD notes-ui](docs/intent/notes-ui/notes-ui-design.md) | Diseño detallado de componentes |

---

## 🔌 Endpoints de la API

| Método | Ruta | Descripción |
|---|---|---|
| `GET` | `/healthz` | Verificación de vida |
| `GET` | `/api/v1/notes` | Listar todas las notas (desc por fecha) |
| `POST` | `/api/v1/notes` | Crear una nota |
| `PATCH` | `/api/v1/notes/{id}/toggle` | Alternar estado PENDING ↔ COMPLETED |

---

## 🗺️ Hoja de Ruta

| Versión | Funcionalidad |
|---|---|
| **v0.1.0** ✅ | Crear, listar y alternar estado de notas |
| **v0.2.0** | Eliminación, edición en línea, migración a MySQL |
| **v0.3.0** | Etiquetas, búsqueda, autenticación |

---

## 🏛️ Flujo de Trabajo LID

Todas las contribuciones deben seguir el flujo de seis fases de Linked-Intent Development:

```
1. HLD  →  2. LLD  →  3. EARS  →  4. Auditoría  →  5. Pruebas  →  6. Código
```

Consulta [GEMINI.md](GEMINI.md) para la configuración completa del agente.

---

## 📄 Licencia

MIT — libre para uso educativo.
