package services

import (
	"context"
	"database/sql"
	"testing"

	"projetoweb2/internal/db"

	"github.com/google/uuid"
)

// 1. Criamos um Mock focado apenas no que o AssetService precisa usar
type MockAssetDB struct {
	db.Querier

	// Precisamos simular a busca de Setor, pois o AssetService valida isso!
	stubGetSetorByID db.Setore
	errGetSetorByID  error

	stubCreateBem db.Ben
	errCreateBem  error

	stubGetBemByID db.Ben
	errGetBemByID  error

	stubUpdateBem db.Ben
	errUpdateBem  error

	errDeleteBem error

	stubListBens []db.Ben
	errListBens  error
}

func (m *MockAssetDB) GetSetorByID(ctx context.Context, id int32) (db.Setore, error) {
	return m.stubGetSetorByID, m.errGetSetorByID
}

func (m *MockAssetDB) CreateBem(ctx context.Context, arg db.CreateBemParams) (db.Ben, error) {
	return m.stubCreateBem, m.errCreateBem
}

func (m *MockAssetDB) GetBemByID(ctx context.Context, id uuid.UUID) (db.Ben, error) {
	return m.stubGetBemByID, m.errGetBemByID
}

func (m *MockAssetDB) UpdateBem(ctx context.Context, arg db.UpdateBemParams) (db.Ben, error) {
	return m.stubUpdateBem, m.errUpdateBem
}

func (m *MockAssetDB) DeleteBem(ctx context.Context, id uuid.UUID) error {
	return m.errDeleteBem
}

func (m *MockAssetDB) ListBens(ctx context.Context) ([]db.Ben, error) {
	return m.stubListBens, m.errListBens
}

func TestCreateBem_Sucesso(t *testing.T) {
	setorID := int32(1)
	mock := &MockAssetDB{
		// Dizemos que o setor ID 1 existe (sem erro)
		errGetSetorByID:  nil,
		stubGetSetorByID: db.Setore{ID: 1, Nome: "TI"},

		// Retorno esperado do Bem criado
		stubCreateBem: db.Ben{
			Nome:   "Notebook Dell",
			Tipo:   "Eletrônico",
			Status: sql.NullString{String: "OCIOSO", Valid: true},
		},
	}
	service := NewAssetService(mock)

	bem, err := service.CreateBem(context.Background(), "Notebook Dell", "Eletrônico", "", &setorID)

	if err != nil {
		t.Errorf("Não esperava erro: %v", err)
	}
	if bem.Nome != "Notebook Dell" {
		t.Errorf("Esperava Notebook Dell, recebeu %s", bem.Nome)
	}
}

func TestCreateBem_ErroCamposVazios(t *testing.T) {
	mock := &MockAssetDB{} // Não chega a chamar o banco
	service := NewAssetService(mock)

	_, err := service.CreateBem(context.Background(), "", "", "", nil)

	if err == nil || err.Error() != "nome e tipo são obrigatórios" {
		t.Errorf("Esperava erro de campos obrigatórios, recebeu: %v", err)
	}
}

func TestCreateBem_ErroSetorNaoEncontrado(t *testing.T) {
	setorIDInvalido := int32(99)
	mock := &MockAssetDB{
		// Simulamos que o banco NÃO ACHOU o setor
		errGetSetorByID: sql.ErrNoRows,
	}
	service := NewAssetService(mock)

	_, err := service.CreateBem(context.Background(), "Mesa", "Móvel", "ATIVO", &setorIDInvalido)

	if err == nil || err.Error() != "setor não encontrado" {
		t.Errorf("Esperava erro de setor não encontrado, recebeu: %v", err)
	}
}

func TestGetBem_Sucesso(t *testing.T) {
	idMock := uuid.New()
	mock := &MockAssetDB{
		stubGetBemByID: db.Ben{ID: idMock, Nome: "Cadeira Gamer"},
	}
	service := NewAssetService(mock)

	bem, err := service.GetBem(context.Background(), idMock)

	if err != nil {
		t.Errorf("Não esperava erro: %v", err)
	}
	if bem.Nome != "Cadeira Gamer" {
		t.Errorf("Esperava Cadeira Gamer, recebeu %s", bem.Nome)
	}
}

func TestGetBem_NaoEncontrado(t *testing.T) {
	mock := &MockAssetDB{
		errGetBemByID: sql.ErrNoRows,
	}
	service := NewAssetService(mock)

	_, err := service.GetBem(context.Background(), uuid.New())

	if err == nil || err.Error() != "bem não encontrado" {
		t.Errorf("Esperava erro 'bem não encontrado', recebeu: %v", err)
	}
}

func TestListBens_Sucesso(t *testing.T) {
	mock := &MockAssetDB{
		stubListBens: []db.Ben{
			{Nome: "Item 1"},
			{Nome: "Item 2"},
		},
	}
	service := NewAssetService(mock)

	bens, err := service.ListBens(context.Background())

	if err != nil {
		t.Errorf("Não esperava erro: %v", err)
	}
	if len(bens) != 2 {
		t.Errorf("Esperava 2 itens, recebeu %d", len(bens))
	}
}

func TestUpdateBem_Sucesso(t *testing.T) {
	idMock := uuid.New()
	setorID := int32(1)
	mock := &MockAssetDB{
		errGetSetorByID: nil, // Setor existe!
		stubUpdateBem:   db.Ben{ID: idMock, Nome: "Monitor LG"},
	}
	service := NewAssetService(mock)

	bem, err := service.UpdateBem(context.Background(), idMock, "Monitor LG", "Eletrônico", "ATIVO", &setorID)

	if err != nil {
		t.Errorf("Não esperava erro: %v", err)
	}
	if bem.Nome != "Monitor LG" {
		t.Errorf("Update falhou, nome recebido: %s", bem.Nome)
	}
}

func TestUpdateBem_ErroCamposVazios(t *testing.T) {
	mock := &MockAssetDB{}
	service := NewAssetService(mock)

	_, err := service.UpdateBem(context.Background(), uuid.New(), "", "", "ATIVO", nil)

	if err == nil || err.Error() != "nome e tipo são obrigatórios" {
		t.Errorf("Esperava erro de campos obrigatórios, recebeu: %v", err)
	}
}

func TestDeleteBem_Sucesso(t *testing.T) {
	mock := &MockAssetDB{
		errDeleteBem: nil,
	}
	service := NewAssetService(mock)

	err := service.DeleteBem(context.Background(), uuid.New())

	if err != nil {
		t.Errorf("Não esperava erro ao deletar, recebeu: %v", err)
	}
}
