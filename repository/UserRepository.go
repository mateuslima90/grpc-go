package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/mateuslima90/grpc-go/configuration"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ObjectID string `bson:"_id" json:"_id"`

	Username string `json:"Username,omitempty"`

	Name string `json:"Name,omitempty"`

	Email string `json:"Email,omitempty"`
}

type UserDTO struct {
	Username string `json:"Username,omitempty"`

	Name string `json:"Name,omitempty"`

	Email string `json:"Email,omitempty"`
}

const DBNAME = "go_db"

func InsertUser(username string, name string, email string) User {
	client := configuration.MongoConnection()
	user := UserDTO{Username: username, Name: name, Email: email}

	collection := client.Database(DBNAME).Collection("users")

	insertResult, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		fmt.Println("Failed to add new user")
		//log.Fatal(err)
	}

	fmt.Println("Inserted user with ID:", insertResult.InsertedID)

	oid := insertResult.InsertedID.(primitive.ObjectID).Hex()

	return User{ObjectID: oid, Username: user.Username, Name: user.Name, Email: user.Email}
}

func GetUserByUsername(username string) User {
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

func GetUserById(id string) User {
	client := configuration.MongoConnection()
	collection := client.Database(DBNAME).Collection("users")

	fmt.Println(id)

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		log.Fatal(err)

	}

	filter := bson.M{"_id": objectId}

	var user User

	result := collection.FindOne(context.Background(), filter)

	fmt.Println(result)
	result.Decode(&user)

	fmt.Println("Found user with email ", user.Email)
	return user
}

func GetAllUser() []User {
	client := configuration.MongoConnection()
	collection := client.Database(DBNAME).Collection("users")
	var users []User
	//cursor, err := collection.Find(context.Background(), bson.M{})

	// Find all users
	userCursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil
	}

	err = userCursor.All(context.Background(), &users)
	if err != nil {
		return nil
	}

	return users

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cursor.Close(context.TODO())

	// for cursor.Next(context.TODO()) {
	// 	var user bson.M
	// 	if err = cursor.Decode(&user); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	users = []User{}
	// 	fmt.Println(users)
	// }
	// return users
}
