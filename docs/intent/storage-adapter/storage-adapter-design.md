# LLD — Segmento: storage-adapter

## Propósito
El propósito primario del adaptador de almacenamiento es desacoplar los componentes visuales de React de la fuente de datos subyacente. Esto permite que la aplicación opere indistintamente contra el backend real en Go (`HttpApiClient`) o de forma autónoma en el navegador (`LocalStorageApiClient`).

## Interfaz StorageAdapter
Definida en `frontend/src/services/storage/types.ts`:
```typescript
export type NoteStatus = 'PENDING' | 'COMPLETED';

export interface Note {
  id: string;
  title: string;
  status: NoteStatus;
  createdAt: string; // ISO 8601
}

export interface StorageAdapter {
  getNotes(): Promise<Note[]>;
  createNote(title: string): Promise<Note>;
  toggleNote(id: string): Promise<Note>;
}
```

## Estrategia de Selección de Adaptador (`createAdapter`)
1. Al iniciar la aplicación, `createAdapter()` en `frontend/src/services/storage/index.ts` evalúa primero si `import.meta.env.VITE_USE_MOCK === 'true'`. Si es verdadero, devuelve `new LocalStorageApiClient()` inmediatamente sin realizar llamadas de red.
2. En caso contrario, realiza una petición `GET` a `${baseUrl}/healthz` (donde `baseUrl` es `import.meta.env.VITE_API_URL ?? 'http://localhost:8080'`) con un tiempo límite de 2000 ms (`AbortSignal.timeout(2000)`).
3. Si la respuesta es `200 OK`, instancia y devuelve `new HttpApiClient()`.
4. Si la petición falla o excede el tiempo de espera, captura la excepción y devuelve `new LocalStorageApiClient()`.

## LocalStorageApiClient
- **Clave de almacenamiento:** `studynotes_v1`.
- **Formato de datos:** Arreglo JSON de objetos `Note`.
- **Ordenamiento:** `getNotes()` ordena las notas por `createdAt` descendente (más recientes primero).

## HttpApiClient
- **URL Base:** `import.meta.env.VITE_API_URL ?? 'http://localhost:8080'`.
- **Operaciones:** Consume `GET /api/v1/notes`, `POST /api/v1/notes` y `PATCH /api/v1/notes/${id}/toggle`, lanzando un `Error` con mensaje en español cuando la respuesta HTTP no es exitosa (`!response.ok`).

## Tabla de Decisiones y Alternativas
| Decisión | Opción elegida | Alternativas consideradas | Justificación |
|----------|----------------|---------------------------|---------------|
| Selección de adaptador | Autodetección con `/healthz` (timeout 2s) + bandera `VITE_USE_MOCK` | Solo variable de entorno en compilación | Permite que la app funcione sin configuración tanto con el backend encendido como sin conexión, y evita esperar 2s en despliegues estáticos como GitHub Pages donde `VITE_USE_MOCK=true`. |
| Generación de UUID local | `crypto.randomUUID()` con respaldo matemático | Librería npm `uuid` | Soportado nativamente en navegadores y Node 20 sin incrementar el tamaño del bundle. |
| Clave en `localStorage` | `studynotes_v1` | `notes` | El sufijo de versión facilita migraciones limpias del esquema local en futuras entregas. |
