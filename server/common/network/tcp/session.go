package tcp

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

// 多个客户端可以建立多个会话
type Session struct {
	UId				uint64 	// 对session进行标识
	Conn 			net.Conn
	packer 			IPacker
	WriteCh			chan *Message
	IsClose			bool
	IsPlayerOnline	bool
	MessageHandler 	func(packet *SessionPacket)	// 处理Session消息
}	

func NewSession(conn net.Conn) *Session {
	return &Session{Conn: conn, packer: &NormalPacker{binary.BigEndian}, WriteCh: make(chan *Message, 1)}
}

// Run不会阻塞
func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

// Read 读取客户端消息
func (s *Session) Read() {
	for {	
		// 设置读超时时间
		err := s.Conn.SetReadDeadline(time.Now().Add(time.Second))
		if err != nil {
			fmt.Println(err)
			continue
		}
		msg, err := s.packer.Unpack(s.Conn)
		if _, ok := err.(net.Error); ok {
			continue
		}

		fmt.Println("server receive message: ", string(msg.Data))
		s.MessageHandler(&SessionPacket{		// 将msg封装成sessionpacket
			Msg: msg,
			Sess: s,
		})
		
	}
	
}


// Write 从写管道中接收消息，使用send编码发送到客户端
func (s *Session) Write() {
	for {
		select {
		case resp := <- s.WriteCh:
			s.send(resp)
		}
	}
}

// send 将消息编码发送到客户端
func (s *Session) send(message *Message) {
	err := s.Conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
		return 
	}
	bytes, err := s.packer.Pack(message)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = s.Conn.Write(bytes)

	if err != nil {
		fmt.Println(err)
	}
}

// SendMsg 向写管道中写入消息
func (s *Session) SendMsg(msg *Message) {
	s.WriteCh <- msg
}