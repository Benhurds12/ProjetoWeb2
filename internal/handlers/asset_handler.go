package handlers

import (
	"encoding/json"
	"net/http"

	"projetoweb2/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type AssetHandler struct {
	Service *services.AssetService
}

func NewAssetHandler(service *services.AssetService) *AssetHandler {
	return &AssetHandler{Service: service}
}

func (h *AssetHandler) CreateBem(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Nome    string `json:"nome"`
		Tipo    string `json:"tipo"`
		Status  string `json:"status"`
		SetorID *int32 `json:"setor_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	bem, err := h.Service.CreateBem(r.Context(), req.Nome, req.Tipo, req.Status, req.SetorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bem)
}

func (h *AssetHandler) GetBem(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "UUID inválido", http.StatusBadRequest)
		return
	}

	bem, err := h.Service.GetBem(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(bem)
}

func (h *AssetHandler) ListBens(w http.ResponseWriter, r *http.Request) {
	bens, err := h.Service.ListBens(r.Context())
	if err != nil {
		http.Error(w, "erro ao listar bens", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bens)
}

func (h *AssetHandler) UpdateBem(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
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
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	bem, err := h.Service.UpdateBem(r.Context(), id, req.Nome, req.Tipo, req.Status, req.SetorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(bem)
}

func (h *AssetHandler) DeleteBem(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "UUID inválido", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteBem(r.Context(), id)
	if err != nil {
		http.Error(w, "erro ao deletar bem", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
