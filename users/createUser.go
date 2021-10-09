package users

import (
	"appointy/dbservice"
	"context"
	"fmt"
	"log"
    	"crypto/sha256"  // Used for hashing passwords

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(newUser User) *mongo.InsertOneResult {
	client, error := dbservice.GetMongoClient()
	fmt.Println(client, error)
	var userCollection = client.Database(dbservice.DB).Collection("users") //Connect to local MongoDB
	pwd := newUser.Password
	h := sha256.New()
	h.Write([]byte(pwd))
	bs := h.Sum(nil)  // Hashed password
	newUser.Password = string (bs[:]) //Convert bytes array to string
	insertResult, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult
}
