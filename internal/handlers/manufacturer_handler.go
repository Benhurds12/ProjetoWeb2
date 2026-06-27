package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"projetoweb2/internal/services"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ManufacturerHandler struct {
	Service *services.ManufacturerService
}

func NewManufacturerHandler(service *services.ManufacturerService) *ManufacturerHandler {
	return &ManufacturerHandler{Service: service}
}

func (h *ManufacturerHandler) CreateFabricante(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Nome string `json:"nome"`
		Cnpj string `json:"cnpj"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	fabricante, err := h.Service.CreateFabricante(
		r.Context(),
		req.Nome,
		req.Cnpj,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fabricante)
}

func (h *ManufacturerHandler) ListFabricantes(w http.ResponseWriter, r *http.Request) {
	fabricantes, err := h.Service.ListFabricantes(r.Context())
	if err != nil {
		log.Println("Erro ao listar fabricantes:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(fabricantes)
}

func (h *ManufacturerHandler) GetFabricante(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	fabricante, err := h.Service.GetFabricante(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(fabricante)
}

func (h *ManufacturerHandler) UpdateFabricante(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	var req struct {
		Nome string `json:"nome"`
		Cnpj string `json:"cnpj"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	fabricante, err := h.Service.UpdateFabricante(
		r.Context(),
		int32(id),
		req.Nome,
		req.Cnpj,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(fabricante)
}

func (h *ManufacturerHandler) DeleteFabricante(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteFabricante(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "erro ao deletar fabricante", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
