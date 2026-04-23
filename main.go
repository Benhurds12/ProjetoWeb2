package main

import (
	"database/sql"
	"log"
	"net/http"

	"projetoweb2/internal/db"
	"projetoweb2/internal/handlers"
	"projetoweb2/internal/middleware"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/projetoweb2?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)

	userHandler := &handlers.UserHandler{
		Queries: queries,
	}

	authHandler := &handlers.AuthHandler{
		Queries: queries,
	}
	sectorHandler := &handlers.SectorHandler{Queries: queries}
	assetHandler := &handlers.AssetHandler{Queries: queries}

	r := chi.NewRouter()

	// Rotas públicas
	r.Post("/users", userHandler.CreateUser)
	r.Post("/login", authHandler.Login)
	r.Post("/logout", authHandler.Logout)

	// Rotas protegidas
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

		r.Post("/bens", assetHandler.CreateBem)
		r.Get("/bens", assetHandler.ListBens)
		r.Get("/bens/{id}", assetHandler.GetBem)
		r.Put("/bens/{id}", assetHandler.UpdateBem)
		r.Delete("/bens/{id}", assetHandler.DeleteBem)

	})

	log.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", r)
}
