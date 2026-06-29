package services

import (
	"context"
	"database/sql"
	"projetoweb2/internal/db"
	"testing"

	"github.com/google/uuid"
)

type MockDB struct {
	db.Querier

	stubGetSetorByNome db.Setore
	errGetSetorByNome  error

	stubCreateSetor db.Setore
	errCreateSetor  error

	// NOVOS CAMPOS PARA GET, UPDATE E DELETE
	stubGetSetorByID db.Setore
	errGetSetorByID  error

	stubUpdateSetor db.Setore
	errUpdateSetor  error

	errDeleteSetor error

	stubListSetores []db.Setore
	errListSetores  error

	stubListSetoresWithBens []db.ListSetoresWithBensRow
	errListSetoresWithBens  error
}

func (m *MockDB) GetSetorByNome(ctx context.Context, nome string) (db.Setore, error) {
	return m.stubGetSetorByNome, m.errGetSetorByNome
}

func (m *MockDB) CreateSetor(ctx context.Context, arg db.CreateSetorParams) (db.Setore, error) {
	return m.stubCreateSetor, m.errCreateSetor
}

func (m *MockDB) GetSetorByID(ctx context.Context, id int32) (db.Setore, error) {
	return m.stubGetSetorByID, m.errGetSetorByID
}

func (m *MockDB) UpdateSetor(ctx context.Context, arg db.UpdateSetorParams) (db.Setore, error) {
	return m.stubUpdateSetor, m.errUpdateSetor
}

func (m *MockDB) DeleteSetor(ctx context.Context, id int32) error {
	return m.errDeleteSetor
}

func (m *MockDB) ListSetores(ctx context.Context) ([]db.Setore, error) {
	return m.stubListSetores, m.errListSetores
}

func (m *MockDB) ListSetoresWithBens(ctx context.Context) ([]db.ListSetoresWithBensRow, error) {
	return m.stubListSetoresWithBens, m.errListSetoresWithBens
}

func TestCreateSetor_Sucesso(t *testing.T) {
	// Preparamos o nosso banco falso dizendo que o setor AINDA NÃO EXISTE (ErrNoRows)
	mock := &MockDB{
		errGetSetorByNome: sql.ErrNoRows,
		stubCreateSetor: db.Setore{
			ID:    1,
			Nome:  "TI",
			Local: "Sala 1",
		},
	}

	service := NewSectorService(mock)

	// Tentamos criar o setor
	setor, err := service.CreateSetor(context.Background(), "TI", "Sala 1")

	// Verificações (Asserts)
	if err != nil {
		t.Errorf("Não esperava erro, mas recebeu: %v", err)
	}
	if setor.ID != 1 {
		t.Errorf("Esperava ID 1, recebeu %d", setor.ID)
	}
}

func TestCreateSetor_ErroNomeDuplicado(t *testing.T) {
	// Preparamos o banco falso dizendo que o setor JÁ EXISTE
	mock := &MockDB{
		errGetSetorByNome: nil, // Sem erro = encontrou no banco!
		stubGetSetorByNome: db.Setore{
			ID:   1,
			Nome: "TI",
		},
	}

	service := NewSectorService(mock)

	// Tentamos criar um setor com o mesmo nome
	_, err := service.CreateSetor(context.Background(), "TI", "Sala 1")

	// Verificações
	if err == nil {
		t.Errorf("Esperava um erro de setor duplicado, mas não deu erro")
	}
	if err.Error() != "já existe um setor com esse nome" {
		t.Errorf("Mensagem de erro errada: %v", err)
	}
}

func TestCreateSetor_ErroCamposVazios(t *testing.T) {
	// Aqui não precisamos configurar retorno no mock,
	// porque a validação barra antes de chegar no banco!
	mock := &MockDB{}
	service := NewSectorService(mock)

	// Tentamos criar um setor sem nome
	_, err := service.CreateSetor(context.Background(), "", "Sala 1")

	if err == nil {
		t.Errorf("Esperava erro de campos obrigatórios, mas passou")
	}
	if err.Error() != "nome e local são obrigatórios" {
		t.Errorf("Mensagem de erro diferente do esperado: %v", err)
	}
}

func TestGetSetor_Sucesso(t *testing.T) {
	mock := &MockDB{
		stubGetSetorByID: db.Setore{ID: 1, Nome: "RH", Local: "Sala 2"},
		errGetSetorByID:  nil,
	}
	service := NewSectorService(mock)

	setor, err := service.GetSetor(context.Background(), 1)

	if err != nil {
		t.Errorf("Não esperava erro: %v", err)
	}
	if setor.Nome != "RH" {
		t.Errorf("Esperava nome RH, recebeu %s", setor.Nome)
	}
}

func TestGetSetor_NaoEncontrado(t *testing.T) {
	mock := &MockDB{
		errGetSetorByID: sql.ErrNoRows, // Simula que o banco não achou nada
	}
	service := NewSectorService(mock)

	_, err := service.GetSetor(context.Background(), 99)

	if err == nil || err.Error() != "setor não encontrado" {
		t.Errorf("Esperava erro 'setor não encontrado', recebeu: %v", err)
	}
}

func TestUpdateSetor_Sucesso(t *testing.T) {
	mock := &MockDB{
		// Quando o service checar se o nome já existe, dizemos que não existe (ErrNoRows)
		errGetSetorByNome: sql.ErrNoRows,
		stubUpdateSetor:   db.Setore{ID: 1, Nome: "TI Atualizado", Local: "Sala 3"},
	}
	service := NewSectorService(mock)

	setor, err := service.UpdateSetor(context.Background(), 1, "TI Atualizado", "Sala 3")

	if err != nil {
		t.Errorf("Não esperava erro: %v", err)
	}
	if setor.Nome != "TI Atualizado" {
		t.Errorf("Update falhou, nome recebido: %s", setor.Nome)
	}
}

func TestUpdateSetor_ErroCamposVazios(t *testing.T) {
	mock := &MockDB{}
	service := NewSectorService(mock)

	_, err := service.UpdateSetor(context.Background(), 1, "", "Sala 3")

	if err == nil || err.Error() != "nome e local são obrigatórios" {
		t.Errorf("Esperava erro de validação, recebeu: %v", err)
	}
}

func TestUpdateSetor_ErroNomeDuplicado(t *testing.T) {
	mock := &MockDB{
		// Simulamos que o banco ACHOU um setor com o mesmo nome, e que o ID dele é 2 (diferente do 1 que queremos atualizar)
		errGetSetorByNome:  nil,
		stubGetSetorByNome: db.Setore{ID: 2, Nome: "TI"},
	}
	service := NewSectorService(mock)

	_, err := service.UpdateSetor(context.Background(), 1, "TI", "Sala 3")

	if err == nil || err.Error() != "já existe um setor com esse nome" {
		t.Errorf("Esperava erro de nome duplicado, recebeu: %v", err)
	}
}

func TestDeleteSetor_Sucesso(t *testing.T) {
	mock := &MockDB{
		errDeleteSetor: nil,
	}
	service := NewSectorService(mock)

	err := service.DeleteSetor(context.Background(), 1)

	if err != nil {
		t.Errorf("Não esperava erro ao deletar, recebeu: %v", err)
	}
}

func TestListSetoresWithBens_AgregacaoSucesso(t *testing.T) {
	// Vamos simular um UUID válido para o Bem
	bemIdValid := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")

	mock := &MockDB{
		// Simulamos o retorno do LEFT JOIN do banco de dados
		stubListSetoresWithBens: []db.ListSetoresWithBensRow{
			{
				SetorID:    1,
				SetorNome:  "TI",
				SetorLocal: "Sala 1",
				// Este setor TEM um bem
				BemID:     uuid.NullUUID{UUID: bemIdValid, Valid: true},
				BemNome:   sql.NullString{String: "Notebook Dell", Valid: true},
				BemTipo:   sql.NullString{String: "Computador", Valid: true},
				BemStatus: sql.NullString{String: "EM USO", Valid: true},
			},
			{
				SetorID:    2,
				SetorNome:  "RH",
				SetorLocal: "Sala 2",
				// Este setor NÃO TEM bens (Valid = false simula o NULL do banco)
				BemID: uuid.NullUUID{Valid: false},
			},
		},
	}

	service := NewSectorService(mock)

	resultado, err := service.ListSetoresWithBens(context.Background())

	// 1. Verifica se não deu erro
	if err != nil {
		t.Errorf("Não esperava erro, recebeu: %v", err)
	}

	// 2. Verifica se agrupou corretamente em 2 setores
	if len(resultado) != 2 {
		t.Fatalf("Esperava 2 setores, recebeu %d", len(resultado))
	}

	// 3. Verifica o Setor de TI (deve ter 1 bem aninhado)
	if resultado[0].Nome != "TI" || len(resultado[0].Bens) != 1 {
		t.Errorf("Esperava 1 bem no setor TI, recebeu %d", len(resultado[0].Bens))
	}
	if resultado[0].Bens[0].Nome != "Notebook Dell" {
		t.Errorf("Nome do bem agregado incorreto, recebeu: %s", resultado[0].Bens[0].Nome)
	}

	// 4. Verifica o Setor de RH (deve ter o slice de bens vazio, tamanho 0)
	if resultado[1].Nome != "RH" || len(resultado[1].Bens) != 0 {
		t.Errorf("Esperava 0 bens no setor RH, recebeu %d", len(resultado[1].Bens))
	}
}
