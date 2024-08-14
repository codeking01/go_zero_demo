package logic

import (
	"context"
	"go_zero_demo/user/rpc/types/user"
	"go_zero_demo/video/api/internal/svc"
	"go_zero_demo/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoLogic {
	return &GetVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoLogic) GetVideo(req *types.VideoReq) (resp *types.VideoRes, err error) {
	//  add your logic here and delete this line
	user1, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: "1",
	})
	logx.Info("req ID: %v", req)

	if err != nil {
		// 记录错误日志
		logx.Errorf("获取用户信息失败: %v", err)
		return nil, err
	}

	if req.Id != user1.Id {
		logx.Info("用户 ID 不匹配")
		return &types.VideoRes{
			Id:   req.Id,
			Name: user1.Name,
		}, nil
	}
	return &types.VideoRes{
		Id:   req.Id,
		Name: user1.Name,
	}, nil
}
