package models

import (
	"fmt"
	"oauth2-server/util/migrations"

	"github.com/jinzhu/gorm"
)

var (
	list = []migrations.MigrationStage{
		{
			Name:     "initial",
			Function: migrate001,
		},
	}
)

// MigrateAll 创建表
func MigrateAll(db *gorm.DB) error {
	return migrations.Migrate(db, list)
}

func migrate001(db *gorm.DB, name string) error {
	if err := db.CreateTable(new(OauthClient)).Error; err != nil {
		return fmt.Errorf("创建表OauthClient失败:%s", err)
	}
	if err := db.CreateTable(new(OauthScope)).Error; err != nil {
		return fmt.Errorf("创建表OauthScope失败:%s", err)
	}
	if err := db.CreateTable(new(OauthRole)).Error; err != nil {
		return fmt.Errorf("创建表OauthRole失败:%s", err)
	}
	if err := db.CreateTable(new(OauthUser)).Error; err != nil {
		return fmt.Errorf("创建表OauthUser失败:%s", err)
	}
	if err := db.CreateTable(new(OauthRefereshToken)).Error; err != nil {
		return fmt.Errorf("创建表OauthRefereshToken失败:%s", err)
	}
	if err := db.CreateTable(new(OauthAccessToken)).Error; err != nil {
		return fmt.Errorf("创建表OauthAccessToken失败:%s", err)
	}
	if err := db.CreateTable(new(OauthAuthorizationCode)).Error; err != nil {
		return fmt.Errorf("创建表OauthAuthorizationCode失败:%s", err)
	}
	err := db.Model(new(OauthUser)).AddForeignKey(
		"role_id", "oauth_role(id)",
		"RESTRICT", "RESTRICT",
	).Error
	if err != nil {
		return fmt.Errorf("创建外键失败: %s", err)
	}
	err = db.Model(new(OauthRefereshToken)).AddForeignKey(
		"client_id", "oauth_client(id)",
		"RESTRICT", "RESTRICT",
	).Error
	if err != nil {
		return fmt.Errorf("创建外键失败:%s", err)
	}
	err = db.Model(new(OauthRefereshToken)).AddForeignKey(
		"user_id", "oauth_user(id)",
		"RESTRICT", "RESTRICT",
	).Error
	if err != nil {
		return fmt.Errorf("创建外键失败:%s", err)
	}
	err = db.Model(new(OauthAccessToken)).AddForeignKey(
		"client_id", "oauth_client(id)",
		"RESTRICT", "RESTRICT",
	).Error
	if err != nil {
		return fmt.Errorf("创建外键失败:%s", err)
	}
	err = db.Model(new(OauthAccessToken)).AddForeignKey(
		"user_id", "oauth_user(id)",
		"RESTRICT", "RESTRICT",
	).Error
	if err != nil {
		return fmt.Errorf("创建外键失败:%s", err)
	}
	err = db.Model(new(OauthAuthorizationCode)).AddForeignKey(
		"client_id", "oauth_client(id)",
		"RESTRICT", "RESTRICT",
	).Error
	if err != nil {
		return fmt.Errorf("创建外键失败:%s", err)
	}
	err = db.Model(new(OauthAuthorizationCode)).AddForeignKey(
		"user_id", "oauth_user(id)",
		"RESTRICT", "RESTRICT",
	).Error
	if err != nil {
		return fmt.Errorf("创建外键失败:%s", err)
	}
	return nil
}
