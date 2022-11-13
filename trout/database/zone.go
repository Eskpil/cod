package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/sirupsen/logrus"
)

type Zone struct {
	Id   string `json:"id" bson:"_id"`
	Fqdn string `json:"fqdn" bson:"fqdn"`

	Name string `json:"name" bson:"name"`
}

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mod := mongo.IndexModel{
		Keys: bson.M{
			"fqdn": 1, // index in ascending order
		},
		Options: &options.IndexOptions{
			Unique: func(b bool) *bool { return &b }(true),
		},
	}

	if name, err := Conn.Collection("zones").Indexes().CreateOne(ctx, mod); err != nil {
		log.Fatalf("Failed to create index: %s on collection \"zones\" because: %v\n", name, err)
	} else {
		log.Infof("Successfully created index: %s on collection \"zones\"\n", name)
	}
}
