package main

import (
	"encoding/json"
	"server/common/logger"
	"time"

	"github.com/lxzan/gws"
	"go.uber.org/zap"
)

const (
	PingInterval = 5 * time.Second // 客户端心跳间隔
	HeartbeatWaitTimeout  = 10 * time.Second // 心跳等待超时时间
)

type WebSocket struct{
	sessions *gws.ConcurrentMap[string, *gws.Conn] // 使用内置的ConcurrentMap存储连接, 可以减少锁冲突，string是客户端querystring传来的name，这里string存储玩家id
}

func NewWebSocket() *WebSocket {
	return &WebSocket{
		sessions: gws.NewConcurrentMap[string, *gws.Conn](16, 128),
	}
}

// 建立连接事件
func (c *WebSocket) OnOpen(socket *gws.Conn) {
	name := MustLoad[string](socket.Session(), "name")
	if conn, ok := c.sessions.Load(name); ok {
		conn.WriteClose(1000, []byte("connection is replaced"))
	}
	_ = socket.SetDeadline(time.Now().Add(PingInterval + HeartbeatWaitTimeout))
	c.sessions.Store(name, socket)	// 将客户端的name作为key, gws.Conn对象(包含会话)作为value存储起来(存到上述的session map中)
	logger.Logger.Info("connect success", zap.String("client name", name))
}

// 关闭连接事件，
func (c *WebSocket) OnClose(socket *gws.Conn, err error) {
	name := MustLoad[string](socket.Session(), "name")	
	sharding := c.sessions.GetSharding(name)	// c.sessions是业务中的map
	sharding.Lock()
	defer sharding.Unlock()

	if conn, ok := sharding.Load(name); ok {	// 拿到name对应的gws.Conn
		key0 := MustLoad[string](socket.Session(), "websocketKey")	// 当前socketsessions是否还存在map中
		if key1 := MustLoad[string](conn.Session(), "websocketKey"); key1 == key0 {
			sharding.Delete(name)
		}
	}

	logger.Logger.Info("one error", zap.String("name", name), zap.String("msg", err.Error()))
}

// 如果某个场景服务挂了、玩家作弊，也要触发关闭连接
func disConnect(target []string) {
	
}

func (c *WebSocket) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + HeartbeatWaitTimeout))
	_ = socket.WriteString("pong")
}

func (c *WebSocket) OnPong(socket *gws.Conn, payload []byte) {}

type Input struct {
	To		string	`json:"to"`	// 要发送给哪个客户端(name)
	Text	string	`json:"text"`
}

/*
处理客户端传递过来的消息
*/
func (c *WebSocket) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	// chrome websocket不支持ping方法, 所以在text frame里面模拟ping
	if b := message.Bytes(); len(b) == 4 && string(b) == "ping" {	// message转换为bytes，如果是一个ping
		c.OnPing(socket, nil)
		return
	}

	var input = &Input{}
	_ = json.Unmarshal(message.Bytes(), input)
	if conn, ok := c.sessions.Load(input.To); ok {
		_ = conn.WriteMessage(gws.OpcodeText, message.Bytes())
	}
}

// 发送消息给客户端
func (c *WebSocket) sendMessage(target string, )

// 获取给定gws.SessionStorage中存储的内容
func MustLoad[T any](session gws.SessionStorage, key string) (v T) {
	if value, exist := session.Load(key); exist {
		v = value.(T)
	}
	return
}