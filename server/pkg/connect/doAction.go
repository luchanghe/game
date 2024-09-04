package connect

import (
	"errors"
	"server/pb"

	"server/action/user"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func doAction(c *gin.Context, result *Data, reqRoute uint32) (proto.Message, error) {
	switch reqRoute {

	case uint32(pb.RouteMap_CS_UserController_enter):

		req := &pb.UserControllerEnter{}
		err := proto.Unmarshal(result.Proto, req)
		if err != nil {
			return nil, err
		}

		res := &pb.UserEnterResponse{}

		user.Enter(c, req, res)

		return res, nil

	default:
		return nil, errors.New("异常的路由枚举")
	}
}
