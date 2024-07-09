package server

import (
	"errors"
	"game/action/test"

	"game/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func doAction(c *gin.Context, result *Data, reqRoute uint32) (proto.Message, error) {
	switch reqRoute {

	case uint32(pb.RouteMap_CS_TestController_getContent):
		req := &pb.TestControllerGetContent{}
		res := &pb.TestGetContentResponse{}
		err := proto.Unmarshal(result.Proto, req)
		if err != nil {
			return nil, err
		}
		test.GetContent(c, req, res)
		return res, nil
	case uint32(pb.RouteMap_CS_TestController_getDefaultContent):
		res := &pb.DefaultResponse{}
		test.GetDefaultContent(c, res)
		return res, nil
	default:
		return nil, errors.New("异常的路由枚举")
	}
}
