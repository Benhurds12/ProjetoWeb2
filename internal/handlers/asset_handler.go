package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"projetoweb2/internal/db"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type AssetHandler struct {
	Queries *db.Queries
}

func (h *AssetHandler) CreateBem(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Nome    string `json:"nome"`
		Tipo    string `json:"tipo"`
		Status  string `json:"status"`
		SetorID *int32 `json:"setor_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	newID := uuid.New()

	status := "OCIOSO"
	if req.Status != "" {
		status = req.Status
	}

	var sectorID sql.NullInt32
	if req.SetorID != nil {
		sectorID = sql.NullInt32{Int32: *req.SetorID, Valid: true}
	} else {
		sectorID = sql.NullInt32{Valid: false}
	}

	bem, err := h.Queries.CreateBem(r.Context(), db.CreateBemParams{
		ID:      newID,
		Nome:    req.Nome,
		Status:  sql.NullString{String: status, Valid: true},
		Tipo:    req.Tipo,
		SetorID: sectorID,
	})
	if err != nil {
		http.Error(w, "Erro ao criar bem: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bem)
}

func (h *AssetHandler) GetBem(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	parsedUUID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "UUID inválido", http.StatusBadRequest)
		return
	}

	bem, err := h.Queries.GetBemByID(r.Context(), parsedUUID)
	if err != nil {
		http.Error(w, "Bem não encontrado", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(bem)
}

func (h *AssetHandler) ListBens(w http.ResponseWriter, r *http.Request) {
	bens, err := h.Queries.ListBens(r.Context())
	if err != nil {
		http.Error(w, "Erro ao listar bens", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(bens)
}

func (h *AssetHandler) UpdateBem(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	parsedUUID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "UUID inválido", http.StatusBadRequest)
		return
	}

	var req struct {
		Nome    string `json:"nome"`
		Tipo    string `json:"tipo"`
		Status  string `json:"status"`
		SetorID *int32 `json:"setor_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	var sectorID sql.NullInt32
	if req.SetorID != nil {
		sectorID = sql.NullInt32{Int32: *req.SetorID, Valid: true}
	}

	bem, err := h.Queries.UpdateBem(r.Context(), db.UpdateBemParams{
		ID:      parsedUUID,
		Nome:    req.Nome,
		Status:  sql.NullString{String: req.Status, Valid: true},
		Tipo:    req.Tipo,
		SetorID: sectorID,
	})
	if err != nil {
		http.Error(w, "Erro ao atualizar bem", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(bem)
}

func (h *AssetHandler) DeleteBem(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	parsedUUID, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "UUID inválido", http.StatusBadRequest)
		return
	}

	if err := h.Queries.DeleteBem(r.Context(), parsedUUID); err != nil {
		http.Error(w, "Erro ao deletar bem", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
