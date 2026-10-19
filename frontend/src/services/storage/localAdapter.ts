import { Note, StorageAdapter } from './types';

const STORAGE_KEY = 'studynotes_v1';

// Adaptador para usar almacenamiento local
export class LocalStorageApiClient implements StorageAdapter {
  
  // @spec STORAGE-ADAPTER-005
  async getNotes(): Promise<Note[]> {
    const data = localStorage.getItem(STORAGE_KEY);
    if (!data) return [];
    
    try {
      const notes: Note[] = JSON.parse(data);
      // Ordena por fecha de creación descendente
      return notes.sort((a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime());
    } catch (error) {
      console.error('Error al parsear las notas locales', error);
      return [];
    }
  }

  // @spec STORAGE-ADAPTER-004
  async createNote(title: string): Promise<Note> {
    const notes = await this.getNotes();
    
    const newNote: Note = {
      id: crypto.randomUUID(), // Genera un UUID v4 en el lado del cliente
      title,
      status: 'PENDING',
      createdAt: new Date().toISOString(),
    };
    
    notes.push(newNote);
    localStorage.setItem(STORAGE_KEY, JSON.stringify(notes));
    return newNote;
  }

  // @spec STORAGE-ADAPTER-006
  async toggleNote(id: string): Promise<Note> {
    const notes = await this.getNotes();
    const noteIndex = notes.findIndex(n => n.id === id);
    
    if (noteIndex === -1) {
      throw new Error('Nota no encontrada');
    }
    
    const note = notes[noteIndex];
    note.status = note.status === 'PENDING' ? 'COMPLETED' : 'PENDING';
    
    localStorage.setItem(STORAGE_KEY, JSON.stringify(notes));
    return note;
  }
}
