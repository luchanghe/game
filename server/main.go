package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/pkg/connect"
	"server/pkg/manage"
)

//go:generate go run ./CreateModelToPb.go
//go:generate go run ./CreatePb.go
//go:generate go run ./CreateAction.go
func main() {
	gin.SetMode(manage.GetConfigManage().GetString("server.mode"))
	r := gin.Default()
	r.GET("/", connect.Handler)
	err := r.Run(":" + manage.GetConfigManage().GetString("server.port"))
	if err != nil {
		fmt.Println("服务启动异常", err.Error())
		return
	}
}
