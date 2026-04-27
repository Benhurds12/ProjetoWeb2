package services

import (
	"context"
	"database/sql"
	"errors"

	"projetoweb2/internal/db"
)

type SupplierService struct {
	Queries *db.Queries
}

func NewSupplierService(q *db.Queries) *SupplierService {
	return &SupplierService{Queries: q}
}

func (s *SupplierService) CreateFornecedor(ctx context.Context, nome, cnpj, contato string) (db.Fornecedore, error) {
	if nome == "" || cnpj == "" {
		return db.Fornecedore{}, errors.New("nome e cnpj são obrigatórios")
	}

	_, err := s.Queries.GetFornecedorByCnpj(ctx, cnpj)
	if err == nil {
		return db.Fornecedore{}, errors.New("cnpj já cadastrado")
	}
	if err != sql.ErrNoRows {
		return db.Fornecedore{}, err
	}

	return s.Queries.CreateFornecedor(ctx, db.CreateFornecedorParams{
		Nome:    nome,
		Cnpj:    cnpj,
		Contato: contato,
	})
}

func (s *SupplierService) ListFornecedores(ctx context.Context) ([]db.Fornecedore, error) {
	return s.Queries.ListFornecedores(ctx)
}

func (s *SupplierService) GetFornecedor(ctx context.Context, id int32) (db.Fornecedore, error) {
	f, err := s.Queries.GetFornecedorByID(ctx, id)
	if err != nil {
		return db.Fornecedore{}, errors.New("fornecedor não encontrado")
	}
	return f, nil
}

func (s *SupplierService) UpdateFornecedor(ctx context.Context, id int32, nome, cnpj, contato string) (db.Fornecedore, error) {
	if nome == "" || cnpj == "" {
		return db.Fornecedore{}, errors.New("nome e cnpj são obrigatórios")
	}
	existing, err := s.Queries.GetFornecedorByCnpj(ctx, cnpj)
	if err == nil && existing.ID != id {
		return db.Fornecedore{}, errors.New("cnpj já cadastrado")
	}
	if err != nil && err != sql.ErrNoRows {
		return db.Fornecedore{}, err
	}

	return s.Queries.UpdateFornecedor(ctx, db.UpdateFornecedorParams{
		ID:      id,
		Nome:    nome,
		Cnpj:    cnpj,
		Contato: contato,
	})
}

func (s *SupplierService) DeleteFornecedor(ctx context.Context, id int32) error {
	return s.Queries.DeleteFornecedor(ctx, id)
}
