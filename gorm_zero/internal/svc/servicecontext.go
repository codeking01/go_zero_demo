package svc

import (
	"go_zero_demo/gorm_zero/common"
	"go_zero_demo/gorm_zero/internal/config"
	"go_zero_demo/gorm_zero/models"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化 GORM，传入配置文件路径
	db := common.InitGorm(c.Mysql.DataSource)
	// 新建数据库
	db.AutoMigrate(&models.UserModel{})

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
