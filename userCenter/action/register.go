package action

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"userCenter/model"
	"userCenter/pkg/manage"
)

type registerRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "数据字段异常"})
		return
	}

	if req.Account == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "账号密码异常"})
		return
	}
	var user model.User
	result := manage.GetMysqlManage().Client.Where("account = ?", req.Account).Find(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该账已存在"})
		return
	}
	passwordMd5 := md5.New()
	passwordMd5.Write([]byte(req.Password))
	user.Account = req.Account
	user.Password = hex.EncodeToString(passwordMd5.Sum(nil))
	result = manage.GetMysqlManage().Client.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "该账已存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "注册成功"})
}
