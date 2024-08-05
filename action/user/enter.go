package user

import (
	"fmt"
	"game/pb"
	"game/pkg/manage/constManage"
	"game/pkg/manage/serverManage"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Enter(c *gin.Context, req *pb.UserControllerEnter, res *pb.UserEnterResponse) {
	token := req.Token
	fmt.Println(token)
	//这里解析通过token获取用户Id，暂时伪装
	uId := int64(10000001)
	conn, _ := c.Get(constManage.Conn)
	serverManage.BindUserConn(uId, conn.(*websocket.Conn))
}
