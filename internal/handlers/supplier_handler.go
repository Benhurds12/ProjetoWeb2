package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"projetoweb2/internal/db"

	"github.com/go-chi/chi/v5"
)

type SupplierHandler struct {
	Queries *db.Queries
}

// CREATE
func (h *SupplierHandler) CreateFornecedor(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Nome    string `json:"nome"`
		Cnpj    string `json:"cnpj"`
		Contato string `json:"contato"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	fornecedor, err := h.Queries.CreateFornecedor(r.Context(), db.CreateFornecedorParams{
		Nome:    req.Nome,
		Cnpj:    req.Cnpj,
		Contato: req.Contato,
	})
	if err != nil {
		http.Error(w, "Erro ao criar fornecedor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fornecedor)
}

// LIST
func (h *SupplierHandler) ListFornecedores(w http.ResponseWriter, r *http.Request) {
	fornecedores, err := h.Queries.ListFornecedores(r.Context())
	if err != nil {
		http.Error(w, "Erro ao listar fornecedores", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(fornecedores)
}

// GET
func (h *SupplierHandler) GetFornecedor(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	fornecedor, err := h.Queries.GetFornecedorByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Fornecedor não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(fornecedor)
}

// UPDATE
func (h *SupplierHandler) UpdateFornecedor(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var req struct {
		Nome    string `json:"nome"`
		Cnpj    string `json:"cnpj"`
		Contato string `json:"contato"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	fornecedor, err := h.Queries.UpdateFornecedor(r.Context(), db.UpdateFornecedorParams{
		ID:      int32(id),
		Nome:    req.Nome,
		Cnpj:    req.Cnpj,
		Contato: req.Contato,
	})
	if err != nil {
		http.Error(w, "Erro ao atualizar fornecedor", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(fornecedor)
}

// DELETE
func (h *SupplierHandler) DeleteFornecedor(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.Queries.DeleteFornecedor(r.Context(), int32(id)); err != nil {
		http.Error(w, "Erro ao deletar fornecedor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
