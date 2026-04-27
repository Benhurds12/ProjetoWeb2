package services

import (
	"context"
	"database/sql"
	"errors"

	"projetoweb2/internal/db"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Queries *db.Queries
}

func NewUserService(q *db.Queries) *UserService {
	return &UserService{Queries: q}
}

func (s *UserService) CreateUser(ctx context.Context, nome, email, cpf, password string) (db.User, error) {
	if nome == "" || email == "" || cpf == "" || password == "" {
		return db.User{}, errors.New("todos os campos são obrigatórios")
	}

	_, err := s.Queries.GetUserByEmail(ctx, email)
	if err == nil {
		return db.User{}, errors.New("email já cadastrado")
	}
	if err != sql.ErrNoRows {
		return db.User{}, err
	}

	_, err = s.Queries.GetUserByCpf(ctx, cpf)
	if err == nil {
		return db.User{}, errors.New("cpf já cadastrado")
	}
	if err != sql.ErrNoRows {
		return db.User{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return db.User{}, err
	}

	return s.Queries.CreateUser(ctx, db.CreateUserParams{
		Nome:     nome,
		Email:    email,
		Cpf:      cpf,
		Password: string(hash),
	})
}

func (s *UserService) GetUser(ctx context.Context, id int32) (db.User, error) {
	user, err := s.Queries.GetUserByID(ctx, id)
	if err != nil {
		return db.User{}, errors.New("usuário não encontrado")
	}
	return user, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]db.User, error) {
	return s.Queries.ListUsers(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, nome, email, cpf string) (db.User, error) {
	if nome == "" || email == "" || cpf == "" {
		return db.User{}, errors.New("nome, email e cpf são obrigatórios")
	}

	existing, err := s.Queries.GetUserByEmail(ctx, email)
	if err == nil && existing.ID != id {
		return db.User{}, errors.New("email já cadastrado")
	}
	if err != nil && err != sql.ErrNoRows {
		return db.User{}, err
	}

	existing, err = s.Queries.GetUserByCpf(ctx, cpf)
	if err == nil && existing.ID != id {
		return db.User{}, errors.New("cpf já cadastrado")
	}
	if err != nil && err != sql.ErrNoRows {
		return db.User{}, err
	}

	return s.Queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:    id,
		Nome:  nome,
		Email: email,
		Cpf:   cpf,
	})
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.Queries.DeleteUser(ctx, id)
}
