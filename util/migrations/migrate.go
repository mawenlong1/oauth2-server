package migrations

import (
	"fmt"
	"mwl/oauth2-server/log"

	"github.com/jinzhu/gorm"
)

// MigrationStage ..
type MigrationStage struct {
	Name     string
	Function func(db *gorm.DB, name string) error
}

// Migrate ..
func Migrate(db *gorm.DB, migrations []MigrationStage) error {
	for _, m := range migrations {
		if MigrationExists(db, m.Name) {
			continue
		}
		// 创建表
		if err := m.Function(db, m.Name); err != nil {
			return err
		}
		if err := SaveMigration(db, m.Name); err != nil {
			return err
		}
	}
	return nil
}

// MigrationExists ..
func MigrationExists(db *gorm.DB, migrationName string) bool {
	migration := new(Migration)
	found := !db.Where("name=?", migrationName).First(migration).RecordNotFound()
	if found {
		log.INFO.Println("跳过%s", migrationName)
	} else {
		log.INFO.Println("运行%s", migrationName)
	}
	return found
}

// SaveMigration ..
func SaveMigration(db *gorm.DB, migrationName string) error {
	migration := new(Migration)
	migration.Name = migrationName
	if err := db.Create(migration).Error; err != nil {
		log.ERROR.Println("保存记录失败", err)
		return fmt.Errorf("保存记录失败:%s", err)
	}
	return nil
}
