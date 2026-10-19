package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"todoapp/backend/internal/handler"
	"todoapp/backend/internal/repository"
	"todoapp/backend/internal/service"
)

// corsMiddleware es un middleware simple para manejar CORS permitiendo solo http://localhost:5173
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Manejar las solicitudes preflight de CORS
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Obtener el puerto desde el entorno o usar el predeterminado
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Inicializar dependencias
	repo := repository.NewMemoryRepository()
	svc := service.NewNoteService(repo)
	h := handler.NewNoteHandler(svc)

	// Configurar enrutador
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	// Rutas
	r.Get("/healthz", h.HealthCheck)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/notes", h.GetNotes)
		r.Post("/notes", h.CreateNote)
		r.Patch("/notes/{id}/toggle", h.ToggleStatus)
	})

	// Configurar servidor HTTP
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Canal para recibir señales del sistema
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Iniciar servidor en una gorutina
	go func() {
		log.Printf("Iniciando servidor en el puerto %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error al iniciar el servidor: %v", err)
		}
	}()

	// Esperar por la señal de terminación
	<-stop
	log.Println("Apagando el servidor...")

	// Crear un contexto con tiempo de espera para el apagado elegante
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Error al apagar el servidor de manera elegante: %v", err)
	}

	log.Println("Servidor detenido correctamente")
}
