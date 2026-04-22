package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"projetoweb2/internal/db"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JwtSecret = []byte("sua-chave-secreta") // troque por uma chave forte

type AuthHandler struct {
	Queries *db.Queries
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	// Busca usuário pelo email
	user, err := h.Queries.GetUserByEmail(r.Context(), input.Email)
	if err != nil {
		http.Error(w, "usuário não encontrado", http.StatusUnauthorized)
		return
	}

	// Compara a senha
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		http.Error(w, "senha incorreta", http.StatusUnauthorized)
		return
	}

	// Gera o token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		http.Error(w, "erro ao gerar token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenString,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// JWT é stateless — o logout é feito no cliente apagando o token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "logout realizado com sucesso",
	})
}
