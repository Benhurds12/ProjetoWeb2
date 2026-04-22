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
	})

	log.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", r)
}
