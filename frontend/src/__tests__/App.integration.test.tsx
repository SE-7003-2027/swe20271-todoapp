import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import '@testing-library/jest-dom';
import App from '../App';

describe('Integración Frontend — App + NoteForm + NoteList + StorageAdapter', () => {
  const originalFetch = global.fetch;

  beforeEach(() => {
    localStorage.clear();
  });

  afterEach(() => {
    global.fetch = originalFetch;
    jest.restoreAllMocks();
  });

  // @spec STORAGE-ADAPTER-003
  // @spec NOTES-UI-005
  // @spec NOTES-UI-007
  test('flujo completo con respaldo LocalStorage: carga estado vacío, crea nota y alterna su estado a COMPLETADO', async () => {
    // Simular que el backend no está disponible para activar LocalStorageApiClient automáticamente
    global.fetch = jest.fn().mockRejectedValue(new Error('Backend inaccesible'));
    const user = userEvent.setup();

    render(<App />);

    // 1. Esperar a que termine la autodetección de salud y muestre el modo sin conexión y la lista vacía
    expect(await screen.findByText(/Modo sin conexión/i)).toBeInTheDocument();
    expect(screen.getByText('No hay notas todavía. ¡Crea tu primera nota arriba!')).toBeInTheDocument();

    // 2. Crear una nota desde NoteForm
    const input = screen.getByPlaceholderText('Escribe una nueva nota...');
    await user.type(input, 'Preparar laboratorio de Jenkins');
    await user.click(screen.getByRole('button', { name: /Crear Nota/i }));

    // 3. Verificar que la nota aparece en NoteList con estado PENDIENTE y el input se limpió
    expect(await screen.findByText('Preparar laboratorio de Jenkins')).toBeInTheDocument();
    expect(screen.getByText('PENDIENTE')).toBeInTheDocument();
    expect(input).toHaveValue('');

    // 4. Hacer clic en "Completar" dentro de NoteItem y verificar la transición a COMPLETADO
    const toggleBtn = screen.getByRole('button', { name: /Completar/i });
    await user.click(toggleBtn);

    await waitFor(() => {
      expect(screen.getByText('COMPLETADO')).toBeInTheDocument();
    });
    expect(screen.getByRole('button', { name: /Reactivar/i })).toBeInTheDocument();
  });

  // @spec STORAGE-ADAPTER-001
  // @spec STORAGE-ADAPTER-002
  // @spec NOTES-UI-007
  test('flujo completo conectado al Backend HTTP (HttpApiClient): salud OK, listar, crear y alternar', async () => {
    const serverNotes: Array<{ id: string; title: string; status: 'PENDING' | 'COMPLETED'; createdAt: string }> = [];

    global.fetch = jest.fn().mockImplementation(async (input: RequestInfo | URL, init?: RequestInit) => {
      const url = String(input);
      const method = init?.method ?? 'GET';

      if (url.endsWith('/healthz')) {
        return { ok: true, status: 200, json: async () => ({ status: 'ok' }) } as Response;
      }
      if (url.endsWith('/api/v1/notes') && method === 'GET') {
        return { ok: true, status: 200, json: async () => [...serverNotes] } as Response;
      }
      if (url.endsWith('/api/v1/notes') && method === 'POST') {
        const body = JSON.parse(String(init?.body ?? '{}'));
        const created = {
          id: 'uuid-servidor-1',
          title: body.title,
          status: 'PENDING' as const,
          createdAt: new Date().toISOString(),
        };
        serverNotes.unshift(created);
        return { ok: true, status: 201, json: async () => created } as Response;
      }
      if (url.includes('/api/v1/notes/uuid-servidor-1/toggle') && method === 'PATCH') {
        serverNotes[0].status = serverNotes[0].status === 'PENDING' ? 'COMPLETED' : 'PENDING';
        return { ok: true, status: 200, json: async () => serverNotes[0] } as Response;
      }
      return { ok: false, status: 404, json: async () => ({ error: 'No encontrado' }) } as Response;
    });

    const user = userEvent.setup();
    render(<App />);

    // 1. Verificar que NO aparece el aviso de modo sin conexión porque /healthz respondió 200 OK
    expect(await screen.findByText('No hay notas todavía. ¡Crea tu primera nota arriba!')).toBeInTheDocument();
    expect(screen.queryByText(/Modo sin conexión/i)).not.toBeInTheDocument();

    // 2. Crear una nota a través de HttpApiClient
    const input = screen.getByPlaceholderText('Escribe una nueva nota...');
    await user.type(input, 'Nota persistida en servidor Go');
    await user.click(screen.getByRole('button', { name: /Crear Nota/i }));

    expect(await screen.findByText('Nota persistida en servidor Go')).toBeInTheDocument();
    expect(screen.getByText('PENDIENTE')).toBeInTheDocument();

    // 3. Alternar estado a través de PATCH /api/v1/notes/{id}/toggle
    await user.click(screen.getByRole('button', { name: /Completar/i }));

    await waitFor(() => {
      expect(screen.getByText('COMPLETADO')).toBeInTheDocument();
    });
  });
});
