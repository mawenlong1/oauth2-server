package config

// DatabaseConfig 数据库相关配置
type DatabaseConfig struct {
	Type         string `default:"mysql"`
	Host         string `default:"localhost"`
	Port         int    `default:"3306"`
	User         string `default:"root"`
	Password     string `default:"123456"`
	DatabaseName string `default:"oauth2-server"`
	MaxIdleConns int    `default:"5"`
	MaxOpenConns int    `default:"5"`
}

// OauthConfig 服务相关配置
type OauthConfig struct {
	AccessTokenLifeTime  int `default:"3600"`
	RefreshTokenLifeTime int `default:"1209600"`
	AuthCodeLifeTime     int `default:"3600"`
}

// SessionConfig web的session的相关配置
type SessionConfig struct {
	Secret   string `default:"test_secret"`
	Path     string `default:"/"`
	MaxAge   int    `default:"604800"`
	HTTPOnly bool   `default:"True"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host         string `default:"localhost"`
	Port         int    `default:"6379"`
	Password     string `default:""`
	NetWork      string `default:"tcp"`
	MaxIdleConns int    `default:"10"`
}

// Config 总体配置
type Config struct {
	Database      DatabaseConfig
	Oauth         OauthConfig
	Session       SessionConfig
	Redis         RedisConfig
	ServerPort    int  `default:"8080"`
	IsDevelopment bool `default:"True"`
}
