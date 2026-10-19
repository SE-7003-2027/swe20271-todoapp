# LLD — Segmento: storage-adapter

## Propósito
El propósito primario del adaptador de almacenamiento es el desacoplamiento de la interfaz gráfica y los estados subyacentes frente a la verdadera fuente de obtención de datos. Esto protege a los componentes de UI ante posibles caídas, migraciones de base de datos e inclusive entornos estáticos que carecen de backend propio.

## Interfaz StorageAdapter
La siguiente interfaz escrita en TypeScript rige el contrato que deben satisfacer todas las implementaciones:
```typescript
export interface StorageAdapter {
  getNotes(): Promise<Note[]>;
  createNote(title: string): Promise<Note>;
  toggleNoteStatus(id: string): Promise<Note>;
}
```

## Estrategia de Selección de Adaptador
1. Al momento de la inicialización de la SPA, se lee la variable de entorno `VITE_USE_MOCK=true`.
2. Si esta configuración está encendida, o si la verificación de salud inicial asíncrona hacia el servidor REST falla de manera inequívoca, se decide instanciar `LocalStorageApiClient`.
3. Si el parámetro en el entorno no fuerza un mockeo, y el backend responde en el healthcheck inicial, entonces el `HttpApiClient` es activado.

## LocalStorageApiClient
- **Clave del caché:** Se utiliza la clave predefinida `studynotes_data_v1` en el LocalStorage.
- **Formato de datos:** Corresponde a una matriz codificada en formato JSON representando arreglos de objetos `Note`.
- **Ordenamiento:** La estrategia es ordenar cada lectura de datos para presentar inicialmente el de fecha de creación más reciente (descendente cronológicamente), garantizando la consistencia visual esperada en `getNotes()`.

## HttpApiClient
- **URLs Base de destino:** Dictado por `VITE_API_BASE_URL` (por default `/api/v1`).
- **Manejo de errores HTTP:** Transforma los errores (especialmente validaciones `400` u otros `500`) devolviendo excepciones fuertemente tipadas y formateadas hacia el nivel superior a fin de que la interfaz notifique al usuario.

## Tabla de Decisiones y Alternativas
| Decisión | Opción elegida | Alternativas consideradas | Justificación |
|----------|----------------|---------------------------|---------------|
| Selección de adaptador | Autodetección o entorno estricto | Exclusivo por variable de entorno | Ofrece flexibilidad y permite el "cero-configuración", al tiempo que garantiza el estado offline para entornos que fallan sus conexiones. |
| Generación de UUID local | La API estándar `crypto.randomUUID()` | Paquete NPM o de terceros | Disponible en todos los navegadores modernos y navegadores seguros; esto evita sobrecargar el bundle resultante. |
| Clave en Storage | Uso de convención versionada (ej. `_v1`) | Clave corta sin revisión (`notes`) | Garantiza la mitigación de errores de deserialización cuando la interfaz cambie radicalmente en versiones futuras. |
