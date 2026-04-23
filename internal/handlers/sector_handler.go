package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"projetoweb2/internal/db"

	"github.com/go-chi/chi/v5"
)

type SectorHandler struct {
	Queries *db.Queries
}

func (h *SectorHandler) CreateSetor(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Nome  string `json:"nome"`
		Local string `json:"local"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	setor, err := h.Queries.CreateSetor(r.Context(), db.CreateSetorParams{
		Nome:  req.Nome,
		Local: req.Local,
	})
	if err != nil {
		http.Error(w, "Erro ao criar setor: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(setor)
}

func (h *SectorHandler) GetSetor(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	setor, err := h.Queries.GetSetorByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Setor não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(setor)
}

func (h *SectorHandler) ListSetores(w http.ResponseWriter, r *http.Request) {
	setores, err := h.Queries.ListSetores(r.Context())
	if err != nil {
		http.Error(w, "Erro ao listar setores", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(setores)
}

func (h *SectorHandler) UpdateSetor(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var req struct {
		Nome  string `json:"nome"`
		Local string `json:"local"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	setor, err := h.Queries.UpdateSetor(r.Context(), db.UpdateSetorParams{
		ID:    int32(id),
		Nome:  req.Nome,
		Local: req.Local,
	})
	if err != nil {
		http.Error(w, "Erro ao atualizar setor: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(setor)
}

func (h *SectorHandler) DeleteSetor(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.Queries.DeleteSetor(r.Context(), int32(id)); err != nil {
		http.Error(w, "Erro ao deletar setor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
