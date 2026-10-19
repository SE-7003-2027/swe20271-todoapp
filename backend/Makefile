.PHONY: run test build lint

# Ejecutar la aplicación en modo desarrollo
run:
	go run cmd/api/main.go

# Ejecutar las pruebas unitarias
test:
	go test -v ./...

# Compilar la aplicación para producción
build:
	go build -o bin/api cmd/api/main.go

# Ejecutar linter
lint:
	go vet ./...
