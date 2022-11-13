package zoneService

import (
	"context"

	"github.com/eskpil/cod/trout/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAll(ctx context.Context) ([]database.Zone, error) {
	zones := []database.Zone{}

	projection := bson.D{
		{"records", 0},
	}

	cursor, err := getCollection().Find(ctx, bson.D{}, options.Find().SetProjection(projection))
	if err != nil {
		return zones, err
	}

	if err = cursor.All(ctx, &zones); err != nil {
		return zones, err
	}

	return zones, nil
}

func GetById(ctx context.Context, zoneId string) (database.Zone, error) {
	var zone database.Zone

	filter := bson.D{
		{"_id", zoneId},
	}
	projection := bson.D{
		{"records", 0},
	}

	err := getCollection().FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&zone)
	if err != nil {
		return zone, err
	}

	return zone, nil
}
