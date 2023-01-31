package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Conn interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func NewPostgreSqlConn(dbUrl string) (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), dbUrl)
}
