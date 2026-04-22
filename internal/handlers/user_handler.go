package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"projetoweb2/internal/db"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Queries *db.Queries
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nome     string `json:"nome"`
		Email    string `json:"email"`
		Cpf      string `json:"cpf"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "erro ao gerar senha", http.StatusInternalServerError)
		return
	}

	user, err := h.Queries.CreateUser(r.Context(), db.CreateUserParams{
		Nome:     input.Nome,
		Email:    input.Email,
		Cpf:      input.Cpf,
		Password: string(hash),
	})

	if err != nil {
		http.Error(w, "erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Queries.ListUsers(r.Context())
	if err != nil {
		http.Error(w, "erro ao listar usuários", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	user, err := h.Queries.GetUserByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "usuário não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	err = h.Queries.DeleteUser(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "erro ao deletar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	var input struct {
		Nome  string `json:"nome"`
		Email string `json:"email"`
		Cpf   string `json:"cpf"`
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	user, err := h.Queries.UpdateUser(r.Context(), db.UpdateUserParams{
		ID:    int32(id),
		Nome:  input.Nome,
		Email: input.Email,
		Cpf:   input.Cpf,
	})
	if err != nil {
		http.Error(w, "erro ao atualizar usuário", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
