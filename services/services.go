package services

import (
	"fmt"
	"oauth2-server/config"
	"oauth2-server/health"
	"oauth2-server/log"
	"oauth2-server/oauth"
	"oauth2-server/session"
	"oauth2-server/user"
	"oauth2-server/web"
	"reflect"

	"github.com/jinzhu/gorm"
	"gopkg.in/boj/redistore.v1"
)

var (
	// HealthService ..
	HealthService health.ServiceInterface
	// OauthService ..
	OauthService oauth.ServiceInterface
	// SessionService ...
	SessionService session.ServiceInterface
	// UserService ...
	UserService user.ServiceInterface
	// WebService ..
	WebService web.ServiceInterface
)

// Init 启动所有服务
func Init(cfg *config.Config, db *gorm.DB) error {
	log.INFO.Println("启动所有服务")
	if nil == reflect.TypeOf(HealthService) {
		HealthService = health.NewService(db)
	}
	if nil == reflect.TypeOf(OauthService) {
		OauthService = oauth.NewService(cfg, db)
	}
	if nil == reflect.TypeOf(SessionService) {
		address := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)
		store, err := redistore.NewRediStore(cfg.Redis.MaxIdleConns, cfg.Redis.NetWork, address, "", []byte(cfg.Session.Secret))
		if err != nil {
			panic(err)
		}
		store.SetMaxAge(24 * 3600)
		SessionService = session.NewService(cfg, store)
	}
	if nil == reflect.TypeOf(UserService) {
		UserService = user.NewService(OauthService)
	}
	if nil == reflect.TypeOf(WebService) {
		WebService = web.NewService(cfg, OauthService, SessionService)
	}
	log.INFO.Println("服务启动完成")
	return nil
}

// Close 关闭所有服务
func Close() {
	log.INFO.Println("关闭所有服务")
}
