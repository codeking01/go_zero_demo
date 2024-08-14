package logic

import (
	"context"
	"fmt"

	"go_zero_demo/apps/login/internal/svc"
	"go_zero_demo/apps/login/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp string, err error) {
	//  add your logic here and delete this line
	fmt.Println(req.UserName, req.Password)
	return "xxxx.xxxx.xxx", nil
}
