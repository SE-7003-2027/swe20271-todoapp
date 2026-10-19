# LLD — Segmento: notes-api

## Propósito
Definir la estructura y el comportamiento del backend escrito en Go responsable de la gestión de notas, proporcionando una API robusta y simple para los clientes que interactúen con el servicio.

## Interfaz de API
- `POST /api/v1/notes` - Crea una nueva nota a partir de un cuerpo JSON.
- `GET /api/v1/notes` - Lista todas las notas existentes en el sistema.
- `PATCH /api/v1/notes/:id/status` - Alterna el estado de completado de una nota específica por su ID.

## Modelo de Dominio
El modelo principal se representa mediante el struct `Note` con los siguientes campos clave:
- `id` (string): Identificador único universal.
- `title` (string): Título de la nota o tarea introducida por el usuario.
- `completed` (boolean): Estado actual de la tarea indicando si está resuelta o pendiente.
- `createdAt` (datetime): Marca de tiempo indicando cuándo se introdujo la nota.

## Flujo de Creación de Nota
1. El cliente envía un payload JSON al endpoint que contiene únicamente el título.
2. El `Handler` intercepta la petición, decodifica el cuerpo y llama a la capa de `Service`.
3. El `Service` realiza la validación asegurándose de que el título no esté vacío y no supere los 100 caracteres.
4. Una vez superada la validación, se genera un UUID aleatorio para identificar la nueva nota en el sistema de manera inequívoca.
5. Se invoca al `Repository` para persistir la nota en la memoria compartida.
6. El `Service` retorna la nota creada al `Handler`, el cual formatea una respuesta JSON exitosa bajo el código `201 Created`.

## Manejo de Errores
- **400 Bad Request:** Estructura de JSON inválida o fallas de validación de reglas de negocio.
- **404 Not Found:** La nota solicitada para ser actualizada no existe.
- **500 Internal Server Error:** Error inesperado y no controlado en el servidor.
Todas las respuestas de error proveen un objeto uniforme JSON, ejemplo: `{ "error": "Mensaje detallado" }`.

## Tabla de Decisiones y Alternativas
| Decisión | Opción elegida | Alternativas consideradas | Justificación |
|----------|----------------|---------------------------|---------------|
| Persistencia de datos | In-memory store | SQLite, MySQL | Ideal para el entregable v0.1.0. Reduce la complejidad de configuración y de dependencia de servicios externos, manteniendo un enfoque minimalista. |
| Generación de UUID | paquete estándar `crypto/rand` | Librería externa (ej. google/uuid) | Previene la contaminación del proyecto con dependencias de terceros; `crypto/rand` es suficiente y altamente seguro. |
| Enrutamiento HTTP | chi router | stdlib (`net/http`), gorilla/mux | `chi` es eficiente, semánticamente afín a los manejadores por defecto de la librería estándar y facilita middlewares. |

## Preguntas Abiertas
- Para la versión v0.2.0, ¿cuál será el plan de migración desde memoria hacia el motor MySQL propuesto sin interrupción agresiva de compatibilidad?
- ¿Existen requerimientos formales de métricas y observabilidad que deba incluir el API como parte del próximo desarrollo?
