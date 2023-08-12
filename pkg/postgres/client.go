package postgres

import (
	"context"
	"github.com/Tobiska/avito-internship/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewClient(cfg *config.Database) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), cfg.Dsn)
}
