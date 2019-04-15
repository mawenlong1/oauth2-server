package database

import (
	"fmt"
	"mwl/oauth2-server/config"
	"mwl/oauth2-server/log"
	"time"

	"github.com/jinzhu/gorm"
	// Driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}
}

// NewDatabase 返回gorm.DB
func NewDatabase(cfg *config.Config) (*gorm.DB, error) {

	if cfg.Database.Type == "mysql" {
		args := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.Database.User,
			cfg.Database.Password,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.DatabaseName)
		db, err := gorm.Open(cfg.Database.Type, args)
		if err != nil {
			log.ERROR.Println("数据库连接错误.", err)
			return db, err
		}
		db.DB().SetMaxIdleConns(cfg.Database.MaxIdleConns)
		db.DB().SetMaxOpenConns(cfg.Database.MaxOpenConns)
		db.LogMode(cfg.IsDevelopment)
		return db, nil
	}
	return nil, fmt.Errorf("不支持的数据库%s", cfg.Database.Type)
}
