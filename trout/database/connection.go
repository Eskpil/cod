package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	log "github.com/sirupsen/logrus"
)

var Conn *mongo.Database

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://cod:test@localhost:27017/cod_db")))

	if err != nil {
		log.Fatalf("Failed to create a new mongodb client: %s\n", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatalf("Failed to connect with the database: %s\n", err)
	}

	log.Info("Connected with the database")

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("Failed to ping the database: %s\n", err)
	}

	log.Info("Pinged the primary database")

	Conn = client.Database("cod_db", nil)
}
