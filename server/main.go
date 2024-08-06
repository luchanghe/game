package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/pkg/connect"
)

func main() {
	r := gin.Default()
	r.GET("/", connect.Handler)
	err := r.Run(":22003")
	if err != nil {
		fmt.Println("服务启动异常", err.Error())
		return
	}
}
