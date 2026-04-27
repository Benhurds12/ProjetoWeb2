package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"projetoweb2/internal/services"

	"github.com/go-chi/chi/v5"
)

type SectorHandler struct {
	Service *services.SectorService
}

func NewSectorHandler(service *services.SectorService) *SectorHandler {
	return &SectorHandler{Service: service}
}

func (h *SectorHandler) CreateSetor(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Nome  string `json:"nome"`
		Local string `json:"local"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	setor, err := h.Service.CreateSetor(r.Context(), req.Nome, req.Local)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(setor)
}

func (h *SectorHandler) GetSetor(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	setor, err := h.Service.GetSetor(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(setor)
}

func (h *SectorHandler) ListSetores(w http.ResponseWriter, r *http.Request) {
	setores, err := h.Service.ListSetores(r.Context())
	if err != nil {
		http.Error(w, "erro ao listar setores", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(setores)
}

func (h *SectorHandler) UpdateSetor(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	var req struct {
		Nome  string `json:"nome"`
		Local string `json:"local"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	setor, err := h.Service.UpdateSetor(r.Context(), int32(id), req.Nome, req.Local)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(setor)
}

func (h *SectorHandler) DeleteSetor(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteSetor(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "erro ao deletar setor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
