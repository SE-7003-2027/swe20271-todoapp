// =============================================================================
// Jenkinsfile — Pipeline Declarativo de CI/CD para StudyNotes / TaskSync
// Propósito educativo: Demostrar Integración Continua (CI), pruebas unitarias,
// compilación y Despliegue Continuo (CD) hacia un entorno Staging Full-Stack
// en http://localhost:8082 (conectado al backend real en Go, sin usar el mock).
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
        // Desactivamos el mock forzado para que la app desplegada use el backend real en Go
        VITE_USE_MOCK = 'false'
        // Apuntamos al mismo origen del entorno Staging (:8082), donde Nginx enruta /api y /healthz al contenedor Go
        VITE_API_URL  = 'http://localhost:8082'
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
                echo '▶ Verificando tipos con TypeScript y generando artefactos en dist/...'
                dir('frontend') {
                    sh 'npm run build'
                }
            }
        }

        stage('6. Despliegue Continuo (CD) — Entorno Staging') {
            steps {
                echo '▶ Publicando binario Go y bundle de React en el entorno Staging (http://localhost:8082)...'
                sh '''
                    mkdir -p /cd-deploy/api /cd-deploy/web
                    cp backend/bin/studynotes-api /cd-deploy/api/studynotes-api.tmp
                    chmod +x /cd-deploy/api/studynotes-api.tmp
                    mv /cd-deploy/api/studynotes-api.tmp /cd-deploy/api/studynotes-api

                    rm -rf /cd-deploy/web/*
                    cp -R frontend/dist/* /cd-deploy/web/
                '''
                echo '▶ Ejecutando Smoke Test contra el entorno Staging recién desplegado...'
                sh '''
                    for i in 1 2 3 4 5; do
                        if curl -fsS http://staging-web/healthz | grep -q '"status":"ok"'; then
                            echo "✓ Smoke Test exitoso: La API en Go y el servidor web en Staging responden correctamente."
                            exit 0
                        fi
                        echo "Esperando a que el binario de Go en Staging inicie (intento $i/5)..."
                        sleep 2
                    done
                    echo "✗ Error: El Smoke Test hacia http://staging-web/healthz no respondió a tiempo."
                    exit 1
                '''
            }
        }
    }

    post {
        success {
            echo '✓ Pipeline CI/CD completado. Aplicación desplegada en vivo en: http://localhost:8082'
            archiveArtifacts artifacts: 'backend/bin/studynotes-api, frontend/dist/**/*', fingerprint: true
        }
        failure {
            echo '✗ El pipeline se detuvo debido a un fallo en las pruebas, compilación o despliegue.'
        }
        always {
            echo '▶ Ejecución de CI/CD finalizada.'
        }
    }
}
