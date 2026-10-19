import React, { useEffect, useState } from 'react';
import { Note, StorageAdapter } from './services/storage/types';
import { createAdapter, LocalStorageApiClient } from './services/storage';
import { NoteForm } from './components/NoteForm';
import { NoteList } from './components/NoteList';

export default function App() {
  const [adapter, setAdapter] = useState<StorageAdapter | null>(null);
  const [notes, setNotes] = useState<Note[]>([]);
  const [loading, setLoading] = useState(true);
  const [globalError, setGlobalError] = useState<string | null>(null);

  useEffect(() => {
    async function initAdapter() {
      try {
        const selectedAdapter = await createAdapter();
        setAdapter(selectedAdapter);
        await loadNotes(selectedAdapter);
      } catch (err: any) {
        setGlobalError('Error al iniciar la aplicación');
      } finally {
        setLoading(false);
      }
    }
    initAdapter();
  }, []);

  const loadNotes = async (currentAdapter: StorageAdapter) => {
    try {
      const fetchedNotes = await currentAdapter.getNotes();
      setNotes(fetchedNotes);
    } catch (err: any) {
      setGlobalError('Error al cargar las notas');
    }
  };

  const handleCreate = async (title: string) => {
    if (!adapter) return;
    setGlobalError(null);
    try {
      await adapter.createNote(title);
      await loadNotes(adapter);
    } catch (err: any) {
      throw err; // El componente hijo maneja su propio error
    }
  };

  const handleToggle = async (id: string) => {
    if (!adapter) return;
    setGlobalError(null);
    try {
      await adapter.toggleNote(id);
      await loadNotes(adapter);
    } catch (err: any) {
      setGlobalError('Error al actualizar la nota');
    }
  };

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-100">
        <p className="text-gray-600">Cargando aplicación...</p>
      </div>
    );
  }

  const isLocal = adapter instanceof LocalStorageApiClient;

  return (
    <div className="min-h-screen bg-gray-100 py-8 px-4">
      <div className="max-w-2xl mx-auto">
        <header className="mb-8 text-center">
          <h1 className="text-3xl font-bold text-gray-800">StudyNotes</h1>
          <p className="text-gray-500 mt-2">Gestiona tus tareas escolares</p>
        </header>

        {isLocal && (
          <div className="mb-4 p-3 bg-yellow-50 text-yellow-800 rounded border border-yellow-200 text-sm text-center">
            Modo sin conexión – los datos se guardan localmente
          </div>
        )}

        {globalError && (
          <div className="mb-4 p-3 bg-red-100 text-red-800 rounded border border-red-200 text-center">
            {globalError}
          </div>
        )}

        <NoteForm onSubmit={handleCreate} />
        
        <main>
          <NoteList notes={notes} onToggle={handleToggle} />
        </main>
      </div>
    </div>
  );
}
