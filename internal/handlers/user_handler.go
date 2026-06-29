package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"projetoweb2/internal/db"
	"projetoweb2/internal/services"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// UserOut é a representação pública de um usuário.
// Omitimos o hash da senha para nunca expô-lo nas respostas da API (OWASP A02 – Cryptographic Failures).
type UserOut struct {
	ID    int32  `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
}

func toUserOut(u db.User) UserOut {
	return UserOut{
		ID:    u.ID,
		Nome:  u.Nome,
		Email: u.Email,
		Cpf:   u.Cpf,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nome     string `json:"nome"`
		Email    string `json:"email"`
		Cpf      string `json:"cpf"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	user, err := h.Service.CreateUser(r.Context(), input.Nome, input.Email, input.Cpf, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(toUserOut(user))
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.ListUsers(r.Context())
	if err != nil {
		http.Error(w, "erro ao listar usuários", http.StatusInternalServerError)
		return
	}

	out := make([]UserOut, 0, len(users))
	for _, u := range users {
		out = append(out, toUserOut(u))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(out)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetUser(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toUserOut(user))
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

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	user, err := h.Service.UpdateUser(r.Context(), int32(id), input.Nome, input.Email, input.Cpf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toUserOut(user))
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteUser(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "erro ao deletar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
