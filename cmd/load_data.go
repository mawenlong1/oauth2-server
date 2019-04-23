package cmd

import (
	"github.com/RichardKnop/go-fixtures"
	"oauth2-server/log"
)

// LoadData 加载数据
func LoadData(paths []string, configFile string) error {
	log.INFO.Println("加载数据")
	cfg, db, err := initConfigDB(configFile)
	if err != nil {
		return err
	}
	defer db.Close()
	return fixtures.LoadFiles(paths, db.DB(), cfg.Database.Type)
}
