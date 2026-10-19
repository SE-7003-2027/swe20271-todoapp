# Especificación de Requisitos — StudyNotes / TaskSync v0.1.0

## 1. Espacio del Problema vs. Espacio de la Solución
**Espacio del Problema:** En el entorno universitario y de aprendizaje, los estudiantes y educadores enfrentan diariamente la sobrecarga de información y necesitan una forma extremadamente rápida, libre de fricción y sin distracciones para anotar tareas pendientes y notas breves. Las herramientas existentes en el mercado usualmente tienen demasiadas opciones configurables (asignación de subtareas complejas, fechas múltiples, sistema estricto de prioridades) que terminan abrumando al usuario y agregan una constante carga cognitiva indeseable que disuade del uso.
**Espacio de la Solución:** Se pretende proporcionar una aplicación web sumamente minimalista, enfocada en la creación y administración de notas a modo de una lista de tareas única. El usuario ingresa a la aplicación web, escribe sin problemas una nota concreta de hasta 100 caracteres, y tiene la potestad de marcarla como completada cuando termine. Fin del flujo funcional.

## 2. Taxonomía de Requisitos
### Funcionales
- **FR-01 (Crear Nota):** El sistema debe permitir a los usuarios crear notas o tareas breves de texto simple de forma rápida a través de la interfaz.
- **FR-02 (Listar Notas):** El sistema deberá presentar todas las notas introducidas previamente por el usuario, priorizando aquellas introducidas más recientemente en el listado visual principal.
- **FR-03 (Alternar Estado):** El sistema debe posibilitar al usuario mediante una interacción marcar las notas visualizadas como "completadas" o regresar su estado a "pendientes".

### No Funcionales
- **Latencia:** Todas las operaciones y respuestas primarias de visualización deben verse y sentirse efectivas frente a la experiencia, procesándose siempre en rangos menores a 200 milisegundos.
- **Portabilidad Estática:** Toda la representación del frontend de interfaz debe operar de una manera sumamente independiente. Se debe soportar totalmente ambientes estáticos apoyándose en persistencia local si el servidor backend principal resultase no alcanzable.
- **Onboarding Cero-Configuración:** La puesta en marcha (deploy) o arranque en entornos de trabajo del desarrollador se da bajo pasos limitados donde bases de datos complejas de terceros sean opcionales y no mandatarias del setup primario.

### Restricciones
- La solución deberá compilar enteramente y permitir fácil despliegue en un modelo estático y tradicional de hospedaje web como GitHub Pages sin requerir middleware intermedio en dicha modalidad.
- Dentro de toda esta fase v0.1.0 inicial, el uso de implementaciones o conectores reales de base de datos como MySQL, SQLite o PostgreSQL queda prohibido y fuera de este alcance.

## 3. Historias de Usuario con Criterios de Aceptación Gherkin

### FR-01: Crear Nota
**Historia de usuario:** Yo como estudiante, quiero crear una nota breve para no olvidar una tarea.
**Criterios de Aceptación:**
- **Escenario: Creación exitosa**
  Dado que me encuentro en la aplicación de forma activa
  Cuando ingreso el texto "Estudiar cálculo" en el formulario y lo envío para su registro
  Entonces la nota recién agregada aparece de inmediato en la parte superior de la lista activa.
- **Escenario: Título vacío en el intento**
  Dado que interactúo con la plataforma y el campo de texto se mantiene en blanco
  Cuando intento presionar de todas formas el botón del formulario
  Entonces el proceso finaliza impidiendo una creación en sistema y se me muestra un error explicativo directo de validación visualmente.

### FR-02: Listar Notas
**Historia de usuario:** Yo como estudiante, quiero ver todas mis tareas anotadas para organizar mi día de mejor manera.
**Criterios de Aceptación:**
- **Escenario: Lista poblada**
  Dado que el historial ya contiene tareas previamente creadas: las notas "A", "B" y por último "C"
  Cuando inicio y abro la aplicación en mi explorador actual
  Entonces se recaba toda la data y observo un listado ordenado mostrando en el tope la nota "C", seguida de "B", y terminando abajo con la nota "A".

### FR-03: Alternar Estado
**Historia de usuario:** Yo como estudiante, quiero marcar una nota específica como completada para poder llevar así control ordenado de mis avances totales.
**Criterios de Aceptación:**
- **Escenario: Completar una tarea en curso**
  Dado que la vista lista al menos una nota o elemento con estado dependiente pendiente
  Cuando decido ejecutar un clic sobre su propio botón o checkbox de estado visual
  Entonces visualizo que su apariencia y estado interno de control ahora se designan como completada exitosamente.
- **Escenario: Desmarcar de terminada a pendiente**
  Dado que en lista percibo una nota cuya tarea se marca como completada hace un tiempo
  Cuando yo acciono nuevamente de forma intencional su estado de palomeo
  Entonces todo se reinvierte y de nuevo percibo el reactivo marcando el registro de la nota a su origen pendiente.

## 4. Catálogo de Casos Límite
El entorno toma acciones correctivas automáticas ante estos conflictos de bordes:
- **Título vacío por relleno de espacios:** Todas las entradas sufren un proceso global de saneamiento `trim()`. Si restan cadenas vacías como resultado, estas fracasarán las validaciones actuando en similitud como un envío no capturado o puramente en blanco.
- **Título con un peso exacto a 100 caracteres:** Este texto largo se evaluará sin rechazos de la infraestructura asumiendo una completitud precisa que transiciona como entrada válida sin excepción alguna a todos sus sistemas y a su almacén receptor en memoria.
- **Título por sobre los límites permitidos:** Todo documento constando un total de 101 caracteres o más fracasa categóricamente en ambos lados: desde el bloqueo forzoso por parámetros HTML en el lado del navegador del cliente y la consiguiente validación `400` directa controlada tras fallar en la capa de servicios por servidor (para quienes operan sin los checks visuales frontales).
- **Toggles masivos repetitivos sobre componentes o clicks apresurados de los listados concurrentes:** Toda interfaz previene y limita que repetitivas transacciones en simultáneo inunden las filas del control adaptativo deteniendo peticiones descarriadas y estableciendo temporalizaciones o bloqueos en la interfaz ("UI") frente a estados temporales asíncronos que paralicen su interacción hasta estar resueltos.
- **Backend indisponible al momento del arranque general:** Si un fallo mayor impidiese total acceso inicial a la infraestructura que corre en `localhost:8080`, el orquestador principal delegará al adaptador StorageAdapter invocar funciones secundarias del `LocalStorageApiClient` capturando su incapacidad y evitando bloquear o tumbar irreversiblemente la experiencia para todo usuario final en navegación sin conexión.
- **Cuotas agotadas del cache web:** Una acumulación gigantesca en las entradas de cliente y posterior alerta de llenado por límite estricto del espacio en navegador `QuotaExceededError` detona interrupciones notificando fallas catastróficas del almacén obligando la revisión técnica del entorno temporal para liberar peso o forzar migraciones necesarias hacia otro origen masivo de sistema.

## 5. Alcance Diferido (v0.2.0 backlog)
Esta propuesta original restringe del código elementos no prioritarios o en evolución de fases posteriores para no aumentar los tiempos:
- **FR-04**: Operaciones para lograr la eliminación real en sistema base de notas creadas.
- **FR-05**: Operadores directos sobre etiquetas que facultarían interactuar cambiando textos pre-creados de un título a otro con interacciones en tiempo real usando el frontend "inline".
- **FR-06**: Propiedades o atributos visuales complementarios agregando tags funcionales de texto.
- Implementación total para autenticaciones centralizadas requiriendo log-ins de validación y posterior despliegue formal de configuraciones sincrónicas online de forma universal a equipos paralelos de un mismo actor principal o titular del usuario en bases como sistema de autenticación persistente con servicios en SQL.
