package edcrypt

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"server/common/logger"
)

/*
md5加密
*/

func Md5(text string) string {
	md5 := md5.New()
	md5.Write([]byte(text))
	digest := md5.Sum(nil)            //md5哈希的结果是128bit
	return hex.EncodeToString(digest) //十六进制编码之后是128/4=32个字符
}



/*
获取uuid

UUID 的字符串表示形式由 32 个十六进制数字组成，以 5 个组显示，由连字符 - 分隔。例如：
123e4567-e89b-12d3-a456-426655440000
*/

func UUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		logger.Logger.Error(err.Error())
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}