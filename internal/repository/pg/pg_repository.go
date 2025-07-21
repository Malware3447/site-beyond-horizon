package pg

import (
	"context"
	"sbh/internal/models/pg"
)

type Repository interface {
	AddUser(ctx context.Context, params *pg.UserDataParams) error
}
