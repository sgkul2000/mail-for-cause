package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sgkul2000/mail-for-cause/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateCause creates a new cause and appends it to the users causes
func CreateCause(email string, cause *types.Cause) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return err
	}
	col := client.Database("mail-for-cause").Collection("users")
	_, err = col.UpdateOne(ctx, bson.M{"email": email}, bson.M{
		"$push": bson.M{"causes": cause},
	})
	if err != nil {
		return err
	}
	return nil
}

// EditCause edits a cause
func EditCause(email string, cause *types.Cause) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return err
	}
	col := client.Database("mail-for-cause").Collection("users")
	_, err = col.UpdateOne(ctx, bson.M{"email": email, "causes.name": cause.Name}, bson.M{
		"$set": bson.M{"causes.$": cause},
	})
	if err != nil {
		return err
	}
	return nil

}

// GetCauses edits a cause
func GetCauses() (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return nil, err
	}
	col := client.Database("mail-for-cause").Collection("users")
	result, err := col.Aggregate(ctx, mongo.Pipeline{
		bson.D{{"$group", bson.D{{"_id", nil}, {"causes", bson.D{{"$push", "$causes"}}}}}},
		bson.D{{"$project", bson.D{{"causes", 1}}}},
		bson.D{{"$unwind", "$causes"}},
		bson.D{{"$unwind", "$causes"}},
		bson.D{{"$group", bson.D{{"_id", nil}, {"causes", bson.D{{"$addToSet", "$causes"}}}}}},
		bson.D{{"$project", bson.D{{"causes", 1}}}},
	})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	err = result.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}
	return results[0]["causes"], nil
}

// GetCause gets a single cause by name
func GetCause(name string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return nil, err
	}
	col := client.Database("mail-for-cause").Collection("users")
	result := col.FindOne(ctx, bson.M{"causes.name": name})
	if err != nil {
		return nil, err
	}
	var user types.User
	err = result.Decode(&user)
	if err != nil {
		fmt.Println("Error Occured herre")
		return nil, err
	}
	for _, cause := range user.Causes {
		if cause.Name == name {
			return cause, nil
		}
	}
	return user, nil

}
