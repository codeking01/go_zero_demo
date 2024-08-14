package logic

import (
	"context"
	"go_zero_demo/user/rpc/internal/svc"
	"go_zero_demo/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	//  add your logic here and delete this line
	logx.Debugf("Received ID: %v", in.Id)

	// 如果传入的 ID 是 "1"
	if in.Id == "1" {
		// 返回固定的信息
		return &user.UserResponse{
			Id:     "1",
			Name:   "湖南省，衡阳市",
			Gender: true,
		}, nil
	}

	// 如果 ID 不是 "1"，返回一个默认值或错误信息
	return &user.UserResponse{
		Id:     in.Id,
		Name:   "未知用户",
		Gender: false,
	}, nil
}
