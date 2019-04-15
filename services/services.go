package services

import (
	"oauth2-server/config"
	"oauth2-server/log"

	"github.com/jinzhu/gorm"
)

// Init 启动所有服务
func Init(cfg *config.Config, db *gorm.DB) error {
	log.INFO.Println("启动所有服务")
	return nil
}

// Init 关闭所有服务
func Close() {
	log.INFO.Println("关闭所有服务")
}
