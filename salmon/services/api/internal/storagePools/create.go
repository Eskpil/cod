package storagePools

import (
	"context"

	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/database"
)

func Create(ctx context.Context, pool models.StoragePool) error {
	if _, err := database.Conn.Exec(ctx, "INSERT INTO storagePools VALUES(?,?,?,?)", pool.Id, pool.Name, pool.TargetPath, pool.Host); err != nil {
		return err
	}
	return nil
}
