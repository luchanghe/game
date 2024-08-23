package manage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
	"sync"
)

type MongoManage struct{}

var mongoManageOnce sync.Once
var mongoManageCache *MongoManage

func GetMongoManage() *MongoManage {
	mongoManageOnce.Do(func() {
		mongoManageCache = &MongoManage{}
	})
	return mongoManageCache
}
func (m *MongoManage) GetMongo() *mongo.Client {
	url := strings.Join([]string{"mongodb://", GetConfigManage().GetString("mongo.host"), ":", GetConfigManage().GetString("mongo.port")}, "")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (m *MongoManage) GetUserDb() *mongo.Collection {
	client := m.GetMongo()
	collection := client.Database(GetConfigManage().GetString("mongo.user_db")).Collection("users")
	return collection
}
