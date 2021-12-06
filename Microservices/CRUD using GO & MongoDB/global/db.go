package global

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB mongo.Database
var collection *mongo.Collection

//DB holds database connection
func connectToDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal("Error connecting to DB")
	}
	DB = *client.Database(dbname)
}

// NewDBContext returns new context according to app performance
func NewDBContext(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d*performance/100)
}

// ConnectToTestDB overwrites DB with the Test Database
func ConnectToTestDB() {
	ctx, cancel := NewDBContext(10 * time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal("Error connect to DB: ", err.Error())
	}
	DB = *client.Database(dbname + "_test")
	ctx, cancel = NewDBContext(30 * time.Second)
	defer cancel()
	collections, _ := DB.ListCollectionNames(ctx, bson.M{})
	for _, collection := range collections {
		ctx, cancel = NewDBContext(10 * time.Second)
		defer cancel()
		DB.Collection(collection).Drop(ctx)
	}
}
