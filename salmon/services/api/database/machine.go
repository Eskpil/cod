package database

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

type MachineModel struct {
	Id       string
	Name     string
	Host     string
	Hostname string
	Groups   string // array seperated by ;;-_-;;
	Fqdn     string
}

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := Conn.Exec(ctx, `
	CREATE TABLE IF NOT EXISTS machines(
		id TEXT NOT NULL PRIMARY KEY,
		name TEXT NOT NULL,
		host TEXT NOT NULL,
		hostname TEXT NOT NULL,
		groups TEXT,
		fqdn TEXT NOT NULL
	)
	`)

	if err != nil {
		log.Fatalf("Failed to create the \"machines\" table: %v\n", err)
	}

	log.Infof("Created or verified that table \"machines\" exists\n")
}
