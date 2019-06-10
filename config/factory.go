package config

import "github.com/jinzhu/configor"

// DefaultConfig 默认配置
var DefaultConfig = &Config{
	Database: DatabaseConfig{
		Type:         "mysql",
		Host:         "localhost",
		Port:         3306,
		User:         "root",
		Password:     "123456",
		DatabaseName: "oauth2-server",
		MaxIdleConns: 5,
		MaxOpenConns: 5,
	},
	Oauth: OauthConfig{
		AccessTokenLifeTime:  3600,
		RefreshTokenLifeTime: 1209600,
		AuthCodeLifeTime:     3600,
	},
	Session: SessionConfig{
		Secret:   "test_secret",
		Path:     "/",
		MaxAge:   86400 * 7,
		HTTPOnly: true,
	},
	Redis: RedisConfig{
		Host:         "localhost",
		Port:         6379,
		Password:     "",
		MaxIdleConns: 5,
	},
	ServerPort:    8080,
	IsDevelopment: true,
}

// NewDefaultConfig 返回默认配置
func NewDefaultConfig() *Config {
	return DefaultConfig
}

// NewConfig 根据配置文件获取配置信息
func NewConfig(configFile string) *Config {
	if configFile != "" {
		config := &Config{}
		_ = configor.Load(config, configFile)
		return config
	}
	return DefaultConfig
}
