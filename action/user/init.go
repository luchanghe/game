package user

import (
	"context"
	"game/model"
	"game/pb"
	"game/pkg/manage/dataManage"
	"game/pkg/manage/userManage"
	"github.com/gin-gonic/gin"
	"log"
)

func Init(c *gin.Context, req *pb.UserControllerInit, res *pb.DefaultResponse) {
	user := model.NewUser()
	user.Id = userManage.GetNextUserId()
	//初始化玩家信息
	user.Name = req.Name
	client := dataManage.GetMongo()
	collection := client.Database("server_1").Collection("user")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
}
