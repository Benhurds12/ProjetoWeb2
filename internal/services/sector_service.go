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

type BemOut struct {
	ID     string `json:"id"`
	Nome   string `json:"nome"`
	Tipo   string `json:"tipo"`
	Status string `json:"status"`
}

type SetorWithBensOut struct {
	ID    int32    `json:"id"`
	Nome  string   `json:"nome"`
	Local string   `json:"local"`
	Bens  []BemOut `json:"bens"`
}

// ListSetoresWithBens busca os setores e agrupa os bens pertencentes a cada um
func (s *SectorService) ListSetoresWithBens(ctx context.Context) ([]SetorWithBensOut, error) {
	rows, err := s.Queries.ListSetoresWithBens(ctx)
	if err != nil {
		return nil, err
	}

	result := []SetorWithBensOut{}
	index := map[int32]int{}

	for _, row := range rows {
		i, exists := index[row.SetorID]

		if !exists {
			i = len(result)
			result = append(result, SetorWithBensOut{
				ID:    row.SetorID,
				Nome:  row.SetorNome,
				Local: row.SetorLocal,
				Bens:  []BemOut{},
			})
			index[row.SetorID] = i
		}

		// Como usamos LEFT JOIN, os campos do "Bem" virão nulos se o setor for vazio.
		if row.BemID.Valid {
			result[i].Bens = append(result[i].Bens, BemOut{
				ID:     row.BemID.UUID.String(),
				Nome:   row.BemNome.String,
				Tipo:   row.BemTipo.String,
				Status: row.BemStatus.String,
			})
		}
	}

	return result, nil
}
