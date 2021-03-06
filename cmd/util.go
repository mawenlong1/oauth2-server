package cmd

import (
	"oauth2-server/config"
	"oauth2-server/database"

	"github.com/jinzhu/gorm"
)

func initConfigDB(configFile string) (*config.Config, *gorm.DB, error) {
	cfg := config.NewConfig(configFile)

	db, err := database.NewDatabase(cfg)
	if err != nil {
		return nil, nil, err
	}
	return cfg, db, nil
}
