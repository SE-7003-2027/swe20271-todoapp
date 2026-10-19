// Definiciones de tipos para el almacenamiento de notas

export type NoteStatus = 'PENDING' | 'COMPLETED';

export interface Note {
  id: string;
  title: string;
  status: NoteStatus;
  createdAt: string; // Formato ISO 8601
}

export interface StorageAdapter {
  getNotes(): Promise<Note[]>;
  createNote(title: string): Promise<Note>;
  toggleNote(id: string): Promise<Note>;
}
