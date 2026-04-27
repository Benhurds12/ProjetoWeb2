package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"projetoweb2/internal/services"

	"github.com/go-chi/chi/v5"
)

type SupplierHandler struct {
	Service *services.SupplierService
}

func NewSupplierHandler(service *services.SupplierService) *SupplierHandler {
	return &SupplierHandler{Service: service}
}

func (h *SupplierHandler) CreateFornecedor(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Nome    string `json:"nome"`
		Cnpj    string `json:"cnpj"`
		Contato string `json:"contato"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	fornecedor, err := h.Service.CreateFornecedor(
		r.Context(),
		req.Nome,
		req.Cnpj,
		req.Contato,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fornecedor)
}

func (h *SupplierHandler) ListFornecedores(w http.ResponseWriter, r *http.Request) {
	fornecedores, err := h.Service.ListFornecedores(r.Context())
	if err != nil {
		http.Error(w, "erro ao listar fornecedores", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(fornecedores)
}

func (h *SupplierHandler) GetFornecedor(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	fornecedor, err := h.Service.GetFornecedor(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(fornecedor)
}

func (h *SupplierHandler) UpdateFornecedor(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	var req struct {
		Nome    string `json:"nome"`
		Cnpj    string `json:"cnpj"`
		Contato string `json:"contato"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	fornecedor, err := h.Service.UpdateFornecedor(
		r.Context(),
		int32(id),
		req.Nome,
		req.Cnpj,
		req.Contato,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(fornecedor)
}

func (h *SupplierHandler) DeleteFornecedor(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteFornecedor(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "erro ao deletar fornecedor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
