package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"time"
)

type Word struct {
	ID        bson.ObjectId `bson:"_id"`
	Word      string        `bson:word`
	lastQuery time.Time     `bson:lastquery`
}

type Book struct {
	Title []string `bson:"title"`
}

func main() {
	var mongodbCredential string = os.Getenv("KINDLIZED_MONGODB")
	session, _ := mgo.Dial(mongodbCredential)
	defer session.Close()

	database := session.DB("kindlized")
	collection := database.C("books")

	var books []Book

	query := collection.Find(bson.M{})
	query.All(&books)

	log.Println(books)
}
