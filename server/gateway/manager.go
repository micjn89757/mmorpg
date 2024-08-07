package main

import (
	"net/http"
	"server/common/logger"

	"github.com/lxzan/gws"
)

/*
网关服务器
*/
type GateWayManager struct {
	wsHandler WebSocket

	// game服务引用
}

func NewGateWayManager() *GateWayManager {
	return &GateWayManager{
		wsHandler: WebSocket{},
	}
}

// TODO: 作为一个客户端连接Game服务器
func (gm *GateWayManager) initAsClient() {

}

// TODO: 作为一个ws服务器为游戏客户端提供服务
func (gm *GateWayManager) initAsServer() {
	logger.InitLogger("gateway")
	upgrader := gws.NewUpgrader(&gm.wsHandler, &gws.ServerOption{
		ParallelEnabled:   true,                                 // 开启并行消息处理
		Recovery:          gws.Recovery,                         // 开启异常恢复
		PermessageDeflate: gws.PermessageDeflate{Enabled: true}, // 开启压缩

		// 在querystring里面传入用户名
		// 把Sec-WebSocket-Key作为连接的key
		// 刷新页面的时候, 会触发上一个连接的OnClose/OnError事件, 这时候需要对比key并删除map里存储的连接
		Authorize: func(r *http.Request, session gws.SessionStorage) bool {
			var name = r.URL.Query().Get("name")
			if name == "" {
				return false
			}
			session.Store("name", name)
			session.Store("websocketKey", r.Header.Get("Sec-WebSocket-Key"))	// 客户端的key
			return true
		},
	})
	
	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := upgrader.Upgrade(writer, request)
		if err != nil {
			return
		}
		go func() {
			socket.ReadLoop() // 此处阻塞会使请求上下文不能顺利被GC
		}()
	})

	http.ListenAndServe(":6666", nil)
}

func (gm *GateWayManager) Run() {
	gm.initAsClient()
	gm.initAsServer()
}

