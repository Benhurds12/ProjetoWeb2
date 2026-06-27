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
	"github.com/go-chi/cors"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	conn, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/projetoweb2?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	var dbName string
	err = conn.QueryRow("SELECT current_database()").Scan(&dbName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Banco conectado:", dbName)

	rows, err := conn.Query(`
	SELECT table_name
	FROM information_schema.tables
	WHERE table_schema = 'public'
	ORDER BY table_name;
`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	log.Println("Tabelas encontradas:")
	for rows.Next() {
		var table string
		rows.Scan(&table)
		log.Println("-", table)
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

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
		},

		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},

		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
		},

		AllowCredentials: true,
	}))

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
