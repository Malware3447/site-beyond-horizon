package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"sbh/internal/models/pg"
	"time"
)

type PgRepository struct {
	db *pgxpool.Pool
}

func NewPgRepository(db *pgxpool.Pool) Repository {
	return &PgRepository{db: db}
}

func (p PgRepository) AddUser(ctx context.Context, params *pg.UserDataParams) error {
	log := logrus.New()
	const q = `
	INSERT INTO user_data (name, mail, countries_id, create_date)
	VALUES ($1, $2, $3, $4)
	`

	createdAt := time.Now()
	nowDate := fmt.Sprintf("%v", createdAt)

	_, err := p.db.Exec(ctx, q, params.Name, params.Mail, params.Countries_id, nowDate)
	if err != nil {
		log.Error(fmt.Sprintf("Ошибка добазу данных: %v", err))
		return fmt.Errorf("failed to upsert pair: %w", err)
	}

	return nil
}
