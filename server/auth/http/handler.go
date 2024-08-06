package main

import (
	"net/http"
	"server/common/db"
	"server/common/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/google/uuid"
)

type Response struct {
	Code  	code    `json:"code"`
	Msg   	string 	`json:"msg"` 
	Data 	any    	`json:"data"`
}

type Request struct {
	Account string	`json:"account"`
	Password string `json:"password"`
}

// Login 登录
func Login(ctx *gin.Context) {
	var err error
	var request Request
	
	if err = ctx.ShouldBindJSON(&request); err != nil {
		logger.Logger.Error("parse request body to json failed", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, Response{
			Code: ParamError,
			Msg: GetCodeMsg(ParamError),
		})
		return
	}

	// 校验
	if len(request.Account) == 0 {
		ctx.JSON(http.StatusBadRequest, Response{Code:ParamError, Msg: "account not null"})
		return
	}

	if len(request.Password) != 32 {
		ctx.JSON(http.StatusBadRequest, Response{Code: ParamError, Msg: "invalid password"})
		return
	}
	user := db.GetUserByAccount(request.Account)
	if user == nil {	// 查询不到用户
		ctx.JSON(http.StatusForbidden, Response{Code: UsernameOrPasswordError, Msg: GetCodeMsg(UsernameOrPasswordError)})
		return
	}

	if user.Password != request.Password {	// 密码错误
		ctx.JSON(http.StatusForbidden, Response{Code: UsernameOrPasswordError, Msg: GetCodeMsg(UsernameOrPasswordError)})
		return
	}

	// 生成一个uuid
	token, err := uuid.NewV7()
	if err != nil {
		logger.Logger.Error("uuid get failed", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, Response{
			Code: TokenGenError,
			Msg: GetCodeMsg(TokenGenError),
		})
		return
	}

	// token 存入redis
	err = db.SetToken(token, request.Account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Code: TokenGenError,
			Msg: GetCodeMsg(TokenGenError),
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Code: LoginSuccess,
		Msg: GetCodeMsg(LoginSuccess),
		Data: token.String(),	// token返回给客户端
	})

}

// Register 注册
func Register(ctx *gin.Context) {
	var err error
	var request Request
	
	if err = ctx.ShouldBindJSON(&request); err != nil {
		logger.Logger.Error("parse request body to json failed", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, Response{
			Code: ParamError,
			Msg: GetCodeMsg(UsernameOrPasswordError),
		})
		return
	}

	// 校验
	if len(request.Account) == 0 {
		ctx.JSON(http.StatusBadRequest, Response{Code:ParamError, Msg: "account not null"})
		return
	}

	if len(request.Password) != 32 {
		ctx.JSON(http.StatusBadRequest, Response{Code: ParamError, Msg: "invalid password, the length of password is 32"})
		return
	}

	user := db.GetUserByAccount(request.Account)
	if user != nil { // 用户已存在
		ctx.JSON(http.StatusForbidden, Response{Code: AccountExist, Msg: GetCodeMsg(AccountExist)})
		return
	}

	// 创建用户
	_, err = db.CreateAccount(request.Account, request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Code: SqlError,
			Msg: GetCodeMsg(SqlError),
		})
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Code: RegistrySuccess,
		Msg: GetCodeMsg(RegistrySuccess),
	})
}