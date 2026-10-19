import React, { useState } from 'react';

interface NoteFormProps {
  onSubmit: (title: string) => Promise<void>;
}

// @spec NOTES-UI-001
// @spec NOTES-UI-002
// @spec NOTES-UI-003
export function NoteForm({ onSubmit }: NoteFormProps) {
  const [title, setTitle] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    
    const trimmedTitle = title.trim();

    if (trimmedTitle.length === 0) {
      setError('El título no puede estar vacío');
      return;
    }

    if (trimmedTitle.length > 100) {
      setError('El título no puede superar los 100 caracteres');
      return;
    }

    try {
      setIsSubmitting(true);
      await onSubmit(trimmedTitle);
      setTitle(''); // Limpiar la entrada en caso de éxito
    } catch (err: any) {
      setError(err.message || 'Error al crear la nota');
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mb-6 p-4 bg-white rounded-lg shadow">
      <div className="flex flex-col sm:flex-row gap-3">
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="Escribe una nueva nota..."
          disabled={isSubmitting}
          className="flex-1 p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <button
          type="submit"
          disabled={isSubmitting}
          className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 disabled:opacity-50 transition-colors"
        >
          {isSubmitting ? 'Guardando...' : 'Crear Nota'}
        </button>
      </div>
      {error && (
        <div className="mt-3 p-2 bg-red-100 text-red-700 text-sm rounded border border-red-200">
          {error}
        </div>
      )}
    </form>
  );
}
