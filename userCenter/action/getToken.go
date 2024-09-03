package action

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"userCenter/model"
	"userCenter/pkg/manage"
)

type getTokenRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func GetToken(c *gin.Context) {
	var req getTokenRequest
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "账号查询异常"})
		return
	}
	fmt.Println(user)
	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "账号不存在"})
		return
	}
	passwordMd5 := md5.New()
	passwordMd5.Write([]byte(req.Password))
	if user.Password != hex.EncodeToString(passwordMd5.Sum(nil)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "账号或密码错误"})
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix() // 过期时间
	claims["accountId"] = user.Id
	keyStr := manage.GetConfigManage().Viper.GetString("token_secret_key")
	tokenString, err := token.SignedString([]byte(keyStr))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "签名服务异常"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"token": tokenString}})

}
