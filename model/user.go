package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type User struct {
	ID        primitive.ObjectID `bson:_id`
	EMail     string             `bson:e_mail`
	FirstName string             `bson:first_name`
	LastName  string             `bson:last_name`
}

func NewUser(email string, firstName string, lastName string) *User {
	return &User{
		EMail:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (u *User) Save(client *mongo.Client) {
	collection := client.Database("user_management").Collection("user")
	collection.InsertOne(context.Background(), *u)
}

func (u *User) Load(client *mongo.Client, email string) error {
	collection := client.Database("user_management").Collection("user")
	return collection.FindOne(context.Background(), bson.M{"e_mail": email}).Decode(&u)
}

func (u *User) Delete(client *mongo.Client, objectIDHex string) error {
	collection := client.Database("user_management").Collection("user")
	objectID, _ := primitive.ObjectIDFromHex(objectIDHex)
	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
	return err
}
