package services

import (
	"context"
	"database/sql"
	"errors"

	"projetoweb2/internal/db"
)

type ManufacturerService struct {
	Queries db.Querier
}

func NewManufacturerService(q db.Querier) *ManufacturerService {
	return &ManufacturerService{Queries: q}
}

func (s *ManufacturerService) CreateFabricante(ctx context.Context, nome, cnpj string) (db.Fabricante, error) {
	if nome == "" || cnpj == "" {
		return db.Fabricante{}, errors.New("nome e cnpj são obrigatórios")
	}

	_, err := s.Queries.GetFabricanteByCnpj(ctx, cnpj)
	if err == nil {
		return db.Fabricante{}, errors.New("cnpj já cadastrado")
	}
	if err != sql.ErrNoRows {
		return db.Fabricante{}, err
	}

	return s.Queries.CreateFabricante(ctx, db.CreateFabricanteParams{
		Nome: nome,
		Cnpj: cnpj,
	})
}

func (s *ManufacturerService) ListFabricantes(ctx context.Context) ([]db.Fabricante, error) {
	return s.Queries.ListFabricantes(ctx)
}

func (s *ManufacturerService) GetFabricante(ctx context.Context, id int32) (db.Fabricante, error) {
	f, err := s.Queries.GetFabricanteByID(ctx, id)
	if err != nil {
		return db.Fabricante{}, errors.New("fabricante não encontrado")
	}
	return f, nil
}

func (s *ManufacturerService) UpdateFabricante(ctx context.Context, id int32, nome, cnpj string) (db.Fabricante, error) {
	if nome == "" || cnpj == "" {
		return db.Fabricante{}, errors.New("nome e cnpj são obrigatórios")
	}

	existing, err := s.Queries.GetFabricanteByCnpj(ctx, cnpj)
	if err == nil && existing.ID != id {
		return db.Fabricante{}, errors.New("cnpj já cadastrado")
	}
	if err != nil && err != sql.ErrNoRows {
		return db.Fabricante{}, err
	}

	return s.Queries.UpdateFabricante(ctx, db.UpdateFabricanteParams{
		ID:   id,
		Nome: nome,
		Cnpj: cnpj,
	})
}

func (s *ManufacturerService) DeleteFabricante(ctx context.Context, id int32) error {
	return s.Queries.DeleteFabricante(ctx, id)
}
