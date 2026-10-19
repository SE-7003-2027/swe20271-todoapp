import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import '@testing-library/jest-dom';
import { NoteForm } from '../components/NoteForm';

describe('NoteForm', () => {
  // @spec NOTES-UI-001
  test('muestra error si el título está vacío y no llama al manejador', async () => {
    const mockSubmit = jest.fn();
    const user = userEvent.setup();
    render(<NoteForm onSubmit={mockSubmit} />);

    const button = screen.getByRole('button', { name: /Crear Nota/i });
    await user.click(button);

    expect(await screen.findByText('El título no puede estar vacío')).toBeInTheDocument();
    expect(mockSubmit).not.toHaveBeenCalled();
  });

  // @spec NOTES-UI-002
  test('muestra error si el título supera los 100 caracteres', async () => {
    const mockSubmit = jest.fn();
    const user = userEvent.setup();
    render(<NoteForm onSubmit={mockSubmit} />);

    const input = screen.getByPlaceholderText('Escribe una nueva nota...');
    const longText = 'a'.repeat(101);
    await user.type(input, longText);

    const button = screen.getByRole('button', { name: /Crear Nota/i });
    await user.click(button);

    expect(await screen.findByText('El título no puede superar los 100 caracteres')).toBeInTheDocument();
    expect(mockSubmit).not.toHaveBeenCalled();
  });

  // @spec NOTES-UI-003
  test('llama al manejador y borra la entrada en caso de éxito', async () => {
    const mockSubmit = jest.fn().mockResolvedValue(undefined);
    const user = userEvent.setup();
    render(<NoteForm onSubmit={mockSubmit} />);

    const input = screen.getByPlaceholderText('Escribe una nueva nota...');
    await user.type(input, 'Mi nueva tarea');

    const button = screen.getByRole('button', { name: /Crear Nota/i });
    await user.click(button);

    expect(mockSubmit).toHaveBeenCalledWith('Mi nueva tarea');

    // Esperar a que el estado asíncrono se actualice antes de verificar el vaciado
    await waitFor(() => {
      expect(input).toHaveValue('');
    });
  });

  test('muestra error si el envío falla', async () => {
    const mockSubmit = jest.fn().mockRejectedValue(new Error('Fallo del servidor'));
    const user = userEvent.setup();
    render(<NoteForm onSubmit={mockSubmit} />);

    const input = screen.getByPlaceholderText('Escribe una nueva nota...');
    await user.type(input, 'Mi nueva tarea');

    const button = screen.getByRole('button', { name: /Crear Nota/i });
    await user.click(button);

    expect(await screen.findByText('Fallo del servidor')).toBeInTheDocument();
  });
});
