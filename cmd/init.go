package cmd

import (
	"mwl/oauth2-server/models"
	"mwl/oauth2-server/util/migrations"
)

// Init 创建表结构
func Init(configFile string) error {
	_, db, err := initConfigDB(configFile)
	if err != nil {
		return err
	}
	defer db.Close()
	if err := migrations.Bootstrap(db); err != nil {
		return err
	}
	if err := models.MigrateAll(db); err != nil {
		return err
	}
	return nil
}
