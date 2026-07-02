package handlers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"projetoweb2/internal/config"
	"projetoweb2/internal/db"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	// Access token de vida curta; o refresh token cobre a sessão longa.
	accessTokenTTL  = 15 * time.Minute
	refreshTokenTTL = 7 * 24 * time.Hour
)

type AuthHandler struct {
	Queries *db.Queries
}

// generateRefreshToken gera um token opaco aleatório e seu hash.
// O valor em texto puro vai para o cliente; só o hash é persistido.
func generateRefreshToken() (plain string, hash string, err error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", "", err
	}
	plain = base64.RawURLEncoding.EncodeToString(b)
	sum := sha256.Sum256([]byte(plain))
	hash = hex.EncodeToString(sum[:])
	return plain, hash, nil
}

func hashRefreshToken(plain string) string {
	sum := sha256.Sum256([]byte(plain))
	return hex.EncodeToString(sum[:])
}

// generateAccessToken cria um JWT de acesso assinado para o usuário.
func generateAccessToken(userID int32, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(accessTokenTTL).Unix(),
	})
	return token.SignedString(config.JwtSecret)
}

// issueRefreshToken cria e persiste um novo refresh token para o usuário,
// retornando o valor em texto puro a ser enviado ao cliente.
func (h *AuthHandler) issueRefreshToken(r *http.Request, userID int32) (string, error) {
	plain, hash, err := generateRefreshToken()
	if err != nil {
		return "", err
	}

	_, err = h.Queries.CreateRefreshToken(r.Context(), db.CreateRefreshTokenParams{
		ID:        uuid.New(),
		UserID:    userID,
		TokenHash: hash,
		ExpiresAt: time.Now().Add(refreshTokenTTL),
	})
	if err != nil {
		return "", err
	}
	return plain, nil
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

	user, err := h.Queries.GetUserByEmail(r.Context(), input.Email)
	if err != nil {
		// Mensagem genérica para não revelar se o e-mail existe (enumeração de usuários).
		http.Error(w, "credenciais inválidas", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		http.Error(w, "credenciais inválidas", http.StatusUnauthorized)
		return
	}

	accessToken, err := generateAccessToken(user.ID, user.Email)
	if err != nil {
		http.Error(w, "erro ao gerar token", http.StatusInternalServerError)
		return
	}

	// refreshToken, err := h.issueRefreshToken(r, user.ID)
	// if err != nil {
	// 	http.Error(w, "erro ao gerar refresh token", http.StatusInternalServerError)
	// 	return
	// }

	refreshToken, err := h.issueRefreshToken(r, user.ID)
	if err != nil {
		log.Printf("Erro ao criar refresh token: %v", err)
		http.Error(w, "erro ao gerar refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// Refresh valida um refresh token, aplica rotação (revoga o antigo e emite um
// novo) e devolve um novo par access/refresh.
func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.RefreshToken == "" {
		http.Error(w, "refresh token não fornecido", http.StatusBadRequest)
		return
	}

	stored, err := h.Queries.GetRefreshTokenByHash(r.Context(), hashRefreshToken(input.RefreshToken))
	if err != nil {
		http.Error(w, "refresh token inválido", http.StatusUnauthorized)
		return
	}

	// Reuso de um token já revogado é sinal de comprometimento:
	// revogamos todos os tokens do usuário por segurança.
	if stored.Revoked {
		_ = h.Queries.RevokeAllUserRefreshTokens(r.Context(), stored.UserID)
		http.Error(w, "refresh token inválido", http.StatusUnauthorized)
		return
	}

	if time.Now().After(stored.ExpiresAt) {
		http.Error(w, "refresh token expirado", http.StatusUnauthorized)
		return
	}

	user, err := h.Queries.GetUserByID(r.Context(), stored.UserID)
	if err != nil {
		http.Error(w, "usuário não encontrado", http.StatusUnauthorized)
		return
	}

	// Rotação: invalida o token usado e emite um novo.
	if err := h.Queries.RevokeRefreshToken(r.Context(), stored.ID); err != nil {
		http.Error(w, "erro ao rotacionar token", http.StatusInternalServerError)
		return
	}

	accessToken, err := generateAccessToken(user.ID, user.Email)
	if err != nil {
		http.Error(w, "erro ao gerar token", http.StatusInternalServerError)
		return
	}

	newRefresh, err := h.issueRefreshToken(r, user.ID)
	if err != nil {
		http.Error(w, "erro ao gerar refresh token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": newRefresh,
	})
}

// Logout revoga o refresh token informado, encerrando a sessão de forma efetiva.
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err == nil && input.RefreshToken != "" {
		stored, err := h.Queries.GetRefreshTokenByHash(r.Context(), hashRefreshToken(input.RefreshToken))
		if err == nil {
			_ = h.Queries.RevokeRefreshToken(r.Context(), stored.ID)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "logout realizado com sucesso",
	})
}
