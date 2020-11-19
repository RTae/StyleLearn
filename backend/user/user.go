package user

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	AuthID   string `bson: "authID" 	json: "authID"`
	UserID   string `bson: "userID" 	json: "userID"`
	Email    string `bson: "name"  		json: "name"`
	Password string `bson: "password" 	json: "password"`
}

func (u *User) checkPassword(passwordCom string) map[string]string {
	logs := make(map[string]string)
	if u.Password == passwordCom {
		logs["status"] = "1"
		logs["msg"] = ""
		logs["result"] = u.UserID
	} else {
		logs["status"] = "215"
		logs["msg"] = "wrong password"
		logs["result"] = ""
	}

	return logs
}

func (u *User) errorHandle(err error) map[string]string {
	var textError string
	var textStatus string
	if err.Error() == "mongo: no documents in result" {
		textError = "User not found"
		textStatus = "215"
	} else {
		textError = err.Error()
		textStatus = "415"
	}

	logs := make(map[string]string)
	logs["status"] = textStatus
	logs["msg"] = textError
	logs["result"] = ""
	return logs
}

// Login user to authication
func (u *User) Login(email, password string) map[string]string {
	// Connect to mongo
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		logs := u.errorHandle(err)
		return logs
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logs := u.errorHandle(err)
		return logs
	}
	defer client.Disconnect(ctx)

	// Query data
	filter := bson.M{"email": email}
	auths := client.Database("stylelearn").Collection("auths")
	err = auths.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		logs := u.errorHandle(err)
		return logs
	}

	// check logic
	logs := u.checkPassword(password)

	return logs
}

// Create new user into platfrom
func (u *User) Create() map[string]string {
	// Connect to mongo
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		logs := u.errorHandle(err)
		return logs
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logs := u.errorHandle(err)
		return logs
	}
	defer client.Disconnect(ctx)

	// Query max data
	auth := client.Database("stylelearn").Collection("auths")
	filter := []bson.M{{
		"$group": bson.M{
			"_id": nil,
			"max": bson.M{"$max": "$authid"},
		}},
	}

	maxCursor, err := auth.Aggregate(ctx, filter)
	if err != nil {
		logs := u.errorHandle(err)
		return logs
	}
	var maxAuthValue []bson.M
	err = maxCursor.All(ctx, &maxAuthValue)
	if err != nil {
		logs := u.errorHandle(err)
		return logs
	}
	authID, _ := maxAuthValue[0]["max"].(string)

	newAuthID, err := increaseID(authID)
	if err != nil {
		logs := u.errorHandle(err)
		return logs
	}

	u.AuthID = newAuthID

	// Insert data
	_, err = auth.InsertOne(ctx, u)
	if err != nil {
		logs := u.errorHandle(err)
		return logs
	}

	logs := make(map[string]string)
	logs["status"] = "1"
	logs["msg"] = ""
	logs["result"] = "1" //u.UserID

	return logs
}

func increaseID(id string) (string, error) {
	digit, err := strconv.Atoi(id[2:])
	if err != nil {
		return "au", err
	}
	digit++
	s := strconv.Itoa(digit)

	newID := "au" + strings.Repeat("0", 8-len(s)) + s
	return newID, nil
}
