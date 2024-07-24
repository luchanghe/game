package main

import (
	"bytes"
	"encoding/binary"
	"game/pb"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

func main() {
	// 服务器地址
	serverAddr := "ws://127.0.0.1:22003/?token=10001" // 替换为你的 WebSocket 服务器地址

	// 连接到 WebSocket 服务器
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()

	j := &pb.TestControllerGetContent{
		Param: 1,
	}
	jx, _ := proto.Marshal(j)
	// 发送消息到服务器
	message := make([]int32, 4)
	message[0] = int32(len(jx)) + 13
	message[1] = 10001
	message[2] = 10002
	message[3] = 1
	// 创建一个 buffer 来存储二进制数据
	buf := new(bytes.Buffer)
	for _, value := range message {
		if err := binary.Write(buf, binary.BigEndian, value); err != nil {
			log.Fatalf("Error writing to buffer: %v", err)
		}
	}
	if err := binary.Write(buf, binary.BigEndian, jx); err != nil {
		log.Fatalf("Error writing to buffer: %v", err)
	}
	err = conn.WriteMessage(websocket.BinaryMessage, buf.Bytes())
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}
	log.Printf("Sent message: %s", message)

	// 读取消息从服务器
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatalf("Error reading message: %v", err)
		}
		log.Printf("Received message: %s", msg)

		// 处理接收到的消息或进行其他操作

		// 你可以根据需要添加停止条件，或者在这里继续读取消息
		// 这里的示例无限循环读取消息
		time.Sleep(1 * time.Second) // 每秒读取一次
	}
}
