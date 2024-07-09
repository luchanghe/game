package main

import (
	"fmt"
	"game/pkg/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/22001", server.Handler)
	err := r.Run(":22001")
	if err != nil {
		fmt.Println("服务启动异常", err.Error())
		return
	}
}
