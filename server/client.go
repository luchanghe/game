package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"server/pb"
	"time"
)

func main() {
	// 先进行登陆
	data := map[string]any{"account": "test", "password": "123456", "sId": 1}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/getToken", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json") // 设置Content-Type头
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取并打印响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var resData struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		fmt.Println("json解析失败", err)
		return
	}

	fmt.Println("获取登陆Token", resData.Data.Token)
	// 服务器地址
	serverAddr := "ws://127.0.0.1:22003" // 替换为你的 WebSocket 服务器地址

	// 连接到 WebSocket 服务器
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatalf("Error connecting to serverManage: %v", err)
	}
	defer conn.Close()

	j := &pb.UserControllerEnter{
		Token: resData.Data.Token,
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
		time.Sleep(1 * time.Second) // 每秒读取一次
	}
}
