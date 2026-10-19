// =============================================================================
// Jenkinsfile — Pipeline Declarativo de CI/CD para StudyNotes / TaskSync
// Propósito educativo: Demostrar la integración continua, pruebas automatizadas,
// compilación del backend en Go y empaquetado del frontend en React/Vite.
// =============================================================================

pipeline {
    agent any

    options {
        // Agregar marcas de tiempo a la salida de consola para facilitar la depuración
        timestamps()
        // Evitar ejecuciones concurrentes conflictivas en la misma rama
        disableConcurrentBuilds()
    }

    environment {
        // En el empaquetado estático activamos el adaptador LocalStorage por defecto
        VITE_USE_MOCK = 'true'
        CGO_ENABLED   = '0'
    }

    stages {
        stage('1. Diagnóstico del Entorno') {
            steps {
                echo '▶ Verificando herramientas instaladas en el agente de Jenkins...'
                sh '''
                    go version
                    node --version
                    npm --version
                '''
            }
        }

        stage('2. Pruebas Unitarias — Backend (Go)') {
            steps {
                echo '▶ Ejecutando suite de pruebas del backend (Repositorio, Servicio y Handler)...'
                dir('backend') {
                    sh 'go test ./... -v -timeout 60s'
                }
            }
        }

        stage('3. Compilación — Backend (Go)') {
            steps {
                echo '▶ Compilando binario estático del servidor API en Go...'
                dir('backend') {
                    sh '''
                        mkdir -p bin
                        go build -ldflags="-s -w" -o bin/studynotes-api ./cmd/api
                    '''
                }
            }
        }

        stage('4. Pruebas Unitarias — Frontend (React + Jest)') {
            steps {
                echo '▶ Instalando dependencias de Node.js y ejecutando pruebas con Jest...'
                dir('frontend') {
                    sh '''
                        npm ci
                        npm test -- --passWithNoTests --watchAll=false
                    '''
                }
            }
        }

        stage('5. Construcción de Producción — Frontend (Vite)') {
            steps {
                echo '▶ Verificando tipos con TypeScript y generando artefactos estáticos en dist/...'
                dir('frontend') {
                    sh 'npm run build'
                }
            }
        }
    }

    post {
        success {
            echo '✓ Pipeline completado exitosamente. Archivando artefactos de la versión...'
            archiveArtifacts artifacts: 'backend/bin/studynotes-api, frontend/dist/**/*', fingerprint: true
        }
        failure {
            echo '✗ El pipeline se detuvo debido a un fallo en las pruebas o en la compilación.'
        }
        always {
            echo '▶ Ejecución de CI/CD finalizada.'
        }
    }
}
