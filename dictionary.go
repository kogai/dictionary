package main

import (
	"github.com/ikawaha/kagome/tokenizer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"time"
)

type Word struct {
	ID        bson.ObjectId `bson:"_id"`
	Word      string        `bson:word`
	LastQuery time.Time     `bson:lastquery`
}

type Book struct {
	Title []string `bson:"title"`
}

func main() {
	var mongodbCredential string = os.Getenv("KINDLIZED_MONGODB")
	session, _ := mgo.Dial(mongodbCredential)
	defer session.Close()

	database := session.DB("kindlized")

	booksCollection := database.C("books")
	query := booksCollection.Find(bson.M{})

	b := new(Book)
	query.One(&b)

	title := b.Title[0]

	t := tokenizer.New()
	tokens := t.Tokenize(title)
	for _, token := range tokens {
		if token.Class == tokenizer.KNOWN {
			log.Println(token.Surface)
			newWord := &Word{
				ID:        bson.NewObjectId(),
				Word:      token.Surface,
				LastQuery: time.Now(),
			}

			q := database.C("words").Find(bson.M{
				"word": token.Surface,
			})
			count, _ := q.Count()
			if count == 0 {
				database.C("words").Insert(newWord)
			}
		}
	}
	// var books []Book
	// query.All(&books)
	// log.Println(books)
}
