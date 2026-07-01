package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// JwtSecret é a chave usada para assinar e validar os tokens JWT.
// É carregada do ambiente (variável JWT_SECRET) e nunca deve ficar
// hardcoded no código-fonte.
var JwtSecret []byte

// Load carrega as variáveis de ambiente a partir do arquivo .env (se existir)
// e inicializa as configurações da aplicação. Deve ser chamada uma vez no
// início do programa (main).
func Load() {
	// Em produção as variáveis já vêm do ambiente; o .env é uma conveniência
	// para o desenvolvimento local, então a ausência dele não é um erro.
	if err := godotenv.Load(); err != nil {
		log.Println("aviso: arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET não definida — configure-a no .env ou no ambiente")
	}
	JwtSecret = []byte(secret)
}

// DatabaseURL retorna a string de conexão com o banco de dados.
// Usa DATABASE_URL do ambiente e, se ausente, cai num default de desenvolvimento.
func DatabaseURL() string {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		return url
	}
	return "postgres://postgres:postgres@localhost:5432/projetoweb2?sslmode=disable"
}
