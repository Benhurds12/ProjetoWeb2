package middleware

import (
	"context"
	"net/http"
	"strings"

	"projetoweb2/internal/handlers"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pega o header Authorization: Bearer <token>
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "token não fornecido", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "formato de token inválido", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// Valida o token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return handlers.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "token inválido", http.StatusUnauthorized)
			return
		}

		// Injeta o user_id no contexto
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"]
		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
