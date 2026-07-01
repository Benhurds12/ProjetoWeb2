package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// HealthHandler expõe endpoints de liveness e readiness para monitoramento
// e orquestradores (Docker/Kubernetes).
type HealthHandler struct {
	DB *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{DB: db}
}

// Healthz é o liveness probe: indica que o processo está de pé.
// Não checa dependências externas.
func (h *HealthHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Readyz é o readiness probe: indica que a aplicação está apta a receber
// tráfego, verificando a conexão com o banco de dados.
func (h *HealthHandler) Readyz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := h.DB.PingContext(r.Context()); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "unavailable",
			"db":     "down",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ready",
		"db":     "up",
	})
}
