package manage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
	"sync"
	"time"
)

type MongoManage struct {
	Client *mongo.Client
}

var mongoManageOnce sync.Once
var mongoManageCache *MongoManage

func GetMongoManage() *MongoManage {
	mongoManageOnce.Do(func() {
		url := strings.Join([]string{"mongodb://", GetConfigManage().GetString("mongo.host"), ":", GetConfigManage().GetString("mongo.port")}, "")
		option := options.Client().ApplyURI(url)
		option = option.SetMaxPoolSize(GetConfigManage().GetUint64("mongo.max_pool_size"))
		option = option.SetMinPoolSize(GetConfigManage().GetUint64("mongo.min_pool_size"))
		option = option.SetMaxConnIdleTime(time.Minute * time.Duration(GetConfigManage().GetUint64("mongo.max_conn_minute")))
		var err error
		mongoManageCache.Client, err = mongo.Connect(context.TODO(), option)
		if err != nil {
			log.Fatal(err)
		}
	})
	return mongoManageCache
}

func (m *MongoManage) GetUserDb() *mongo.Database {
	return m.Client.Database(GetConfigManage().GetString("mongo.server_db_name"))
}
