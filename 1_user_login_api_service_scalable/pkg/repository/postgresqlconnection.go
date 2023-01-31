package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Conn interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func NewPostgreSqlConn(dbUrl string) (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), dbUrl)
}
