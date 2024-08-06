package main

type code uint

const (
	RegistrySuccess code = 400
	LoginSuccess code = 200
	ParamError code  = iota + 10000
	SqlError 
	AccountExist 
	UsernameOrPasswordError
	TokenGenError
)


var codemsg = map[code]string{
	RegistrySuccess: "注册成功",
	LoginSuccess: "登录成功",
	ParamError: "参数异常",
	SqlError: "数据库异常",
	AccountExist: "账号已存在",
	UsernameOrPasswordError: "账号或密码错误",
	TokenGenError: "token 生成异常",
}


func GetCodeMsg(c code) string {
	return codemsg[c]
}