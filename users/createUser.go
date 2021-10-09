package users

import (
	"appointy/dbservice"
	"context"
	"fmt"
	"log"
    "crypto/sha256"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(newUser User) *mongo.InsertOneResult {
	client, error := dbservice.GetMongoClient()
	fmt.Println(client, error)
	var userCollection = client.Database(dbservice.DB).Collection("users")
	pwd := newUser.Password
	h := sha256.New()
	h.Write([]byte(pwd))
	bs := h.Sum(nil)
	newUser.Password = string (bs[:])
	insertResult, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult
}
