package records

import (
	"github.com/eskpil/cod/trout/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCollection() *mongo.Collection {
	return database.Conn.Collection("records")
}
