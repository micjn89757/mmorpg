package tcp

import (
	"encoding/binary"
	"io"
	"net"
	"time"
)

type NormalPacker struct {
	ByteOrder binary.ByteOrder
}

// func NewNormalPack(order binary.ByteOrder) *NormalPacker {
// 	return &NormalPacker{
// 		ByteOrder: order,
// 	}
// }

// Pack	将message序列化并发送  
func (p *NormalPacker) Pack(msg *Message) ([]byte, error) {
	// TODO: 一般来说打包数据的时候，都会在前面加上一个魔法数，用来确认消息的有效性
	buffer := make([]byte, 8 + 8 + len(msg.Data))  // |记录总长度 8B|id 8B|data|
	p.ByteOrder.PutUint64(buffer[0:8], uint64(len(buffer)))  
	p.ByteOrder.PutUint64(buffer[8:16], msg.Id)
	copy(buffer[16:], msg.Data)
	return buffer, nil
}

// Unpack 从tcp连接中拿到字节流数据，并进行解析
func (p *NormalPacker) Unpack(reader io.Reader) (*Message, error) {
	// var err error
	// 客户端会发来很多请求，这里的读超时时间限制为稍微＞时长最长的请求值就可以，在不同环境可能时间不同，稍微大一些较好
	err := reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second * 10))

	if err != nil {
		return nil, err
	}

	buffer := make([]byte, 8+8) // 只获取长度和id
	_, err = io.ReadFull(reader, buffer)	// 已经将数据读出
	if err != nil {
		return nil, err
	}

	totalLen := p.ByteOrder.Uint64(buffer[:8])
	id := p.ByteOrder.Uint64(buffer[8:])
	dataSize := totalLen - 16
	dataBuffer := make([]byte, dataSize)	// 单独获取数据
	_, err = io.ReadFull(reader, dataBuffer)
	if err != nil {
		return nil, err
	}

	msg := &Message{
		Id: id,
		Data: dataBuffer,
	}

	return msg, nil 
}