package logic

import (
	"context"
	"errors"
	"go_zero_demo/gorm_zero/models"

	"go_zero_demo/gorm_zero/internal/svc"
	"go_zero_demo/gorm_zero/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (pd *user.UserInfoResponse, err error) {
	var model models.UserModel
	err = l.svcCtx.DB.Take(&model).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	return &user.UserInfoResponse{
		UserId:   uint32(model.ID),
		Username: model.Username,
	}, nil
}
