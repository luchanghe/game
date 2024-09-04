package tool

import (
	"github.com/gin-gonic/gin"
	"server/define"
	"server/pkg/sysConst"
)

type GameError struct {
	Code    int
	Close   bool
	Message string
}

func (g GameError) Error() string {
	return g.Message
}

func NewGameError(code int, close bool) *GameError {
	return &GameError{
		code, close, define.ErrorMap[code],
	}
}
func SetGameError(c *gin.Context, code int) {
	c.Set(sysDefined.Error, NewGameError(
		code, false,
	))
}

func SetGameErrorAndCloseConn(c *gin.Context, code int) {
	c.Set(sysDefined.Error, NewGameError(
		code, true,
	))
}
