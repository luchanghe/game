package server

import (
	"bytes"
	"encoding/binary"
	"game/pkg/manage/userManage"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
)

var upgraded = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 可以在这里添加更多的验证逻辑，比如检查请求的来源
		return true
	},
}

func Handler(c *gin.Context) {
	conn, err := upgraded.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("连接关闭异常:", err.Error())
			return
		}
	}(conn)
	onOpen(conn)
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			onClose(conn, err)
			return
		}
		onMessage(c, conn, messageType, message)
	}
}

type Data struct {
	Head  [4]uint32
	Proto []byte
}

func onMessage(c *gin.Context, conn *websocket.Conn, messageType int, message []byte) {
	userId := 100
	ok := userManage.Lock(userId)
	if ok {
		//todo 这里后面要变成错误返回
		log.Fatal("请求频繁，请稍后再试")
	}
	defer func() {
		userManage.Unlock(userId)
	}()
	buffer := bytes.NewReader(message)
	var result Data
	for i := 0; i < 4; i++ {
		err := binary.Read(buffer, binary.BigEndian, &result.Head[i])
		if err != nil {
			log.Fatal("二进制数据读取异常:", err)
		}
	}
	u, err := userManage.GetUserFormUid(userId)
	if err != nil {
		log.Fatal("获取用户异常:", err)
		return
	}
	c.Set("user", u)
	remainingBytes := message[16:]
	result.Proto = remainingBytes
	reqRoute := result.Head[1]
	resRoute := result.Head[1] + 1
	res, err := doAction(c, &result, reqRoute)
	if err != nil {
		log.Fatal("异常的返回:", err)
		return
	}
	resData, err := proto.Marshal(res)
	if err != nil {
		log.Fatal("异常的返回:", result.Head[1])
		return
	}
	lenValue := uint32(len(resData) + 13)
	reqId := result.Head[3] // 请求 ID 值
	errorCode := "\x00"     // 错误码
	var buf bytes.Buffer
	err = binary.Write(&buf, binary.BigEndian, lenValue)
	if err != nil {
		log.Fatal("写入返回失败,lenValue:", err)
		return
	}
	err = binary.Write(&buf, binary.BigEndian, resRoute)
	if err != nil {
		log.Fatal("写入返回失败,resRoute:", err)
		return
	}
	err = binary.Write(&buf, binary.BigEndian, reqId)
	if err != nil {
		log.Fatal("写入返回失败,reqId:", err)
		return
	}
	buf.WriteString(errorCode)
	buf.Write(resData)
	packedData := buf.Bytes()
	err = conn.WriteMessage(messageType, packedData)
	if err != nil {
		log.Fatal("消息发送异常:", err.Error())
		return
	}
}

func onClose(conn *websocket.Conn, err error) {
	log.Println("websocket连接关闭:", err.Error())
}

func onOpen(conn *websocket.Conn) {
	log.Println("新的websocket连接")
}
