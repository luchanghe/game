package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"server/pb"
	"server/pkg/manage"
)

func Enter(c *gin.Context, req *pb.UserControllerEnter, res *pb.UserEnterResponse) {
	token := req.Token
	fmt.Println(token)
	//这里解析通过token获取用户Id，暂时伪装
	uId := int64(10000001)
	conn, _ := c.Get(manage.Conn)
	manage.GetServerManage().BindUserConn(uId, conn.(*websocket.Conn))
}
