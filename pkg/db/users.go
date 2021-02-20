package db

import (
	"context"
	"os"
	"time"

	"github.com/sgkul2000/mail-for-cause/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser  creates a user
func CreateUser(u types.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return primitive.NilObjectID, err
	}
	col := client.Database("mail-for-cause").Collection("users")
	u.EncryptPassword()
	data, err := bson.Marshal(u)
	if err != nil {
		return primitive.NilObjectID, err
	}

	result, err := col.InsertOne(ctx, data)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

// FindUser finds a single user
func FindUser(email string) (types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return types.User{}, err
	}
	col := client.Database("mail-for-cause").Collection("users")

	result := col.FindOne(ctx, bson.M{
		"email": email,
	})
	if result.Err() != nil {
		return types.User{}, err
	}
	var document bson.M = make(bson.M)
	err = result.Decode(document)
	if err != nil {
		return types.User{}, err
	}
	bytes, err := bson.Marshal(document)
	if err != nil {
		return types.User{}, err
	}

	var newUser types.User
	bson.Unmarshal(bytes, &newUser)
	return newUser, nil
}
