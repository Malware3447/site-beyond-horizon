package pg

import (
	"context"
	Model "sbh/internal/models/pg"
	"sbh/internal/repository/pg"
)

type PgService struct {
	repo pg.Repository
}

func NewPgService(repo pg.Repository) *PgService {
	return &PgService{repo: repo}
}

func (p PgService) AddUser(ctx context.Context, params *Model.UserDataParams) error {
	return p.repo.AddUser(ctx, params)
}
