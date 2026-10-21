import { Note, StorageAdapter } from './types';

// Adaptador para peticiones HTTP hacia el backend en Go
export class HttpApiClient implements StorageAdapter {
  private baseUrl: string;

  constructor() {
    this.baseUrl = process.env.VITE_API_URL || 'http://localhost:8080';
  }

  // @spec STORAGE-ADAPTER-002
  async getNotes(): Promise<Note[]> {
    const response = await fetch(`${this.baseUrl}/api/v1/notes`);
    if (!response.ok) {
      throw new Error('Error al obtener las notas del servidor');
    }
    return response.json();
  }

  // @spec STORAGE-ADAPTER-002
  async createNote(title: string): Promise<Note> {
    const response = await fetch(`${this.baseUrl}/api/v1/notes`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ title }),
    });
    if (!response.ok) {
      throw new Error('Error al crear la nota en el servidor');
    }
    return response.json();
  }

  // @spec STORAGE-ADAPTER-002
  async toggleNote(id: string): Promise<Note> {
    const response = await fetch(`${this.baseUrl}/api/v1/notes/${id}/toggle`, {
      method: 'PATCH',
    });
    if (!response.ok) {
      throw new Error('Error al cambiar el estado de la nota en el servidor');
    }
    return response.json();
  }
}
