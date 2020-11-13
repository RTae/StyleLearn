package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type env struct {
	Foo string `json: foo`
	Bar string `json: bar`
}

type User struct {
	Id       string `bson: "_id"`
	Email    string `bson: "email"`
	Name     string `bson: "name"`
	Password string `bson: "passwrod"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	/*
		f := os.Getenv("FOO")
		b := os.Getenv("BAR")

			e := env{
				Foo: f,
				Bar: b,
			}
	*/
	a := connetMongo()
	json.NewEncoder(w).Encode(a)

}

func connetMongo() []bson.M {
	u := User{}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://rate:123qwe@cluster0.sbzww.mongodb.net/sample_mflix?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	sampleMflix := client.Database("sample_mflix")
	usersCollection := sampleMflix.Collection("users")

	cursor, err := usersCollection.Find(ctx, bson.M{}).Decode(&u)
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	return episodes
}
