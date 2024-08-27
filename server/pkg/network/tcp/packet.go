package tcp

import "net"

// 客户端数据包格式
type ClientPacket struct {
	Msg 	*Message 
	Conn 	net.Conn
}

// Session数据包格式
type SessionPacket struct {
	Msg  *Message
	Sess *Session
}
