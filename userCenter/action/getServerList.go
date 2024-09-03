package action

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"userCenter/model"
	"userCenter/pkg/manage"
)

func GetServerList(c *gin.Context) {
	var list []model.ServerList
	result := manage.GetMysqlManage().Client.Find(&list)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "区服列表请求异常"})
		fmt.Println(result.Error)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"data": list})
}
