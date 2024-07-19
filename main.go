package main

import (
	"fmt"
	"game/pkg/server"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", server.Handler)
	err := r.Run(":22003")
	if err != nil {
		fmt.Println("服务启动异常", err.Error())
		return
	}
}
