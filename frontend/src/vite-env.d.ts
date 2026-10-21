/// <reference types="vite/client" />

declare const process: {
  env: {
    VITE_USE_MOCK?: string;
    VITE_API_URL?: string;
    [key: string]: string | undefined;
  };
};
