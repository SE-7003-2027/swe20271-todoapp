import { HttpApiClient } from './httpAdapter';
import { LocalStorageApiClient } from './localAdapter';
import { StorageAdapter } from './types';

export { HttpApiClient, LocalStorageApiClient };
export type { StorageAdapter };

// @spec STORAGE-ADAPTER-001
// @spec STORAGE-ADAPTER-002
// @spec STORAGE-ADAPTER-003
export async function createAdapter(): Promise<StorageAdapter> {
  // Si estamos en modo sin conexión forzado por configuración
  if (process.env.VITE_USE_MOCK === 'true') {
    return new LocalStorageApiClient();
  }

  const baseUrl = process.env.VITE_API_URL || 'http://localhost:8080';
  const controller = new AbortController();
  const timeoutId = setTimeout(() => controller.abort(), 2000);

  try {
    const response = await fetch(`${baseUrl}/healthz`, {
      signal: controller.signal, // Límite de tiempo de 2000ms compatible con todos los entornos
    });

    if (response.ok) {
      return new HttpApiClient();
    }
  } catch (error) {
    console.warn('El backend no respondió a tiempo, usando almacenamiento local.');
  } finally {
    clearTimeout(timeoutId);
  }

  // Si falló o excedió el tiempo
  return new LocalStorageApiClient();
}
