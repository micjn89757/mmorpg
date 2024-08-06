package db

import (
	"context"
	"server/common/logger"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

/*
Redis 存入token信息
*/


const (
	TOKEN_PREFIX = "token_"
)

// 把token写入redis
func SetToken(token uuid.UUID, account string) error {
	client := GetRedisClient()
	ctx := context.Background()
	if err := client.Set(ctx, TOKEN_PREFIX+token.String(), account, time.Hour * 0).Err(); err != nil { // 设置不过期
		logger.Logger.Error("write token to redis failed", zap.String("token", token.String()), zap.String("account", account), zap.Error(err))
		return err
	}
	return nil
}


// 获取token对应的用户信息
func GetToken(token string) (string, error) {
	client := GetRedisClient()
	ctx := context.Background()
	account, err := client.Get(ctx, TOKEN_PREFIX+token).Result()

	if err != nil {
		logger.Logger.Error("get account from redis failed", zap.String("account", account), zap.Error(err))
		return "", err
	}

	return account, nil
}


// 删除用户信息
func DelToken(token string) error {
	client := GetRedisClient()
	ctx := context.Background()

	_, err := client.Del(ctx, TOKEN_PREFIX+token).Result()
	if err != nil {
		logger.Logger.Error("del token failed", zap.Error(err))
		return err 
	}

	return nil
}