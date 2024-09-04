package connect

import (
	"bytes"
	"encoding/binary"
	"github.com/gin-gonic/gin"
	"github.com/go-stack/stack"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"net/http"
	"server/define"
	"server/model"
	"server/pb"
	"server/pkg/manage"
	"server/pkg/sysConst"
	"server/tool"
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
		gameError := onMessage(c, conn, message)
		if gameError != nil && gameError.Close {
			onClose(c, conn, gameError)
			return
		}
	}

}

type Data struct {
	Head  [4]uint32
	Proto []byte
}

func onMessage(c *gin.Context, conn *websocket.Conn, message []byte) *tool.GameError {
	buffer := bytes.NewReader(message)
	var result Data
	for i := 0; i < 4; i++ {
		err := binary.Read(buffer, binary.BigEndian, &result.Head[i])
		if err != nil {
			return tool.NewGameError(define.ProtoDecodeFait, false)
		}
	}

	remainingBytes := message[16:]
	result.Proto = remainingBytes
	reqRoute := result.Head[1]
	resRoute := result.Head[1] + 1
	reqId := result.Head[3] // 请求 ID 值
	if reqRoute == uint32(pb.RouteMap_CS_UserController_enter) {
		//登陆不需要获取当前用户
		res, err := doAction(c, &result, reqRoute)
		if err != nil {
			return tool.NewGameError(define.ActionFuncFait, true)
		}
		user, ok := c.Get(sysDefined.ActionUser)
		if !ok {
			return tool.NewGameError(define.ActionFuncFait, true)
		}
		manage.GetServerManage().BindUserConn(user.(*model.User).Id, conn)
		err = manage.GetServerManage().PbSendToClient(res, conn, reqId, reqRoute, resRoute)
		if err != nil {
			return tool.NewGameError(define.SendToClientFait, true)
		}
	} else {
		userId, ok := manage.GetServerManage().ConnToUserMap[conn]
		if !ok {
			return tool.NewGameError(define.UserNoLogin, true)
		}
		u, ok := manage.GetUser(c, userId)
		if !ok {
			return tool.NewGameError(define.RequestOften, true)
		}
		c.Set(sysDefined.ActionUser, u)
		ok = manage.Lock(userId)
		if ok {
			return tool.NewGameError(define.RequestOften, false)
		}
		defer func() {
			manage.Unlock(userId)
		}()
		res, err := doAction(c, &result, reqRoute)
		if err != nil {
			return tool.NewGameError(define.RouteError, false)
		}
		change := manage.GetUserChange(c)
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
		err = manage.GetServerManage().PbSendToClient(res, conn, reqId, reqRoute, resRoute)
		if err != nil {
			return tool.NewGameError(define.SendToClientFait, false)
		}
		if len(change) > 0 {
			err = manage.GetServerManage().SendUserChangeMessage(change)
			if err != nil {
				return tool.NewGameError(define.SendToClientFait, false)
			}
		}
	}
	if gameErr, ok := c.Get(sysDefined.Error); ok {
		return gameErr.(*tool.GameError)
	}
	return nil
}

func onClose(c *gin.Context, conn *websocket.Conn, err error) {
	defer conn.Close()
	manage.GetServerManage().DelUserConn(conn)
	if err != nil {
		log.Println("消息发送异常:", err.Error())
		log.Println(err, stack.Trace().String())
	}
}

func onOpen(c *gin.Context, conn *websocket.Conn) {

}
