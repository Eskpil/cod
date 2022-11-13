package zoneService

import (
	"context"

	"github.com/eskpil/cod/trout/database"
)

func Create(ctx context.Context, zone database.Zone) error {
	_, err := getCollection().InsertOne(ctx, zone)
	return err
}
