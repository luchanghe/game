package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/model"
	"server/pb"
	"server/pkg/manage/configManage"
	"server/pkg/manage/dataManage"
	"server/pkg/manage/userManage"
)

func Init(c *gin.Context, req *pb.UserControllerInit, res *pb.DefaultResponse) {
	user := model.NewUser()
	user.Id = userManage.GetNextUserId()
	//初始化玩家信息
	user.Name = req.Name
	client := dataManage.GetMongo()
	collection := client.Database(configManage.GetConfig().GetString("mongo.user_db")).Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
	}
}
