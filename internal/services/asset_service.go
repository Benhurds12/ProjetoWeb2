package services

import (
	"context"
	"database/sql"
	"errors"

	"projetoweb2/internal/db"

	"github.com/google/uuid"
)

type AssetService struct {
	Queries db.Querier
}

func NewAssetService(q db.Querier) *AssetService {
	return &AssetService{Queries: q}
}

func (s *AssetService) CreateBem(ctx context.Context, nome, tipo, status string, setorID *int32, fornecedorID *int32, fabricanteID *int32) (db.Ben, error) {
	if nome == "" || tipo == "" {
		return db.Ben{}, errors.New("nome e tipo são obrigatórios")
	}

	newID := uuid.New()

	if status == "" {
		status = "OCIOSO"
	}

	var sector sql.NullInt32
	var supplier sql.NullInt32
	var manufacturer sql.NullInt32

	if setorID != nil {
		_, err := s.Queries.GetSetorByID(ctx, *setorID)
		if err != nil {
			return db.Ben{}, errors.New("setor não encontrado")
		}
		sector = sql.NullInt32{Int32: *setorID, Valid: true}
	}

	if fornecedorID != nil {
		_, err := s.Queries.GetFornecedorByID(ctx, *fornecedorID)
		if err != nil {
			return db.Ben{}, errors.New("fornecedor não encontrado")
		}

		supplier = sql.NullInt32{
			Int32: *fornecedorID,
			Valid: true,
		}
	}

	if fabricanteID != nil {
		_, err := s.Queries.GetFabricanteByID(ctx, *fabricanteID)
		if err != nil {
			return db.Ben{}, errors.New("fabricante não encontrado")
		}

		manufacturer = sql.NullInt32{
			Int32: *fabricanteID,
			Valid: true,
		}
	}

	return s.Queries.CreateBem(ctx, db.CreateBemParams{
		ID:           newID,
		Nome:         nome,
		Status:       sql.NullString{String: status, Valid: true},
		Tipo:         tipo,
		SetorID:      sector,
		FornecedorID: supplier,
		FabricanteID: manufacturer,
	})
}

func (s *AssetService) GetBem(ctx context.Context, id uuid.UUID) (db.Ben, error) {
	bem, err := s.Queries.GetBemByID(ctx, id)
	if err != nil {
		return db.Ben{}, errors.New("bem não encontrado")
	}
	return bem, nil
}

func (s *AssetService) ListBens(ctx context.Context) ([]db.Ben, error) {
	return s.Queries.ListBens(ctx)
}

func (s *AssetService) UpdateBem(ctx context.Context, id uuid.UUID, nome, tipo, status string, setorID *int32, fornecedorID *int32, fabricanteID *int32) (db.Ben, error) {
	if nome == "" || tipo == "" {
		return db.Ben{}, errors.New("nome e tipo são obrigatórios")
	}

	var sector sql.NullInt32
	var supplier sql.NullInt32
	var manufacturer sql.NullInt32

	if setorID != nil {
		_, err := s.Queries.GetSetorByID(ctx, *setorID)
		if err != nil {
			return db.Ben{}, errors.New("setor não encontrado")
		}
		sector = sql.NullInt32{Int32: *setorID, Valid: true}
	}

	if fornecedorID != nil {
		_, err := s.Queries.GetFornecedorByID(ctx, *fornecedorID)
		if err != nil {
			return db.Ben{}, errors.New("fornecedor não encontrado")
		}

		supplier = sql.NullInt32{
			Int32: *fornecedorID,
			Valid: true,
		}
	}

	if fabricanteID != nil {
		_, err := s.Queries.GetFabricanteByID(ctx, *fabricanteID)
		if err != nil {
			return db.Ben{}, errors.New("fabricante não encontrado")
		}

		manufacturer = sql.NullInt32{
			Int32: *fabricanteID,
			Valid: true,
		}
	}

	return s.Queries.UpdateBem(ctx, db.UpdateBemParams{
		ID:           id,
		Nome:         nome,
		Status:       sql.NullString{String: status, Valid: true},
		Tipo:         tipo,
		SetorID:      sector,
		FornecedorID: supplier,
		FabricanteID: manufacturer,
	})
}

func (s *AssetService) DeleteBem(ctx context.Context, id uuid.UUID) error {
	return s.Queries.DeleteBem(ctx, id)
}
