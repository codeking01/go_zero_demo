package logic

import (
	"context"
	"go_zero_demo/gorm_zero/models"

	"go_zero_demo/gorm_zero/internal/svc"
	"go_zero_demo/gorm_zero/types/user"

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

func (l *UserCreateLogic) UserCreate(in *user.UserCreateRequest) (pd *user.UserCreateResponse, err error) {
	pd = new(user.UserCreateResponse)
	var model models.UserModel
	err = l.svcCtx.DB.Take(&model, "username = ?", in.Username).Error
	if err == nil {
		pd.Err = "该用户名已存在!!"
		return
	}
	model = models.UserModel{
		Username: in.Username,
		Password: in.Password,
	}
	err = l.svcCtx.DB.Create(&model).Error
	if err != nil {
		logx.Error(err)
		pd.Err = err.Error()
		err = nil
		return
	}
	pd.UserId = uint32(model.ID)
	return &user.UserCreateResponse{
		UserId: uint32(model.ID),
	}, nil
}
