package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"game/pb"
	"game/pkg/manage/userManage"
	"github.com/gin-gonic/gin"
	"github.com/go-stack/stack"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var upgraded = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 可以在这里添加更多的验证逻辑，比如检查请求的来源
		return true
	},
}
var connections = make(map[int]*websocket.Conn)
var mu sync.Mutex

func Handler(c *gin.Context) {
	var err error
	conn, err := upgraded.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	token := c.Query("token")
	if token == "" {
		return
	}
	//假装一下token就是用户ID 以后改为鉴权
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("连接关闭异常:", err.Error())
			return
		}
	}(conn)
	uId, _ := strconv.Atoi(token)
	c.Set(userManage.ConnUid, uId)
	onOpen(c, conn)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			onClose(c, conn, err)
			return
		}
		err = onMessage(c, conn, message)
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
	userId := 100
	ok := userManage.Lock(userId)
	if ok {
		//todo 这里后面要变成错误返回
		log.Println("请求频繁，请稍后再试")
	}
	defer func() {
		userManage.Unlock(userId)
	}()
	buffer := bytes.NewReader(message)
	var result Data
	for i := 0; i < 4; i++ {
		err := binary.Read(buffer, binary.BigEndian, &result.Head[i])
		if err != nil {
			return err
		}
	}
	u := userManage.GetUser(c, userId)
	c.Set(userManage.ActionUser, u)
	remainingBytes := message[16:]
	result.Proto = remainingBytes
	reqRoute := result.Head[1]
	resRoute := result.Head[1] + 1
	reqId := result.Head[3] // 请求 ID 值
	fmt.Println("reqRoute", reqRoute, "resRoute", resRoute, "reqId", reqId, "proto", result.Proto)
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
	err = pbSendToClient(res, conn, reqId, reqRoute, resRoute)
	if err != nil {
		return err
	}
	if len(change) > 0 {
		err = sendUserChangeMessage(change)
		if err != nil {
			return err
		}
	}
	return nil
}

func pbSendToClient(res proto.Message, conn *websocket.Conn, reqId uint32, reqRoute uint32, resRoute uint32) error {
	resData, err := proto.Marshal(res)
	if err != nil {
		return err
	}
	lenValue := uint32(len(resData) + 13)
	errorCode := "\x00" // 错误码
	var buf bytes.Buffer
	err = binary.Write(&buf, binary.BigEndian, lenValue)
	if err != nil {
		return err
	}
	err = binary.Write(&buf, binary.BigEndian, resRoute)
	if err != nil {
		return err
	}
	err = binary.Write(&buf, binary.BigEndian, reqId)
	if err != nil {
		return err
	}
	buf.WriteString(errorCode)
	buf.Write(resData)
	packedData := buf.Bytes()
	err = conn.WriteMessage(websocket.BinaryMessage, packedData)
	if err != nil {
		return err
	}
	return nil
}

func sendUserChangeMessage(change map[int][]*userManage.ChangeCommand) error {
	for otherUid, commands := range change {
		conn, ok := connections[otherUid]
		if !ok {
			continue
		}
		changeMessage := &pb.ChangeMessage{
			ChangeCommand: []*pb.ChangeMessage_Command{},
		}
		for _, data := range commands {
			changeMessage.ChangeCommand = append(changeMessage.ChangeCommand, &pb.ChangeMessage_Command{
				Object:       data.Object,
				Operate:      data.Operate,
				OperateValue: data.OperateValue,
			})
		}
		res := &pb.DefaultResponse{C: changeMessage}
		err := pbSendToClient(res, conn, uint32(0), uint32(0), uint32(0))
		if err != nil {
			return err
		}
	}
	return nil
}

func onClose(c *gin.Context, conn *websocket.Conn, err error) {
	mu.Lock()
	defer mu.Unlock()
	uId, ok := c.Get(userManage.ConnUid)
	if !ok {
		panic("用户ID不存在,不应该走到这里")
		return
	}
	delete(connections, uId.(int))
	if err != nil {
		log.Println(err, stack.Trace().String())
	}

}

func onOpen(c *gin.Context, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	uId, ok := c.Get(userManage.ConnUid)
	if !ok {
		panic("用户ID不存在,不应该走到这里")
		return
	}
	userId := uId.(int)
	if _, ok := connections[userId]; ok {
		//要T掉原先的人，目前不做处理
	}
	connections[userId] = conn
}
