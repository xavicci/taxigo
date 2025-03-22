package main

import (
	"log"
	"net/http"
	"taxiya/internal/auth"
	"taxiya/internal/server"
)

func main() {
	// Inicializar el servicio de autenticación
	authService := auth.NewAuthService()

	// Inicializar el handler HTTP
	httpHandler := server.NewHTTPHandler(authService)

	// Configurar rutas HTTP
	http.HandleFunc("/api/login", httpHandler.Login)
	http.HandleFunc("/api/register", httpHandler.Register)

	// Iniciar servidor HTTP
	go func() {
		log.Printf("HTTP Server listening at :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Iniciar servidor gRPC
	if err := server.StartServer(":50051"); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
