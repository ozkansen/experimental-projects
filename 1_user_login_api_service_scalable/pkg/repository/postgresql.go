package repository

import (
	"context"
)

var _ Repository = (*PostgreSqlDB)(nil)

type PostgreSqlDB struct {
	conn Conn
}

func NewPostgreSqlDB(conn Conn) *PostgreSqlDB {
	return &PostgreSqlDB{conn: conn}
}

func (p *PostgreSqlDB) GetByUsername(ctx context.Context, username string) (*User, error) {
	sql := `SELECT id,username,passwd FROM "public".users WHERE username=$1`
	var user User
	if err := p.conn.QueryRow(ctx, sql, username).Scan(&user.ID, &user.Username, &user.Passwd); err != nil {
		return nil, err
	}
	return &user, nil
}
