package svc

import "auth/internal/config"

/*
	logic层需要用到的依赖
*/

type ServiceContext struct {
	Config config.Config
	
}