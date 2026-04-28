package services

import (
	"context"
	"database/sql"
	"testing"

	"projetoweb2/internal/db"
)

type MockUserDB struct {
	db.Querier

	stubGetUserByEmail db.User
	errGetUserByEmail  error

	stubGetUserByCpf db.User
	errGetUserByCpf  error

	stubCreateUser db.User
	errCreateUser  error

	stubGetUserByID db.User
	errGetUserByID  error

	stubUpdateUser db.User
	errUpdateUser  error

	stubListUsers []db.User
	errListUsers  error

	errDeleteUser error
}

func (m *MockUserDB) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	return m.stubGetUserByEmail, m.errGetUserByEmail
}

func (m *MockUserDB) GetUserByCpf(ctx context.Context, cpf string) (db.User, error) {
	return m.stubGetUserByCpf, m.errGetUserByCpf
}

func (m *MockUserDB) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return m.stubCreateUser, m.errCreateUser
}

func (m *MockUserDB) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return m.stubGetUserByID, m.errGetUserByID
}

func (m *MockUserDB) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	return m.stubUpdateUser, m.errUpdateUser
}

func (m *MockUserDB) ListUsers(ctx context.Context) ([]db.User, error) {
	return m.stubListUsers, m.errListUsers
}

func (m *MockUserDB) DeleteUser(ctx context.Context, id int32) error {
	return m.errDeleteUser
}

func TestCreateUser_Sucesso(t *testing.T) {
	mock := &MockUserDB{
		errGetUserByEmail: sql.ErrNoRows, // Email não existe (pode criar)
		errGetUserByCpf:   sql.ErrNoRows, // CPF não existe (pode criar)
		stubCreateUser:    db.User{ID: 1, Nome: "Ben Hur", Email: "ben@teste.com"},
	}
	service := NewUserService(mock)

	user, err := service.CreateUser(context.Background(), "Ben Hur", "ben@teste.com", "12345678900", "senhaSegura123")

	if err != nil {
		t.Errorf("Não esperava erro, recebeu: %v", err)
	}
	if user.Nome != "Ben Hur" {
		t.Errorf("Esperava Ben Hur, recebeu %s", user.Nome)
	}
}

func TestCreateUser_ErroCamposVazios(t *testing.T) {
	mock := &MockUserDB{}
	service := NewUserService(mock)

	_, err := service.CreateUser(context.Background(), "", "teste@teste.com", "123", "123")

	if err == nil || err.Error() != "todos os campos são obrigatórios" {
		t.Errorf("Esperava erro de campos obrigatórios, recebeu: %v", err)
	}
}

func TestCreateUser_ErroEmailDuplicado(t *testing.T) {
	mock := &MockUserDB{
		errGetUserByEmail:  nil, // nil significa que o banco ACHOU o email
		stubGetUserByEmail: db.User{ID: 2, Email: "ben@teste.com"},
	}
	service := NewUserService(mock)

	_, err := service.CreateUser(context.Background(), "Outro Nome", "ben@teste.com", "11111111111", "senha123")

	if err == nil || err.Error() != "email já cadastrado" {
		t.Errorf("Esperava erro de email já cadastrado, recebeu: %v", err)
	}
}

func TestCreateUser_ErroCpfDuplicado(t *testing.T) {
	mock := &MockUserDB{
		errGetUserByEmail: sql.ErrNoRows, // Email liberado
		errGetUserByCpf:   nil,           // CPF já existe
		stubGetUserByCpf:  db.User{ID: 3, Cpf: "12345678900"},
	}
	service := NewUserService(mock)

	_, err := service.CreateUser(context.Background(), "Outro", "novo@teste.com", "12345678900", "senha")

	if err == nil || err.Error() != "cpf já cadastrado" {
		t.Errorf("Esperava erro de CPF já cadastrado, recebeu: %v", err)
	}
}

func TestGetUser_Sucesso(t *testing.T) {
	mock := &MockUserDB{
		stubGetUserByID: db.User{ID: 1, Nome: "João"},
	}
	service := NewUserService(mock)

	user, err := service.GetUser(context.Background(), 1)
	if err != nil || user.Nome != "João" {
		t.Errorf("Erro inesperado no GetUser: %v", err)
	}
}

func TestListUsers_Sucesso(t *testing.T) {
	mock := &MockUserDB{
		stubListUsers: []db.User{{Nome: "User 1"}, {Nome: "User 2"}},
	}
	service := NewUserService(mock)

	users, err := service.ListUsers(context.Background())
	if err != nil || len(users) != 2 {
		t.Errorf("Erro inesperado no ListUsers: %v", err)
	}
}

func TestUpdateUser_Sucesso(t *testing.T) {
	mock := &MockUserDB{
		errGetUserByEmail: sql.ErrNoRows,
		errGetUserByCpf:   sql.ErrNoRows,
		stubUpdateUser:    db.User{ID: 1, Nome: "João Atualizado"},
	}
	service := NewUserService(mock)

	user, err := service.UpdateUser(context.Background(), 1, "João Atualizado", "joao@teste.com", "00000000000")

	if err != nil || user.Nome != "João Atualizado" {
		t.Errorf("Erro inesperado no UpdateUser: %v", err)
	}
}

func TestDeleteUser_Sucesso(t *testing.T) {
	mock := &MockUserDB{
		errDeleteUser: nil,
	}
	service := NewUserService(mock)

	err := service.DeleteUser(context.Background(), 1)
	if err != nil {
		t.Errorf("Erro inesperado no DeleteUser: %v", err)
	}
}
