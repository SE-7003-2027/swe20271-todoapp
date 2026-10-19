import React from 'react';
import { Note } from '../services/storage/types';

interface NoteItemProps {
  note: Note;
  onToggle: (id: string) => void;
}

// @spec NOTES-UI-004
// @spec NOTES-UI-005
export function NoteItem({ note, onToggle }: NoteItemProps) {
  const isCompleted = note.status === 'COMPLETED';
  
  // Opciones explícitas compatibles con Node/jsdom y navegadores modernos
  const formattedDate = new Date(note.createdAt).toLocaleString('es-MX', {
    day: '2-digit',
    month: 'short',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  });

  return (
    <div className={`p-4 mb-3 border rounded-lg shadow-sm flex items-center justify-between transition-colors ${isCompleted ? 'bg-gray-50 border-gray-200' : 'bg-white border-gray-300'}`}>
      <div className="flex flex-col">
        <h3 className={`text-lg font-medium ${isCompleted ? 'text-gray-500 line-through' : 'text-gray-800'}`}>
          {note.title}
        </h3>
        <span className="text-xs text-gray-500 mt-1">Creado: {formattedDate}</span>
        
        <div className="mt-2">
          <span className={`text-xs font-semibold px-2 py-1 rounded-full ${isCompleted ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'}`}>
            {isCompleted ? 'COMPLETADO' : 'PENDIENTE'}
          </span>
        </div>
      </div>
      
      <button
        onClick={() => onToggle(note.id)}
        className={`ml-4 px-3 py-1 text-sm font-medium rounded transition-colors ${
          isCompleted 
            ? 'text-gray-600 bg-gray-200 hover:bg-gray-300' 
            : 'text-white bg-green-600 hover:bg-green-700'
        }`}
      >
        {isCompleted ? 'Reactivar' : 'Completar'}
      </button>
    </div>
  );
}
