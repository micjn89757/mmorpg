package logic

import (
	"context"
	"mmorpg/server/pkg/logger"
	"mmorpg/server/service/auth/rpc/internal/svc"

	"go.uber.org/zap"
)

type CheckTokenLogic struct {
	ctx context.Context 
	svcCtx *svc.ServiceContext
	logger zap.Logger
}

func NewCheckTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckTokenLogic {
	return &CheckTokenLogic{
		ctx: ctx,
		svcCtx: svcCtx,
		logger: logger.GetLogger(),
	}
}