package services

import (
	"oauth2-server/config"
	"oauth2-server/health"
	"oauth2-server/log"
	"reflect"

	"github.com/jinzhu/gorm"
)

// HealthService ..
var HealthService health.ServiceInterface

// UseHealthService ..
func UseHealthService(h health.ServiceInterface) {
	HealthService = h
}

// Init 启动所有服务
func Init(cfg *config.Config, db *gorm.DB) error {
	log.INFO.Println("启动所有服务")
	if nil == reflect.TypeOf(HealthService) {
		HealthService = health.NewService(db)
	}
	return nil
}

// Close 关闭所有服务
func Close() {
	log.INFO.Println("关闭所有服务")
}
