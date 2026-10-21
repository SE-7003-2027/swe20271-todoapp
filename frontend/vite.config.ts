// Configuración de Vite
import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

declare const process: { env: Record<string, string | undefined> };

export default defineConfig({
  base: './',
  plugins: [react()],
  define: {
    'process.env.VITE_USE_MOCK': JSON.stringify(process.env.VITE_USE_MOCK ?? ''),
    'process.env.VITE_API_URL': JSON.stringify(process.env.VITE_API_URL ?? ''),
  },
});
