package manage

import (
	"bytes"
	"encoding/binary"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"server/pb"
	"sync"
)

type ServerManage struct {
	userToConnMap map[int64]*websocket.Conn
	connToUserMap map[*websocket.Conn]int64
	mu            sync.Mutex
}

var serverManageOnce sync.Once
var serveManageCache *ServerManage

func GetServerManage() *ServerManage {
	serverManageOnce.Do(func() {
		serveManageCache = &ServerManage{
			userToConnMap: make(map[int64]*websocket.Conn),
			connToUserMap: make(map[*websocket.Conn]int64),
		}
	})
	return serveManageCache
}

func (s *ServerManage) BindUserConn(uId int64, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.userToConnMap[uId] = conn
	s.connToUserMap[conn] = uId
}

func (s *ServerManage) DelUserConn(conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if uId, ok := s.connToUserMap[conn]; ok {
		delete(s.userToConnMap, uId)
		delete(s.connToUserMap, conn)
	}
}

func (s *ServerManage) SendUserChangeMessage(change map[int64][]*ChangeCommand) error {
	for otherUid, commands := range change {
		conn, ok := s.userToConnMap[otherUid]
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
		err := s.PbSendToClient(res, conn, uint32(0), uint32(0), uint32(0))
		if err != nil {
			return err
		}
	}
	return nil
}

func (*ServerManage) PbSendToClient(res proto.Message, conn *websocket.Conn, reqId uint32, reqRoute uint32, resRoute uint32) error {
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
