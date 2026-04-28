package services

import (
	"context"
	"database/sql"
	"testing"

	"projetoweb2/internal/db"
)

type MockSupplierDB struct {
	db.Querier

	stubGetFornecedorByCnpj db.Fornecedore
	errGetFornecedorByCnpj  error

	stubCreateFornecedor db.Fornecedore
	errCreateFornecedor  error

	stubGetFornecedorByID db.Fornecedore
	errGetFornecedorByID  error

	stubUpdateFornecedor db.Fornecedore
	errUpdateFornecedor  error

	stubListFornecedores []db.Fornecedore
	errListFornecedores  error

	errDeleteFornecedor error
}

func (m *MockSupplierDB) GetFornecedorByCnpj(ctx context.Context, cnpj string) (db.Fornecedore, error) {
	return m.stubGetFornecedorByCnpj, m.errGetFornecedorByCnpj
}

func (m *MockSupplierDB) CreateFornecedor(ctx context.Context, arg db.CreateFornecedorParams) (db.Fornecedore, error) {
	return m.stubCreateFornecedor, m.errCreateFornecedor
}

func (m *MockSupplierDB) GetFornecedorByID(ctx context.Context, id int32) (db.Fornecedore, error) {
	return m.stubGetFornecedorByID, m.errGetFornecedorByID
}

func (m *MockSupplierDB) UpdateFornecedor(ctx context.Context, arg db.UpdateFornecedorParams) (db.Fornecedore, error) {
	return m.stubUpdateFornecedor, m.errUpdateFornecedor
}

func (m *MockSupplierDB) ListFornecedores(ctx context.Context) ([]db.Fornecedore, error) {
	return m.stubListFornecedores, m.errListFornecedores
}

func (m *MockSupplierDB) DeleteFornecedor(ctx context.Context, id int32) error {
	return m.errDeleteFornecedor
}

func TestCreateFornecedor_Sucesso(t *testing.T) {
	mock := &MockSupplierDB{
		errGetFornecedorByCnpj: sql.ErrNoRows, // CNPJ livre
		stubCreateFornecedor:   db.Fornecedore{ID: 1, Nome: "Dell LTDA", Cnpj: "12345678000199"},
	}
	service := NewSupplierService(mock)

	f, err := service.CreateFornecedor(context.Background(), "Dell LTDA", "12345678000199", "contato@dell.com")

	if err != nil {
		t.Errorf("Não esperava erro: %v", err)
	}
	if f.Nome != "Dell LTDA" {
		t.Errorf("Esperava Dell LTDA, recebeu %s", f.Nome)
	}
}

func TestCreateFornecedor_ErroCamposVazios(t *testing.T) {
	mock := &MockSupplierDB{}
	service := NewSupplierService(mock)

	_, err := service.CreateFornecedor(context.Background(), "", "", "contato")

	if err == nil || err.Error() != "nome e cnpj são obrigatórios" {
		t.Errorf("Esperava erro de campos obrigatórios, recebeu: %v", err)
	}
}

func TestCreateFornecedor_ErroCnpjDuplicado(t *testing.T) {
	mock := &MockSupplierDB{
		errGetFornecedorByCnpj:  nil, // CNPJ já achado no banco
		stubGetFornecedorByCnpj: db.Fornecedore{ID: 2, Cnpj: "12345678000199"},
	}
	service := NewSupplierService(mock)

	_, err := service.CreateFornecedor(context.Background(), "Outra Empresa", "12345678000199", "")

	if err == nil || err.Error() != "cnpj já cadastrado" {
		t.Errorf("Esperava erro de CNPJ duplicado, recebeu: %v", err)
	}
}

func TestGetFornecedor_Sucesso(t *testing.T) {
	mock := &MockSupplierDB{
		stubGetFornecedorByID: db.Fornecedore{ID: 1, Nome: "Apple Brasil"},
	}
	service := NewSupplierService(mock)

	f, err := service.GetFornecedor(context.Background(), 1)
	if err != nil || f.Nome != "Apple Brasil" {
		t.Errorf("Erro inesperado no GetFornecedor: %v", err)
	}
}

func TestListFornecedores_Sucesso(t *testing.T) {
	mock := &MockSupplierDB{
		stubListFornecedores: []db.Fornecedore{{Nome: "Empresa A"}, {Nome: "Empresa B"}},
	}
	service := NewSupplierService(mock)

	lista, err := service.ListFornecedores(context.Background())
	if err != nil || len(lista) != 2 {
		t.Errorf("Erro inesperado no ListFornecedores: %v", err)
	}
}

func TestUpdateFornecedor_Sucesso(t *testing.T) {
	mock := &MockSupplierDB{
		errGetFornecedorByCnpj: sql.ErrNoRows,
		stubUpdateFornecedor:   db.Fornecedore{ID: 1, Nome: "LG Eletronics"},
	}
	service := NewSupplierService(mock)

	f, err := service.UpdateFornecedor(context.Background(), 1, "LG Eletronics", "99999999000199", "contato@lg.com")

	if err != nil || f.Nome != "LG Eletronics" {
		t.Errorf("Erro inesperado no UpdateFornecedor: %v", err)
	}
}

func TestDeleteFornecedor_Sucesso(t *testing.T) {
	mock := &MockSupplierDB{
		errDeleteFornecedor: nil,
	}
	service := NewSupplierService(mock)

	err := service.DeleteFornecedor(context.Background(), 1)
	if err != nil {
		t.Errorf("Erro inesperado no DeleteFornecedor: %v", err)
	}
}
