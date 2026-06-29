package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"projetoweb2/internal/config"
	"projetoweb2/internal/db"
	"projetoweb2/internal/handlers"
	"projetoweb2/internal/middleware"
	"projetoweb2/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	config.Load()

	conn, err := sql.Open("pgx", config.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)

	userService := services.NewUserService(queries)

	userHandler := handlers.NewUserHandler(userService)

	authHandler := &handlers.AuthHandler{
		Queries: queries,
	}
	sectorService := services.NewSectorService(queries)
	sectorHandler := handlers.NewSectorHandler(sectorService)
	assetService := services.NewAssetService(queries)
	assetHandler := handlers.NewAssetHandler(assetService)
	supplierService := services.NewSupplierService(queries)
	supplierHandler := handlers.NewSupplierHandler(supplierService)
	manufacturerService := services.NewManufacturerService(queries)
	manufacturerHandler := handlers.NewManufacturerHandler(manufacturerService)
	healthHandler := handlers.NewHealthHandler(conn)

	r := chi.NewRouter()

	// Cabeçalhos de segurança em todas as respostas (OWASP A05).
	r.Use(middleware.SecurityHeaders)

	// Health checks para monitoramento e orquestradores.
	r.Get("/healthz", healthHandler.Healthz)
	r.Get("/readyz", healthHandler.Readyz)

	r.Post("/users", userHandler.CreateUser)

	// Rate limit nas rotas sensíveis de autenticação para mitigar
	// brute-force / credential stuffing (OWASP A07 – Identification and Authentication Failures).
	r.Group(func(r chi.Router) {
		r.Use(httprate.LimitByIP(10, time.Minute))
		r.Post("/login", authHandler.Login)
		r.Post("/refresh", authHandler.Refresh)
	})

	r.Post("/logout", authHandler.Logout)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Get("/users", userHandler.ListUsers)
		r.Get("/users/{id}", userHandler.GetUser)
		r.Put("/users/{id}", userHandler.UpdateUser)
		r.Delete("/users/{id}", userHandler.DeleteUser)

		r.Post("/setores", sectorHandler.CreateSetor)
		r.Get("/setores", sectorHandler.ListSetores)
		r.Get("/setores/{id}", sectorHandler.GetSetor)
		r.Put("/setores/{id}", sectorHandler.UpdateSetor)
		r.Delete("/setores/{id}", sectorHandler.DeleteSetor)
		r.Get("/setores-bens", sectorHandler.ListSetoresWithBens)

		r.Post("/bens", assetHandler.CreateBem)
		r.Get("/bens", assetHandler.ListBens)
		r.Get("/bens/{id}", assetHandler.GetBem)
		r.Put("/bens/{id}", assetHandler.UpdateBem)
		r.Delete("/bens/{id}", assetHandler.DeleteBem)

		r.Post("/fornecedores", supplierHandler.CreateFornecedor)
		r.Get("/fornecedores", supplierHandler.ListFornecedores)
		r.Get("/fornecedores/{id}", supplierHandler.GetFornecedor)
		r.Put("/fornecedores/{id}", supplierHandler.UpdateFornecedor)
		r.Delete("/fornecedores/{id}", supplierHandler.DeleteFornecedor)

		r.Post("/fabricantes", manufacturerHandler.CreateFabricante)
		r.Get("/fabricantes", manufacturerHandler.ListFabricantes)
		r.Get("/fabricantes/{id}", manufacturerHandler.GetFabricante)
		r.Put("/fabricantes/{id}", manufacturerHandler.UpdateFabricante)
		r.Delete("/fabricantes/{id}", manufacturerHandler.DeleteFabricante)
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Sobe o servidor em uma goroutine para não bloquear o tratamento de sinais.
	go func() {
		log.Println("Servidor rodando em :8080")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("erro ao iniciar servidor: %v", err)
		}
	}()

	// Aguarda um sinal de interrupção (Ctrl+C) ou término (SIGTERM do Docker).
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Encerrando servidor...")

	// Dá até 15s para as requisições em andamento terminarem antes de fechar.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("erro no shutdown do servidor: %v", err)
	}

	if err := conn.Close(); err != nil {
		log.Printf("erro ao fechar conexão com o banco: %v", err)
	}

	log.Println("Servidor encerrado com sucesso")
}
