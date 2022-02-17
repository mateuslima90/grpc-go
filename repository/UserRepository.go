package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/mateuslima90/grpc-go/configuration"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/mateuslima90/grpc-go/entities"
)



const DBNAME = "go_db"

func InsertUser(username string, name string, email string) (User, error) {
	client, errConnection := configuration.MongoConnection()

	if errConnection != nil {
		return User{}, errConnection
	}

	user := UserDTO{Username: username, Name: name, Email: email}

	collection := client.Database(DBNAME).Collection("users.go")

	insertResult, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		fmt.Println("Failed to add new user")
		//log.Fatal(err)
	}

	fmt.Println("Inserted user with ID:", insertResult.InsertedID)

	oid := insertResult.InsertedID.(primitive.ObjectID).Hex()

	return User{ObjectID: oid, Username: user.Username, Name: user.Name, Email: user.Email}, nil
}

func GetUserByUsername(username string) (User, error) {
	client, errConnection := configuration.MongoConnection()

	if errConnection != nil {
		return User{}, errConnection
	}

	collection := client.Database(DBNAME).Collection("users.go")

	filter := bson.M{"username": username}

	var user User

	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found user with email ", user.Email)
	return user, nil
}

func GetUserById(id string) (User, error) {
	client, errConnection := configuration.MongoConnection()

	if errConnection != nil {
		return User{}, errConnection
	}

	collection := client.Database(DBNAME).Collection("users.go")

	fmt.Println(id)

	objectId, errObjectID := primitive.ObjectIDFromHex(id)

	if errObjectID != nil {
		log.Fatal(errObjectID)
	}

	filter := bson.M{"_id": objectId}

	var user User

	result := collection.FindOne(context.Background(), filter)

	fmt.Println(result)
	result.Decode(&user)

	fmt.Println("Found user with email ", user.Email)
	return user, nil
}

func GetAllUser() ([]User, error) {
	client, errConnection := configuration.MongoConnection()

	if errConnection != nil {
		return []User{}, errConnection
	}

	collection := client.Database(DBNAME).Collection("users.go")
	var users []User
	//cursor, err := collection.Find(context.Background(), bson.M{})

	// Find all users.go
	userCursor, errCursor := collection.Find(context.Background(), bson.M{})
	if errCursor != nil {
		return nil, errCursor
	}

	errUserCursor := userCursor.All(context.Background(), &users)
	if errUserCursor != nil {
		return nil, errUserCursor
	}

	return users, nil
}
