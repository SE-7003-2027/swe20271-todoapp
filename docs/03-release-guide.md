# Guía de Liberación — StudyNotes / TaskSync

Esta guía establece el proceso estandarizado y los puntos de verificación requeridos para realizar el lanzamiento oficial de nuevas versiones del producto base hacia los ambientes correspondientes de la v0.1.0 inicial.

## 1. Definición de Terminado (DoD)
Para certificar una nueva característica como apta de liberación y terminada para entrega final debe estrictamente acatar estas reglas:
- [ ] Todas las pruebas asociadas al proyecto, que verifiquen procesos unitarios y validen el comportamiento esperado (end-to-end), transicionan como exitosas garantizando estabilidad y cubrimiento.
- [ ] Las ejecuciones asociadas a la generación de archivos finales de uso distribuido del cliente y código general de servidor transcurren ininterrumpidamente originando salidas íntegras sin marcación alguna de problemas y alertas de *build*.
- [ ] Los recursos asociados como descripciones escritas acerca del diseño actual, pautas documentadas de requisitos funcionales y diagramas base están sincrónicamente reflejando los estados vigentes y presentes en código recién alterado.
- [ ] El registro en los bitácoras bajo formato `CHANGELOG.md` provee resúmenes que ilustran a detalle los sucesos recientes de mejoras a disposición.
- [ ] Todo requisito validado con la sintaxis `@spec` implementada en los EARS concuerdan en precisión y estado validado junto con el estado del desarrollo a la par.

## 2. Flujo de Liberación con Git
El método preciso a ejecutar en la manipulación central de versiones asume el trabajo directamente operado a través del manejador Git para lograr consistencia durante el lanzamiento en la rama original del repositorio de trabajo.

Asegúrate y localiza en consola la ruta local principal donde ejecutarás el chequeo:
1. Revisa tu posición estableciendo que la cabeza de trabajo apunte exactamente sobre la versión madre llamada *main* conteniendo todos tus registros finales:
   ```bash
   git checkout main
   git pull origin main
   ```
2. Despliega al tiempo actual un indicador inmutable con apuntes descriptivos validando los trabajos para formar esta revisión:
   ```bash
   git tag -a v0.1.0 -m "Liberación v0.1.0"
   ```
3. Distribuye esta revisión hacia el entorno colaborativo empujándola del estado de desarrollo estático actual e invocando su integración a procesos CI que formen liberaciones o builds listos automáticos:
   ```bash
   git push origin v0.1.0
   ```

## 3. Mantenimiento del CHANGELOG
En cada iteración mantenemos una disciplina transparente registrando toda actividad funcional del producto adaptando y adhiriéndose a las convenciones probadas establecidas globalmente bajo formato de "Keep a Changelog".

Ejemplo de cómo reportarlo debidamente dentro del archivo de control histórico:
```markdown
## [0.1.0] - 2023-10-15
### Añadido
- Interfaz gráfica principal moderna implementando interacciones utilizando React 18 en entorno TypeScript.
- Soporte para creación veloz con entrada controlada e indexación visible básica a la lista central de todas las tareas (FR-01, FR-02).
- Capacidad integrada permitiendo manipular su estado alternando de forma visual y de control en memoria su conclusión a finalizadas (FR-03).
- Compatibilidad absoluta integrando funcionamiento pasivo *offline* a través de un auto-fallback del adaptador empleando un guardado dentro del `LocalStorage`.
```

## 4. Configuración de GitHub Pages
Proceso unívoco manual implementado al generar la vista a presentar frente al entorno aislado si utilizas GitHub:
1. Dirígete e ingresa en la plataforma web al apartado respectivo dentro del destino de los repositorios hacia la vista general de la aplicación seleccionando **Settings** para luego identificar el menú específico **Pages**.
2. Modifica la ruta inicial donde se despliegan estos servicios a la propia rama destinada desde el flujo principal por el control como `gh-pages` con su respectiva ruta web desde un sistema automatizado de *GitHub Actions*.
3. De igual manera, no prescindas confirmar que en las variables de contexto e integridad dentro de la base general del repositorio para *Secrets*, se halla el atributo clave establecido: `VITE_USE_MOCK=true`. Eso permitirá que la versión a alojar instale y contenga exclusivamente al emulador del archivo local.

## 5. Docker
Como fase complementaria final disponemos y damos oportunidad del arranque completo garantizando interactividad conjunta real de sistemas sobre ambientes localizados. Se deberá instalar la herramienta:
1. Lanza o abre tu emulador de terminal dentro de tu propio disco y navega estrictamente hacia la base matriz raíz desde donde la base de los archivos operan en un ambiente estandarizado.
2. Inyecta este simple comando originario orquestando ambas funciones en el esquema actual sin atrapar a tu terminal y liberando los servicios de manera asíncrona:
   ```bash
   docker compose up -d
   ```
3. Verifica la normal disposición funcional confirmando los procesos navegando sobre tu explorador web e invocando:
   - Capa del cliente visual, el cual despliega el aplicativo en **Frontend:** http://localhost:5173
   - Y verifica indirectamente el entorno asíncrono para comprobación llamando al **Backend API:** http://localhost:8080/api/v1/notes (esto se presentará mediante la respuesta como un simple código estricto arrojando `[]` bajo primera lectura al arranque del programa).
