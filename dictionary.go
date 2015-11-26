package main

import (
	"github.com/ikawaha/kagome/tokenizer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"strings"
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

	query := collection.Find(bson.M{})

	b := new(Book)
	query.One(&b)

	title := b.Title[0]

	t := tokenizer.New()
	tokens := t.Tokenize(title)
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			log.Println("%s\n", token.Surface)
			continue
		}
		features := strings.Join(token.Features(), ",")
		log.Println("%s\t%v\n", token.Surface, features)
	}
	// var books []Book
	// query.All(&books)
	// log.Println(books)
}
