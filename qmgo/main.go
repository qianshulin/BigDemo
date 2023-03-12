package main

import (
	"context"
	"fmt"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var DataSource = "mongodb://bc:bc142536@81.71.13.173:7017/admin"
var Database = "userAdmin"

type AppUser struct {
	Address    string `json:"address"`
	CreateTime int64  `json:"create_time"`
	LastTime   int64  `json:"last_time"`
}

func main() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: DataSource})
	if err != nil {
		log.Panicln(err, "===================")
	}
	db := client.Database(Database)

	client.Database("userAdmin")
	coll := db.Collection("appuser")
	var m []AppUser
	err = coll.Find(ctx, bson.M{}).All(&m)
	for _, user := range m {
		fmt.Println(user.Address)
	}
}
