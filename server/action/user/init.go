package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/model"
	"server/pb"
	"server/pkg/manage"
)

func Init(c *gin.Context, req *pb.UserControllerInit, res *pb.DefaultResponse) {
	user := model.NewUser()
	user.Id = manage.GetNextUserId()
	//初始化玩家信息
	user.Name = req.Name
	collection := manage.GetMongoManage().GetUserDb()
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
	}
}
