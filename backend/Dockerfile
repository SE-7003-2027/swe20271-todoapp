# Etapa de construcción
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copiar archivos de módulos y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/api cmd/api/main.go

# Etapa de ejecución
FROM alpine:latest

WORKDIR /app

# Instalar certificados raíz para conexiones seguras
RUN apk --no-cache add ca-certificates

# Copiar el binario compilado de la etapa anterior
COPY --from=builder /app/bin/api .

# Exponer el puerto
EXPOSE 8080

# Comando para iniciar la aplicación
CMD ["./api"]
