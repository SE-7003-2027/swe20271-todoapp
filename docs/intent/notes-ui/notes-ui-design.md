# LLD — Segmento: notes-ui

## Propósito
El diseño principal de la capa front-end (React) es establecer las interacciones fluídas de la interfaz gráfica y mapear consistentemente los cambios de los elementos con la instancia del Adapter inyectado. Todo debe tener un enfoque en un código base sencillo, ligero, libre de abstracciones complejas.

## Árbol de Componentes
El layout general recae en un sistema de anidación simple:
- `App` (Raíz)
  - `NoteForm` (Control para registro de tareas)
  - `NoteList` (Componente orquestador del renderizado en bucle)
    - `NoteItem` (Entidad unitaria del bloque visual para alternar un registro)

## Flujo de Estado
- **Notes State:** En la jerarquía principal y propagado como "properties" (props). Representa un arreglo de tareas recibidas por el Storage Adapter.
- **Adapter Bootstrap:** Realizado a nivel de inicio mediante hooks en el componente superior de React; emite transiciones (banderas o *flags* visuales) tales como estados de carga ("Cargando configuraciones iniciales").
- **Estados de Carga y Error:** Se centralizan en mecanismos directos de interrupción visual, deshabilitando todos los elementos de formulario activos y bloqueando clics hasta finalizar la respuesta correspondiente al Adapter para eludir transacciones duplicadas.

## Validación del Cliente
- **Limpieza (Trim) de Cadenas:** Toda captura antes del procesamiento retira prefijos de espacios inútiles en ambos extremos de la descripción para evaluar adecuadamente sus propiedades antes del paso por red.
- **Reglas del límite de caracteres (100) y longitud cero:** Limitantes estrictas; sus alertas de validación correspondientes en el documento reaccionan si se incumple alguna regla informando visualmente de inmediato. Textos esperados y mostrados puramente en idioma español: "El título no puede estar vacío" y "El título excede el límite de 100 caracteres".

## Tabla de Decisiones y Alternativas
| Decisión | Opción elegida | Alternativas consideradas | Justificación |
|----------|----------------|---------------------------|---------------|
| Motor CSS y estilos | Tailwind CSS | CSS Modules, SASS, Global Styles | Acelera la etapa visual drásticamente usando clases utilitarias predecibles. Minimiza la fricción para nombramientos. |
| Suite de pruebas del Front-end | Vitest | Jest | Integración automática, robusta y optimizada si compartimos todo el andamiaje del ecosistema bajo herramientas Vite. |
| Notificación de alertas | Errores en línea directamente (Inline UI) | Mensajes flotantes asíncronos (Toasts) | Incrementa la concisión y minimalismo esperado; presenta notificaciones localizadas a lo largo del flujo original del usuario y no satura su área global de pantalla. |
