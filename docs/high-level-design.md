# Diseño de Alto Nivel — StudyNotes / TaskSync

## Contexto y Propósito
El problema que resuelve este proyecto es la falta de una herramienta minimalista para gestionar notas y tareas enfocada en entornos educativos, sin la sobrecarga de características complejas que abruman al estudiante. Su objetivo educativo primario es demostrar los principios de diseño de software usando el paradigma de Desarrollo Basado en Intenciones Enlazadas (LID), rebanado vertical (*Vertical Slicing*) y entrega continua (CI/CD).

## Arquitectura del Sistema
La arquitectura es un monolito por capas estructurado de la siguiente forma:
- **Handler**: Maneja el enrutamiento HTTP (`chi` v5), decodificación JSON y códigos de estado.
- **Service**: Contiene la lógica de negocio, generación de UUID v4 y validaciones de dominio.
- **Repository**: Administra el acceso a datos mediante la interfaz `NoteRepository` (implementada en memoria con `sync.RWMutex` en `v0.1.0`).

En el frontend (React 18 + Vite + TypeScript + Tailwind CSS), se utiliza el patrón **Adapter** (`StorageAdapter`) para aislar la interfaz de usuario de la fuente de datos (`HttpApiClient` o `LocalStorageApiClient`). La topología del proyecto incluye soporte para Docker Compose tanto para ejecución local como para un pipeline de CI/CD en **Jenkins** que despliega en un entorno **Staging Full-Stack** (`:8082` y `:8083`).

## Tenets
1. **Mínima complejidad sobre abstracción prematura:** Se prioriza código simple, legible y directo en lugar de abstracciones teóricas que no proporcionan valor inmediato para una versión temprana.
2. **La demo debe funcionar sin conexión:** El `LocalStorageApiClient` es considerado un modo de primera clase y no solo una contingencia. El producto está diseñado para operar desde el navegador en entornos estáticos como GitHub Pages.
3. **La documentación y las pruebas son código:** Todo lo referente a especificaciones EARS, diseños de bajo nivel (LLD) y la pirámide completa de pruebas (unitarias, integración y E2E) se entregan junto con el código fuente y bloquean el despliegue si fallan.

## Índice de Segmentos LLD
| Segmento LLD | Descripción |
|--------------|-------------|
| [notes-api](./intent/notes-api/notes-api-design.md) | Diseño del backend en Go (Endpoints REST, Servicio, Repositorio y pruebas de integración/E2E) |
| [storage-adapter](./intent/storage-adapter/storage-adapter-design.md) | Adaptadores de persistencia (`HttpApiClient`, `LocalStorageApiClient`) y autodetección de salud |
| [notes-ui](./intent/notes-ui/notes-ui-design.md) | Componentes React (`App`, `NoteForm`, `NoteList`, `NoteItem`) e integración de interfaz |

## Decisiones Arquitectónicas
- **ADR-001: Monolito en capas (vs. microservicios):** Se eligió un diseño de monolito por su valor pedagógico. Su simplicidad permite entender la separación de responsabilidades (`Handler -> Service -> Repository`) sin la sobrecarga operativa de microservicios.
- **ADR-002: Portabilidad estática:** Se utiliza el patrón Adapter en el frontend para permitir el despliegue estático en GitHub Pages (`LocalStorageApiClient`) y la conexión al backend real en Go (`HttpApiClient`) en entornos locales y de Staging.
- **ADR-003: Pirámide de pruebas ligera (Unitarias + Integración + E2E Post-Despliegue):** Además de las pruebas unitarias aisladas por capa, el sistema verifica la integración real entre componentes (`httptest.Server` en Go y montaje completo de `<App />` en React) y ejecuta una suite E2E contra el entorno Staging desplegado por Jenkins antes de dar por válido el pipeline, sin depender de navegadores headless pesados.

## Fuera del Alcance v0.1.0
Las siguientes funciones se han postergado de esta iteración:
- Eliminación de notas (`FR-04`).
- Edición en línea de notas previamente creadas (`FR-05`).
- Etiquetas y sistema de categorización (`FR-06`).
- Autenticación y cuentas de usuario.
