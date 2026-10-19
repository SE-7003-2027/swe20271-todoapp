// Configuración de Jest para StudyNotes / TaskSync
module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'jsdom',
  // Cada archivo de prueba importa '@testing-library/jest-dom' directamente
  moduleNameMapper: {
    '\\.(css|less|scss|sass)$': 'identity-obj-proxy',
  },
  transform: {
    '^.+\\.(ts|tsx)$': ['ts-jest', {
      tsconfig: 'tsconfig.json'
    }]
  }
};

