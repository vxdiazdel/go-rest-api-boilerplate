package stores

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/vxdiazdel/rest-api/internal/logger"
)

type PostgresStore struct {
	ctx context.Context
	db  *pgx.Conn
	lg  logger.ILogger
}

func NewPostgresStore(
	ctx context.Context,
	db *pgx.Conn,
	lg logger.ILogger,
) *PostgresStore {
	return &PostgresStore{
		ctx: ctx,
		db:  db,
		lg:  lg,
	}
}

func (s *PostgresStore) Ping(ctx context.Context) error {
	return nil
}

func (s *PostgresStore) Ctx() context.Context {
	return s.ctx
}

func (s *PostgresStore) DB() *pgx.Conn {
	return s.db
}

func (s *PostgresStore) Lg() logger.ILogger {
	return s.lg
}
