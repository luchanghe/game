package serverManage

import (
	"bytes"
	"encoding/binary"
	"game/pb"
	"game/pkg/manage/userManage"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"sync"
)

var userToConnMap = make(map[int64]*websocket.Conn)
var connToUserMap = make(map[*websocket.Conn]int64)
var mu sync.Mutex

func BindUserConn(uId int64, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	userToConnMap[uId] = conn
	connToUserMap[conn] = uId
}

func DelUserConn(conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	if uId, ok := connToUserMap[conn]; ok {
		delete(userToConnMap, uId)
		delete(connToUserMap, conn)
	}
}

func SendUserChangeMessage(change map[int64][]*userManage.ChangeCommand) error {
	for otherUid, commands := range change {
		conn, ok := userToConnMap[otherUid]
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
		err := PbSendToClient(res, conn, uint32(0), uint32(0), uint32(0))
		if err != nil {
			return err
		}
	}
	return nil
}

func PbSendToClient(res proto.Message, conn *websocket.Conn, reqId uint32, reqRoute uint32, resRoute uint32) error {
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
