package main

import (
	"server/common/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"` 
	Data int    `json:"uid"`
}

type LoginRequest struct {
	Username string	`json:"username"`
	Password string `json:"password"`
}

// Login
func Login(ctx *gin.Context) {
	var loginRequest LoginRequest
	
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		logger.Logger.Error("parse request body to json failed", zap.Error(err))
	}

	logger.Logger.Info("request body:", zap.String("user:", loginRequest.Username), zap.String("pass:", loginRequest.Password))


}

// Register
func Register(ctx *gin.Context) {
	
}