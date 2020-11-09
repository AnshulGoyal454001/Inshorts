package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//InsertPost is a insert function
func InsertPost(title string, subtitle string, content string) {
	//Connect Mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//Insert
	post := Article{title, subtitle, content}
	collection := client.Database("inShotsDB").Collection("posts")
	insertResult, err := collection.InsertOne(context.TODO(), post)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted post ID:", insertResult.InsertedID)
	fmt.Println("Inserted post ID:", title)
}

//GetPost is insteting in the database
func GetPost(idString string, w http.ResponseWriter) {
	fmt.Println("inside GetPost")
	//Connect Mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//Post Inside Database
	collection := client.Database("inShotsDB").Collection("posts")
	id, err := primitive.ObjectIDFromHex(idString)
	filter := bson.M{"_id": id}
	var post Article
	err2 := collection.FindOne(context.TODO(), filter).Decode(&post)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Found Post", post.Title)
	fmt.Fprintf(w, post.Title, post.Subtitle, post.Content)
}

func searchArticleByName(qString string) {
	fmt.Println("inside searchArticleByName")
	//Connect Mongodb
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	//Post Inside Database
	collection := client.Database("inShotsDB").Collection("posts")
	// id, err := primitive.ObjectIDFromHex(qString)
	fmt.Println(qString)

	filterTitle := bson.D{{"title", primitive.Regex{Pattern: qString, Options: ""}}}
	filterSubtitle := bson.D{{"subtitle", primitive.Regex{Pattern: qString, Options: ""}}}
	filterContent := bson.D{{"content", primitive.Regex{Pattern: qString, Options: ""}}}

	InsideTitle, err1 := collection.Find(context.TODO(), filterTitle)
	InsideSubtitle, err2 := collection.Find(context.TODO(), filterSubtitle)
	InsideContent, err3 := collection.Find(context.TODO(), filterContent)
	if err2 != nil || err1 != nil || err3 != nil {
		log.Fatal(err2)
	}

	for InsideTitle.Next(context.TODO()) {
		var post Article
		InsideTitle.Decode(&post)
		Articles = append(Articles, Article{Title: post.Title, Subtitle: post.Subtitle, Content: post.Content})
	}
	for InsideSubtitle.Next(context.TODO()) {
		var post Article
		InsideSubtitle.Decode(&post)
		Articles = append(Articles, Article{Title: post.Title, Subtitle: post.Subtitle, Content: post.Content})
	}
	for InsideContent.Next(context.TODO()) {
		var post Article
		InsideContent.Decode(&post)
		Articles = append(Articles, Article{Title: post.Title, Subtitle: post.Subtitle, Content: post.Content})
	}

	fmt.Println("Found post with title", Articles)
}
