package config_test

import (
	"oauth2-server/config"
	"oauth2-server/log"
	"testing"
)

func TestConfig(t *testing.T) {
	log.INFO.Println("获取默认配置")
	log.INFO.Println(config.NewDefaultConfig())
	log.INFO.Println("根据配置文件获取配置")
	log.INFO.Println(config.NewConfig("../config.yml"))
}
