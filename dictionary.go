package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

type Book struct {
	ID          bson.ObjectId `bson:"_id"`
	Author      []string      `bson:"author"`
	Title       []string      `bson:"title"`
	Url         []string      `bson:"url"`
	IsKindlized bool          `bson:"iskindlized"`
}

func main() {
	var mongodbCredential string = os.Getenv("KINDLIZED_MONGODB")
	session, _ := mgo.Dial(mongodbCredential)
	defer session.Close()

	database := session.DB("kindlized")
	collection := database.C("books")

	b := new(Book)
	count, _ := collection.Count()
	query := collection.Find(bson.M{})
	query.One(&b)

	log.Println(b)
	log.Println(count)
}
