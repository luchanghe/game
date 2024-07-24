package test

import (
	"fmt"
	"game/pb"
	"github.com/gin-gonic/gin"
)

func GetContent(c *gin.Context, req *pb.TestControllerGetContent, res *pb.TestGetContentResponse) {
	fmt.Println("进入成功，获取param值为", req.Param)
}
