package database

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

// type Interface struct {
// 	Id      string   `json:"id" yaml:"id"`
// 	Name    string   `json:"name" yaml:"name"`
// 	Mac     string   `json:"mac" yaml:"mac"`
// 	IpAddrs []IpAddr `json:"addrs" yaml:"addrs"`
// }

// type IpAddr struct {
// 	Type   int32  `json:"type" yaml:"type"`
// 	Addr   string `json:"addr" yaml:"addr"`
// 	Prefix uint32 `json:"prefix" yaml:"prefix"`
// }

type IpAddrModel struct {
	Type   int32
	Addr   string
	Prefix uint32
}

type InterfaceModel struct {
	Id        string
	Mac       string
	MachineId string
	Name      string
}

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Interfaces
	{
		_, err := Conn.Exec(ctx, `
			CREATE TABLE IF NOT EXISTS interfaces(
				id TEXT NOT NULL PRIMARY KEY,
				mac TEXT NOT NULL,
				machineid TEXT NOT NULL,
				name TEXT NOT NULL
			)
		`)

		if err != nil {
			log.Fatalf("Failed to create the \"interfaces\" table: %v\n", err)
		}

		log.Infof("Created or verified that table \"interfaces\" exists\n")
	}

	// addrs
	{
		_, err := Conn.Exec(ctx, `
			CREATE TABLE IF NOT EXISTS addrs(
				id TEXT NOT NULL PRIMARY KEY,
				ifaceid TEXT NOT NULL,
				type INTEGER NOT NULL,
				addr STRING NOT NULL,
				prefix INTEGER NOT NULL
			)
		`)

		if err != nil {
			log.Fatalf("Failed to create the \"addrs\" table: %v\n", err)
		}

		log.Infof("Created or verified that tabel \"addrs\" exists\n")
	}
}
