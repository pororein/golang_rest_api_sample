package user

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	EMail     string             `json:"e_mail" bson:"e_mail"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
}

func NewUser(email string, firstName string, lastName string) *User {
	return &User{
		EMail:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
}

func (u *User) Save(client *mongo.Client) (result *mongo.InsertOneResult, err error) {
	collection := client.Database("user_management").Collection("user")
	return collection.InsertOne(context.Background(), *u)
}

func (u *User) Update(client *mongo.Client, email string) (result *mongo.UpdateResult, err error) {
	collection := client.Database("user_management").Collection("user")
	return collection.UpdateOne(
		context.Background(),
		bson.D{primitive.E{Key: "e_mail", Value: email}},
		bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "e_mail", Value: u.EMail},
			primitive.E{Key: "first_name", Value: u.FirstName},
			primitive.E{Key: "last_name", Value: u.LastName}}}})
}

func (u *User) Load(client *mongo.Client, email string) error {
	collection := client.Database("user_management").Collection("user")
	return collection.FindOne(context.Background(), bson.D{{"e_mail", email}}).Decode(u)
}

func (u *User) Delete(client *mongo.Client, email string) (result *mongo.DeleteResult, err error) {
	collection := client.Database("user_management").Collection("user")
	return collection.DeleteOne(context.Background(), bson.D{{"e_mail", email}})
}
