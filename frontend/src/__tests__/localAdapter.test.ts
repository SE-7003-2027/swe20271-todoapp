import { LocalStorageApiClient } from '../services/storage/localAdapter';

beforeAll(() => {
  Object.defineProperty(global, 'crypto', {
    value: {
      randomUUID: () => 'uuid-1234',
    },
  });
});

describe('LocalStorageApiClient', () => {
  let adapter: LocalStorageApiClient;

  beforeEach(() => {
    localStorage.clear();
    adapter = new LocalStorageApiClient();
  });

  // @spec STORAGE-ADAPTER-004
  test('createNote genera UUID y estado PENDING', async () => {
    const note = await adapter.createNote('Test nota');
    expect(note.id).toBe('uuid-1234');
    expect(note.title).toBe('Test nota');
    expect(note.status).toBe('PENDING');
    
    const stored = JSON.parse(localStorage.getItem('studynotes_v1') || '[]');
    expect(stored).toHaveLength(1);
    expect(stored[0].title).toBe('Test nota');
  });

  // @spec STORAGE-ADAPTER-005
  test('getNotes devuelve notas ordenadas descendientemente por fecha', async () => {
    const note1 = await adapter.createNote('Nota 1');
    
    const stored = JSON.parse(localStorage.getItem('studynotes_v1') || '[]');
    stored[0].createdAt = '2020-01-01T00:00:00.000Z';
    localStorage.setItem('studynotes_v1', JSON.stringify(stored));
    
    const note2 = await adapter.createNote('Nota 2');
    
    const notes = await adapter.getNotes();
    expect(notes[0].title).toBe('Nota 2'); 
    expect(notes[1].title).toBe('Nota 1'); 
  });

  // @spec STORAGE-ADAPTER-006
  test('toggleNote cambia el estado PENDING a COMPLETED y viceversa', async () => {
    const note = await adapter.createNote('Mi nota');
    expect(note.status).toBe('PENDING');
    
    const completedNote = await adapter.toggleNote(note.id);
    expect(completedNote.status).toBe('COMPLETED');
    
    const revertedNote = await adapter.toggleNote(note.id);
    expect(revertedNote.status).toBe('PENDING');
  });
});
