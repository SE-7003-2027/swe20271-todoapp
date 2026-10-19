import { HttpApiClient } from './httpAdapter';
import { LocalStorageApiClient } from './localAdapter';
import { StorageAdapter } from './types';

export { HttpApiClient, LocalStorageApiClient };
export type { StorageAdapter };

// @spec STORAGE-ADAPTER-001
// @spec STORAGE-ADAPTER-002
// @spec STORAGE-ADAPTER-003
export async function createAdapter(): Promise<StorageAdapter> {
  // Si estamos en modo de prueba forzado
  if (import.meta.env.VITE_USE_MOCK === 'true') {
    return new LocalStorageApiClient();
  }

  const baseUrl = import.meta.env.VITE_API_URL ?? 'http://localhost:8080';

  try {
    const response = await fetch(`${baseUrl}/healthz`, {
      signal: AbortSignal.timeout(2000), // Límite de tiempo de 2000ms
    });

    if (response.ok) {
      return new HttpApiClient();
    }
  } catch (error) {
    console.warn('El backend no respondió a tiempo, usando almacenamiento local.');
  }

  // Si falló o excedió el tiempo
  return new LocalStorageApiClient();
}
