# LLD — Segmento: notes-ui

## Propósito
Definir la jerarquía de componentes en React, las reglas de validación en el cliente y la integración entre la interfaz gráfica y el `StorageAdapter`.

## Árbol de Componentes
- `App` (`frontend/src/App.tsx`): Componente raíz que inicializa el adaptador con `createAdapter()`, gestiona el estado `notes`, muestra el indicador de modo sin conexión cuando el adaptador es `LocalStorageApiClient` y coordina `handleCreate` y `handleToggle`.
  - `NoteForm` (`frontend/src/components/NoteForm.tsx`): Formulario controlado con validación de cliente (`trim`, no vacío y máximo 100 caracteres) y manejo de errores en línea.
  - `NoteList` (`frontend/src/components/NoteList.tsx`): Renderiza el mensaje de estado vacío o la lista de componentes `NoteItem`.
    - `NoteItem` (`frontend/src/components/NoteItem.tsx`): Muestra el título, fecha formateada en `es-MX`, insignia de estado (`PENDIENTE` / `COMPLETADO`) y botón de acción (`Completar` / `Reactivar`).

## Flujo de Estado e Integración
1. **Inicialización (`App`):** Al montarse, invoca `createAdapter()` y muestra `"Cargando aplicación..."` hasta que se resuelve la selección del adaptador y la primera carga de `getNotes()`.
2. **Creación (`NoteForm` $\rightarrow$ `App` $\rightarrow$ `NoteList`):** Al enviar un título válido, `NoteForm` invoca `onSubmit(trimmedTitle)`. Si resuelve sin error, limpia el campo de texto y `App` recarga el listado desde el adaptador para reflejar la nueva nota en `NoteList`. Si el adaptador lanza un error, `NoteForm` muestra el banner de error y conserva el texto escrito por el usuario.
3. **Alternancia (`NoteItem` $\rightarrow$ `App` $\rightarrow$ `NoteList`):** Al hacer clic en `Completar` o `Reactivar`, `App` invoca `adapter.toggleNote(id)` y actualiza el listado para reflejar el nuevo estado.

## Tabla de Decisiones y Alternativas
| Decisión | Opción elegida | Alternativas consideradas | Justificación |
|----------|----------------|---------------------------|---------------|
| Estilos | Tailwind CSS v3 | CSS Modules, Styled Components | Clases utilitarias semánticas que mantienen los componentes concisos y fáciles de inspeccionar en clase. |
| Runner de pruebas | Jest + `ts-jest` + React Testing Library | Vitest | Familiaridad pedagógica estándar con bloques `describe/test/expect` y pruebas de integración completas del árbol `<App />`. |
| Notificación de errores | Banners de error en línea (Inline UI) | Notificaciones flotantes (Toasts) | Más simple, determinista en pruebas automatizadas y sin dependencias adicionales. |
