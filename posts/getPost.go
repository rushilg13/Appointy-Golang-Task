package posts

import (
	"appointy/dbservice"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPost(id string) Post{
	var post Post
	client, error := dbservice.GetMongoClient()
	if error != nil {	fmt.Println(error)  }
	var postCollection = client.Database(dbservice.DB).Collection("posts")
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	postCollection.FindOne(context.TODO(), filter).Decode(&post)
	return post
}
