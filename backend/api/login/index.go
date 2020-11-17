package handler

import (
	"backend/user"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Handler Login Auth MongoDB
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		email := r.FormValue("email")
		pasword := r.FormValue("password")
		logs := login(email, pasword)
		json.NewEncoder(w).Encode(logs)
	} else {
		fmt.Fprintf(w, "Please use post medthod")
	}
}

func login(email, password string) map[string]string {
	u := user.User{}
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
	filter := bson.M{"email": email}
	auths := client.Database("stylelearn").Collection("auths")
	err = auths.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		logs := u.ErrorHandle(err)
		return logs
	}

	// check logic
	logs := u.CheckPassword(password)

	return logs
}
