package storagePools

import (
	"context"

	"github.com/eskpil/salmon/pkg/models"
	"github.com/eskpil/salmon/services/api/database"
)

func GetById(ctx context.Context, poolId string) (models.StoragePool, error) {
	pool := models.StoragePool{}

	row := database.Conn.ExecRow(ctx, `SELECT id, name, target_path, host FROM storagePools WHERE id = ?`, poolId)
	if err := row.Scan(&pool.Id, &pool.Name, &pool.TargetPath, &pool.Host); err != nil {
		return pool, err
	}

	return pool, nil
}
