package migrations

import (
	"fmt"
	"mwl/oauth2-server/log"

	"github.com/jinzhu/gorm"
)

// Bootstrap 创建migration表
func Bootstrap(db *gorm.DB) error {
	migrationName := "bootstrap_migration"
	migration := new(Migration)
	migration.Name = migrationName

	exists := nil == db.Where("name=?", migrationName).First(migration).Error
	if exists {
		log.INFO.Println("表已经存在", migrationName)
		return nil
	}
	// 创建表
	if err := db.CreateTable(new(Migration)).Error; err != nil {
		return fmt.Errorf("创建表migration失败%s", err)
	}
	if err := db.Create(migration).Error; err != nil {
		return fmt.Errorf("表migration插入数据失败%s", err)
	}
	return nil
}
