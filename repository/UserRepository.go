package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/mateuslima90/grpc-go/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type userMongoDB struct {
	connection *mongo.Client
}

func NewMongo(connection *mongo.Client) UserRepository {
	return &userMongoDB{connection: connection}
}

const DBNAME = "go_db"
const COLLECTION = "users"

func (u userMongoDB) CreateUser(username string, name string, email string) (entities.User, error) {

	user := entities.UserDTO{Username: username, Name: name, Email: email}

	collection := u.connection.Database(DBNAME).Collection(COLLECTION)

	insertResult, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		fmt.Println("Failed to add new user")
		//log.Fatal(err)
	}

	fmt.Println("Inserted user with ID:", insertResult.InsertedID)

	oid := insertResult.InsertedID.(primitive.ObjectID).Hex()

	return entities.User{ObjectID: oid, Username: user.Username, Name: user.Name, Email: user.Email}, nil
}

func (u userMongoDB) GetUserByUsername(username string) (entities.User, error) {

	collection := u.connection.Database(DBNAME).Collection(COLLECTION)

	filter := bson.M{"username": username}

	var user entities.User

	err := collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found user with email ", user.Email)
	return user, nil
}

func (u userMongoDB) GetUserById(id string) (entities.User, error) {

	collection := u.connection.Database(DBNAME).Collection(COLLECTION)

	fmt.Println(id)

	objectId, errObjectID := primitive.ObjectIDFromHex(id)

	if errObjectID != nil {
		log.Println("Id bad formation")
		//log.Fatal(errObjectID)
	}

	filter := bson.M{"_id": objectId}

	var user entities.User

	result := collection.FindOne(context.Background(), filter)

	fmt.Println(result)
	result.Decode(&user)
	if user.Username == "" && user.Email == "" {
		return user, errors.New("Id not found")
	}

	fmt.Println("Found user with email ", user.Email)
	return user, nil
}

func (u userMongoDB) GetAllUser() ([]entities.User, error) {

	collection := u.connection.Database(DBNAME).Collection(COLLECTION)
	var users []entities.User
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
