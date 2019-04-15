package database_test

import (
	"mwl/oauth2-server/config"
	"mwl/oauth2-server/database"
	"mwl/oauth2-server/log"
	"testing"
)

func TestDatabase(t *testing.T) {
	cfg := config.NewConfig("../config.yml")
	log.INFO.Println(cfg)

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.WARNING.Println("数据库连接错误。", err)
		return
	}
	log.INFO.Println("数据库连接成功。")
	log.INFO.Println(db)
}
