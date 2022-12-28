package database

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var Conn *Connection = nil

type Connection struct {
	Conn *sql.DB
}

func init() {
	if Conn == nil {
		db, err := sql.Open("sqlite3", "./test.db")

		if err != nil {
			log.Fatalf("Failed to connect with the database: %v\n", err)
		}

		conn := &Connection{
			Conn: db,
		}

		log.Info("Connected with the database")
		Conn = conn
	}
}

func (c *Connection) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return c.Conn.ExecContext(ctx, query, args...)
}

func (c *Connection) ExecRow(ctx context.Context, query string, args ...any) *sql.Row {
	return c.Conn.QueryRowContext(ctx, query, args...)
}

func (c *Connection) ExecRows(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return c.Conn.QueryContext(ctx, query, args...)
}
