package connect

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-stack/stack"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"net/http"
	"server/pb"
	"server/pkg/manage/constManage"
	"server/pkg/manage/serverManage"
	"server/pkg/manage/userManage"
)

var upgraded = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 可以在这里添加更多的验证逻辑，比如检查请求的来源
		return true
	},
}

func Handler(c *gin.Context) {
	var err error
	conn, err := upgraded.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	onOpen(c, conn)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			onClose(c, conn, err)
			return
		}
		err = onMessage(c, conn, message)
		fmt.Println(123)
		if err != nil {
			onClose(c, conn, err)
			return
		}
	}

}

type Data struct {
	Head  [4]uint32
	Proto []byte
}

func onMessage(c *gin.Context, conn *websocket.Conn, message []byte) error {
	buffer := bytes.NewReader(message)
	var result Data
	for i := 0; i < 4; i++ {
		err := binary.Read(buffer, binary.BigEndian, &result.Head[i])
		if err != nil {
			return err
		}
	}

	remainingBytes := message[16:]
	result.Proto = remainingBytes
	reqRoute := result.Head[1]
	resRoute := result.Head[1] + 1
	reqId := result.Head[3] // 请求 ID 值
	fmt.Println("reqRoute", reqRoute, "resRoute", resRoute, "reqId", reqId, "proto", result.Proto)
	if reqRoute == uint32(pb.RouteMap_CS_UserController_init) || reqRoute == uint32(pb.RouteMap_CS_UserController_init) {
		//登陆和注册不需要获取当前用户
		res, err := doAction(c, &result, reqRoute)
		if err != nil {
			return err
		}
		err = serverManage.PbSendToClient(res, conn, reqId, reqRoute, resRoute)
		if err != nil {
			return err
		}
	} else {
		userId := int64(10)
		u := userManage.GetUser(c, userId)
		c.Set(constManage.ActionUser, u)
		ok := userManage.Lock(userId)
		if ok {
			//todo 这里后面要变成错误返回
			log.Println("请求频繁，请稍后再试")
		}
		defer func() {
			userManage.Unlock(userId)
		}()
		res, err := doAction(c, &result, reqRoute)
		if err != nil {
			return err
		}
		change := userManage.GetUserChange(c)
		changeData, ok := change[userId]
		if ok {
			fieldDescriptor := res.ProtoReflect().Descriptor().Fields().ByName("c")
			changeMessage := &pb.ChangeMessage{
				ChangeCommand: []*pb.ChangeMessage_Command{},
			}
			for _, diff := range changeData {
				changeMessage.ChangeCommand = append(changeMessage.ChangeCommand, &pb.ChangeMessage_Command{Object: diff.Object, Operate: diff.Operate, OperateValue: diff.OperateValue})
			}
			changeMessageValue := protoreflect.ValueOf(changeMessage.ProtoReflect())
			res.ProtoReflect().Set(fieldDescriptor, changeMessageValue)
			delete(change, userId)
		}
		err = serverManage.PbSendToClient(res, conn, reqId, reqRoute, resRoute)
		if err != nil {
			return err
		}
		if len(change) > 0 {
			err = serverManage.SendUserChangeMessage(change)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func onClose(c *gin.Context, conn *websocket.Conn, err error) {
	serverManage.DelUserConn(conn)
	if err != nil {
		log.Println("消息发送异常:", err.Error())
		log.Println(err, stack.Trace().String())
	}
}

func onOpen(c *gin.Context, conn *websocket.Conn) {

}
