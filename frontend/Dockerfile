# Etapa de construcción
FROM node:20-alpine AS build

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos los archivos de dependencias
COPY package.json package-lock.json* ./

# Instalamos las dependencias
RUN npm install

# Copiamos el resto del código
COPY . .

# Construimos la aplicación con entorno MOCK por defecto
ENV VITE_USE_MOCK=true
RUN npm run build

# Etapa de producción
FROM nginx:alpine

# Copiamos los archivos estáticos construidos al directorio de nginx
COPY --from=build /app/dist /usr/share/nginx/html

# Exponemos el puerto 80
EXPOSE 80

# Iniciamos nginx
CMD ["nginx", "-g", "daemon off;"]
