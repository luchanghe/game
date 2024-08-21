package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/pkg/connect"
	"server/pkg/manage/configManage"
)

func main() {
	//启动服务
	r := gin.Default()
	r.GET("/", connect.Handler)
	err := r.Run(":" + configManage.GetConfig().GetString("server.port"))
	if err != nil {
		fmt.Println("服务启动异常", err.Error())
		return
	}
}
