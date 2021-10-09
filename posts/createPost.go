package posts

import (
	"appointy/dbservice"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePost(newPost Post) *mongo.InsertOneResult {
	client, error := dbservice.GetMongoClient()
	fmt.Println(client, error)
	var postCollection = client.Database(dbservice.DB).Collection("posts")
	insertResult, err := postCollection.InsertOne(context.TODO(), newPost)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult
}

