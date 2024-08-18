package common

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// MySQLConfig 用于映射 YAML 文件中的 MySQL 配置
type MySQLConfig struct {
	DataSource string `yaml:"DataSource"`
}

// Config 用于映射整个配置文件
type Config struct {
	MySQL MySQLConfig `yaml:"MySQL"`
}

func InitGorm(configPath string) *gorm.DB {
	// 使用解析后的数据库连接字符串初始化 GORM
	db, err := gorm.Open(mysql.Open(configPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	return db
}
