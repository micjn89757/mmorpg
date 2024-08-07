package main

import (
	"context"
	"net"
	"server/common/db"
	"server/common/idl/gen/auth"
	"server/common/logger"

	"google.golang.org/grpc"
)

/*
grpc鉴权服务
*/

type Auth struct {
	auth.UnimplementedAuthServer
}


func (a *Auth) CheckToken(ctx context.Context, tokenReq *auth.CheckTokenReq) (*auth.CheckTokenRes, error) {
	token := tokenReq.GetToken() // 获取客户端发送的Token

	account, err := db.GetToken(token) // 去redis查询Token对应的用户信息
	errMsg := "valid token error"
	if err != nil {
		logger.Logger.Error("token验证失败")
		return &auth.CheckTokenRes{
			Data: &auth.CheckTokenResData{
				Account: "",
			},
			Error: &errMsg,
		}, err
	}

	return &auth.CheckTokenRes{
		Data: &auth.CheckTokenResData{
			Account: account,	// 返回账号信息
		},
	}, nil
}


// 允许rpc服务
func (a *Auth) Run() {
	authRpc := grpc.NewServer()
	auth.RegisterAuthServer(authRpc, a)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	err = authRpc.Serve(listener)
	
	if err != nil {
		panic(err)
	}
}