package useractionlogic

import (
	"context"

	"AIproto/internal/svc"
	"AIproto/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserCreateLogic) UserCreate(in *__.UserCreateRequest) (*__.UserCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &__.UserCreateResponse{}, nil
}
