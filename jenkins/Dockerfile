# ==============================================================================
# Dockerfile de Jenkins para Laboratorio de CI/CD — StudyNotes / TaskSync
# Incluye Jenkins LTS + Go + Node.js + npm preinstalados para ejecutar el
# Jenkinsfile directamente en el nodo principal durante las sesiones de clase.
# ==============================================================================
FROM jenkins/jenkins:lts-jdk17

USER root

# Instalar dependencias del sistema, Go y Node.js 20.x
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    git \
    gnupg \
    make \
    golang-go \
    && curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y --no-install-recommends nodejs \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Crear directorio compartido de despliegue continuo (CD) con permisos para jenkins
RUN mkdir -p /cd-deploy/api /cd-deploy/web && chown -R jenkins:jenkins /cd-deploy

# Verificar versiones instaladas durante la construcción de la imagen
RUN go version && node --version && npm --version

USER jenkins

# Preinstalar plugins esenciales para Declarative Pipelines y visualización de etapas
RUN jenkins-plugin-cli --plugins \
    workflow-aggregator \
    git \
    pipeline-stage-view \
    timestamper \
    ws-cleanup
