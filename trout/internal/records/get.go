package records

import (
	"context"
	"fmt"

	"github.com/eskpil/cod/trout/database"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAll(ctx context.Context, zoneId string) ([]database.Record, error) {
	records := []database.Record{}

	filter := bson.D{
		{"zone_id", zoneId},
	}

	cursor, err := getCollection().Find(ctx, filter)

	if err != nil {
		return records, err
	}

	if err := cursor.All(ctx, &records); err != nil {
		return records, err
	}

	return records, nil
}

func SearchSpecific(ctx context.Context, recordType uint16, fqdn string) (database.Record, bool, error) {
	filter := bson.D{
		{"fqdn", fqdn},
		{"type", recordType},
	}

	cursor, err := getCollection().Find(ctx, filter)

	if err != nil {
		return database.Record{}, false, err
	}

	var records []database.Record

	if err := cursor.All(ctx, &records); err != nil {
		return database.Record{}, false, err
	}

	if 0 >= len(records) {
		return database.Record{}, false, nil
	}

	return records[0], true, nil
}

func Search(ctx context.Context, fqdn string) ([]database.Record, error) {
	records := new([]database.Record)
	filter := bson.D{
		{"fqdn", fqdn},
	}

	cursor, err := getCollection().Find(ctx, filter)

	if err != nil {
		return *records, err
	}

	if err := cursor.All(ctx, records); err != nil {
		return *records, fmt.Errorf("Could not find any records for fqdn: %s", fqdn)
	}

	return *records, nil
}
