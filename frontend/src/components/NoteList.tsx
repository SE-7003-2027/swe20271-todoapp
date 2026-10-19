import React from 'react';
import { Note } from '../services/storage/types';
import { NoteItem } from './NoteItem';

interface NoteListProps {
  notes: Note[];
  onToggle: (id: string) => void;
}

// @spec NOTES-UI-006
export function NoteList({ notes, onToggle }: NoteListProps) {
  if (notes.length === 0) {
    return (
      <div className="p-8 text-center bg-gray-50 rounded-lg border border-dashed border-gray-300">
        <p className="text-gray-500">
          No hay notas todavía. ¡Crea tu primera nota arriba!
        </p>
      </div>
    );
  }

  return (
    <div className="flex flex-col gap-2">
      {notes.map((note) => (
        <NoteItem key={note.id} note={note} onToggle={onToggle} />
      ))}
    </div>
  );
}
