# Diseño de Alto Nivel — StudyNotes / TaskSync

## Contexto y Propósito
El problema que resuelve este proyecto es la falta de una herramienta minimalista para gestionar notas y tareas enfocada en entornos educativos, sin la sobrecarga de características complejas que abruman al estudiante. Su objetivo educativo primario es demostrar los principios de diseño de software usando el paradigma de Desarrollo Basado en Intenciones Enlazadas (LID).

## Arquitectura del Sistema
La arquitectura es un monolito por capas estructurado de la siguiente forma:
- **Handler**: Maneja las solicitudes HTTP.
- **Service**: Contiene la lógica de negocio y las validaciones correspondientes.
- **Repository**: Administra el acceso a datos y persistencia.

En el frontend, se utiliza el patrón Adapter para aislar la interfaz de usuario de las fuentes de datos. La topología del proyecto incluye soporte para Docker Compose, facilitando la orquestación y el despliegue local de la aplicación de manera integral.

## Tenets
1. **Mínima complejidad sobre abstracción prematura:** Se prioriza código simple, legible y directo en lugar de abstracciones teóricas que no proporcionan valor inmediato para una versión temprana.
2. **La demo debe funcionar sin conexión:** El `LocalStorageAdapter` es considerado un modo de primera clase y no una alternativa de contingencia (fallback). El producto está diseñado para operar directa y eficientemente desde el navegador en cualquier escenario.
3. **La documentación es código:** Todo lo referente a especificaciones EARS, diseños de bajo nivel (LLD) y pruebas funcionales se entregan junto con la base del código fuente en su versión de control correspondiente.

## Índice de Segmentos LLD
| Segmento LLD | Descripción |
|--------------|-------------|
| [notes-api](./intent/notes-api/notes-api-design.md) | Diseño del backend en Go (Endpoints, Modelo de Dominio, Servicios) |
| [storage-adapter](./intent/storage-adapter/storage-adapter-design.md) | Adaptadores de persistencia y consumo de datos para aislar la UI |
| [notes-ui](./intent/notes-ui/notes-ui-design.md) | Componentes React y gestión del estado de la interfaz gráfica |

## Decisiones Arquitectónicas
- **ADR-001: Monolito en capas (vs. microservicios):** Se eligió un diseño de monolito por su valor pedagógico. La simplicidad del mismo permite entender adecuadamente la separación de responsabilidades sin enfrentar la alta sobrecarga operativa y de configuración de los microservicios.
- **ADR-002: Portabilidad estática:** Se utiliza el patrón adapter en el frontend para asegurar que este se pueda construir y desplegar puramente como un sitio estático (ideal para GitHub Pages), conmutando al modo local cuando un backend servidor no esté disponible.

## Fuera del Alcance v0.1.0
Las siguientes funciones se han postergado de esta iteración:
- Eliminación de notas (FR-04).
- Edición inline de notas previamente creadas (FR-05).
- Etiquetas y sistema de categorización (FR-06).
- Autenticación y cuentas de usuario.
