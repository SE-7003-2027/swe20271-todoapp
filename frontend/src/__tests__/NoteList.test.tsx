import React from 'react';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom';
import { NoteList } from '../components/NoteList';
import { Note } from '../services/storage/types';

describe('NoteList', () => {
  // @spec NOTES-UI-006
  test('muestra mensaje de estado vacío cuando no hay notas', () => {
    render(<NoteList notes={[]} onToggle={() => {}} />);
    expect(screen.getByText('No hay notas todavía. ¡Crea tu primera nota arriba!')).toBeInTheDocument();
  });

  // @spec NOTES-UI-004
  test('renderiza notas con estado correcto', () => {
    const notes: Note[] = [
      { id: '1', title: 'Comprar leche', status: 'PENDING', createdAt: new Date().toISOString() },
      { id: '2', title: 'Hacer tarea', status: 'COMPLETED', createdAt: new Date().toISOString() },
    ];
    
    render(<NoteList notes={notes} onToggle={() => {}} />);
    
    expect(screen.getByText('Comprar leche')).toBeInTheDocument();
    expect(screen.getByText('Hacer tarea')).toBeInTheDocument();
    
    expect(screen.getByText('PENDIENTE')).toBeInTheDocument();
    expect(screen.getByText('COMPLETADO')).toBeInTheDocument();
  });
});
