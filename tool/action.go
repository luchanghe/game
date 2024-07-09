package tool

import (
	"game/been/redis/user"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) *user.User {
	u, _ := c.Get("user")
	return u.(*user.User)
}
