package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/mateuslima90/grpc-go/configuration"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string `json:"Username,omitempty"`

	Name string `json:"Name,omitempty"`

	Email string `json:"Email,omitempty"`
}

const DBNAME = "go_db"

// func mongoConnection() {
// 	mongoTemplate, err := configuration.MongoConnection()
// 	if err != nil {
// 		log.Fatalf("Could not be connect with database")
// 	}

// 	return mongoTemplate
// }

func InsertUser(name string, email string) {
	client := configuration.MongoConnection()
	user := User{"mateuslima90@gmail.com", name, email}

	collection := client.Database(DBNAME).Collection("users")

	insertResult, err := collection.InsertOne(context.TODO(), user)

	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("Inserted user with ID:", insertResult.InsertedID)

}

func InsertUser2(username string, name string, email string) {
	client := configuration.MongoConnection()
	user := User{username, name, email}

	collection := client.Database(DBNAME).Collection("users")

	insertResult, err := collection.InsertOne(context.TODO(), user)

	if err != nil {

		fmt.Println("Failed to add new user")
		//log.Fatal(err)

	}

	fmt.Println("Inserted user with ID:", insertResult.InsertedID)

}

func GetUser(username string) User {
	client := configuration.MongoConnection()
	collection := client.Database(DBNAME).Collection("users")

	filter := bson.M{"username": username}

	var user User

	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("Found user with email ", user.Email)
	return user
}
