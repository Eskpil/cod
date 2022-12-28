package interfaces

import (
	"context"

	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/database"
	"github.com/google/uuid"
)

func createInterface(ctx context.Context, mac string, addr models.IpAddr) error {
	if _, err := database.Conn.Exec(ctx, "INSERT INTO addrs VALUES(?, ?, ?, ?, ?)", uuid.New().String(), mac, addr.Type, addr.Addr, addr.Prefix); err != nil {
		return err
	}
	return nil
}

func Create(ctx context.Context, machineId string, iface models.Interface) error {
	id := uuid.New().String()
	if _, err := database.Conn.Exec(ctx, "INSERT INTO interfaces VALUES(?, ?, ?, ?)", id, iface.Mac, machineId, iface.Name); err != nil {
		return err
	}

	for _, addr := range iface.IpAddrs {
		if err := createInterface(ctx, id, addr); err != nil {
			return err
		}
	}

	return nil
}
