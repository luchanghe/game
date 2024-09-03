package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"userCenter/action"
	"userCenter/pkg/manage"
)

func main() {
	//连接初始化
	manage.GetMysqlManage()

	gin.SetMode(manage.GetConfigManage().GetString("userCenter.mode"))
	r := gin.Default()
	r.POST("/register", action.Register)
	r.POST("/getToken", action.GetToken)
	r.POST("/getServerList", action.GetServerList)
	err := r.Run(":" + manage.GetConfigManage().GetString("userCenter.port"))
	if err != nil {
		fmt.Println("服务启动异常", err.Error())
		return
	}
}
