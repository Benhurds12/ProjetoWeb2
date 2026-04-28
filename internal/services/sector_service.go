package services

import (
	"context"
	"database/sql"
	"errors"

	"projetoweb2/internal/db"
)

type SectorService struct {
	Queries db.Querier
}

func NewSectorService(q db.Querier) *SectorService {
	return &SectorService{Queries: q}
}

func (s *SectorService) CreateSetor(ctx context.Context, nome, local string) (db.Setore, error) {
	if nome == "" || local == "" {
		return db.Setore{}, errors.New("nome e local são obrigatórios")
	}

	_, err := s.Queries.GetSetorByNome(ctx, nome)
	if err == nil {
		return db.Setore{}, errors.New("já existe um setor com esse nome")
	}
	if err != sql.ErrNoRows {
		return db.Setore{}, err
	}

	return s.Queries.CreateSetor(ctx, db.CreateSetorParams{
		Nome:  nome,
		Local: local,
	})
}

func (s *SectorService) GetSetor(ctx context.Context, id int32) (db.Setore, error) {
	setor, err := s.Queries.GetSetorByID(ctx, id)
	if err != nil {
		return db.Setore{}, errors.New("setor não encontrado")
	}
	return setor, nil
}

func (s *SectorService) ListSetores(ctx context.Context) ([]db.Setore, error) {
	return s.Queries.ListSetores(ctx)
}

func (s *SectorService) UpdateSetor(ctx context.Context, id int32, nome, local string) (db.Setore, error) {
	if nome == "" || local == "" {
		return db.Setore{}, errors.New("nome e local são obrigatórios")
	}

	existing, err := s.Queries.GetSetorByNome(ctx, nome)
	if err == nil && existing.ID != id {
		return db.Setore{}, errors.New("já existe um setor com esse nome")
	}
	if err != nil && err != sql.ErrNoRows {
		return db.Setore{}, err
	}

	return s.Queries.UpdateSetor(ctx, db.UpdateSetorParams{
		ID:    id,
		Nome:  nome,
		Local: local,
	})
}

func (s *SectorService) DeleteSetor(ctx context.Context, id int32) error {
	return s.Queries.DeleteSetor(ctx, id)
}
