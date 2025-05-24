package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func NewPostgresConn(
	ctx context.Context,
	address string,
) *pgx.Conn {
	conn, err := pgx.Connect(ctx, address)
	if err != nil {
		panic(fmt.Errorf("connect to postgres: %w", err))
	}

	return conn
}
