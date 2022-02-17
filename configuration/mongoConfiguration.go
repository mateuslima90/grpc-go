package configuration

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnection() (*mongo.Client, error) {

	uriDB := os.Getenv("URIDB")

	uri := "mongodb://"+uriDB+"/go_db"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://172.17.0.3:27017/go_db"))
	if err != nil {
		log.Fatal("MONGOCONNECTION: mongo.NewClient", err)
	}
	log.Printf("Uri connection mongodb://%v/go_db\n", uriDB)
	log.Println("MongoDB Connect")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("MONGOCONNECTION: CREATING A CONNECTION", err)
	}

	return client, nil
}

