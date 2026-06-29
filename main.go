package main

import (
	"database/sql"
	"log"
	"net/http"

	"projetoweb2/internal/db"
	"projetoweb2/internal/handlers"
	"projetoweb2/internal/middleware"
	"projetoweb2/internal/services"

	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/projetoweb2?sslmode=disable")
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

	r := chi.NewRouter()

	r.Post("/users", userHandler.CreateUser)
	r.Post("/login", authHandler.Login)
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

	log.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", r)
}
