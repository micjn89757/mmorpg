package main

import (
	"server/common/middleware"
	"server/common/logger"
)

type Auth struct {
	Addr string
}

func NewAuth(addr string) *Auth {
	return &Auth{
		Addr: addr,
	}
}

func (auth *Auth) Run() {
	logger.InitLogger("auth/auth.log") // 初始化logger
	router := initRouter(middleware.Cors(), middleware.GinLogger(logger.Logger))
	if err := router.Run(auth.Addr); err != nil {
		panic("auth server start failed")
	} else {
		logger.Logger.Info("auth server start success!")
	}
}