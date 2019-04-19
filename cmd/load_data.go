package cmd

import "github.com/RichardKnop/go-fixtures"

// LoadData 加载数据
func LoadData(paths []string, configFile string) error {
	cfg, db, err := initConfigDB(configFile)
	if err != nil {
		return err
	}
	defer db.Close()
	return fixtures.LoadFiles(paths, db.DB(), cfg.Database.Type)
}
