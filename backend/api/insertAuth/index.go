package handler

import (
	"backend/user"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Handler Insert Auth MongoDB
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		authID := r.FormValue("authID")
		userID := r.FormValue("userID")
		email := r.FormValue("email")
		password := r.FormValue("password")
		logs := insert(authID, userID, email, password)
		json.NewEncoder(w).Encode(logs)
	} else {
		fmt.Fprintf(w, "Please use post medthod")
	}
}

func insert(authID, userID, email, password string) map[string]string {
	u := user.User{
		AuthID:   authID,
		UserID:   userID,
		Email:    email,
		Password: password,
	}
	// Connect to mongo
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		logs := u.ErrorHandle(err)
		return logs
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logs := u.ErrorHandle(err)
		return logs
	}
	defer client.Disconnect(ctx)

	// Query data
	auth := client.Database("stylelearn").Collection("auths")
	opts := options.Find()
	opts.SetSort(bson.D{{"authID", -1}})
	sortCursor, err := auth.Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Fatal(err)
	}
	var episodesSorted []bson.M
	if err = sortCursor.All(ctx, &episodesSorted); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodesSorted)

	/*
		// Insert data
		_, err = auth.InsertOne(ctx, u)
		if err != nil {
			logs := u.ErrorHandle(err)
			return logs
		}
	*/

	logs := make(map[string]string)
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = "1" //u.UserID

	return logs
}
