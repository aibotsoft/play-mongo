package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

type Podcast struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Title         string             `bson:"title,omitempty"`
	Author        string             `bson:"author,omitempty"`
	NewFeature    string             `bson:"fu,omitempty"`
	SecondFeature string             `bson:"fu,omitempty"`
	Tags          []string           `bson:"tags,omitempty"`
}

func main() {
	fmt.Println(os.Getenv("ATLAS_URI"))
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("ATLAS_URI")))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	database := client.Database("quickstart")
	podcastsCollection := database.Collection("podcasts")
	fmt.Println(time.Since(start))
	//episodesCollection := database.Collection("episodes")
	//err = client.Ping(ctx, readpref.Primary())
	podcast := Podcast{
		Title:  "The Polyglot Developer",
		Author: "Nic Raboy",
		Tags:   []string{"development", "programming", "coding"},
	}
	insertResult, err := podcastsCollection.InsertOne(ctx, podcast)
	if err != nil {
		panic(err)
	}
	fmt.Println(time.Since(start))

	fmt.Println(insertResult.InsertedID)

}
