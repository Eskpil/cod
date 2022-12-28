package database

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

// type StoragePool struct {
// 	Id         string `json:"id"`
// 	Name       string `json:"name"`
// 	TargetPath string `json:"target_path"`
//	Host	   string `json:"host"`
// }

type StoragePoolModel struct {
	Id         string
	Name       string
	TargetPath string
	Host       string
}

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := Conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS storagePools(
			id TEXT NOT NULL PRIMARY KEY,
			name TEXT NOT NULL,
			target_path TEXT NOT NULL,
			host TEXT NOT NULL
		)	
	`)

	if err != nil {
		log.Fatalf("Failed to create the \"storagePools\" table: %v\n", err)
	}

	log.Infof("Created or verified that table \"storagePools\" exists\n")
}
