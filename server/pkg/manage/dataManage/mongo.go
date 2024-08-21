package dataManage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"server/pkg/manage/configManage"
	"strings"
)

func GetMongo() *mongo.Client {
	url := strings.Join([]string{"mongodb://", configManage.GetConfig().GetString("mongo.host"), ":", configManage.GetConfig().GetString("mongo.port")}, "")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	return client
}
