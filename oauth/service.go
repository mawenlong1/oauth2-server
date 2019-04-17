package oauth

import (
	"github.com/jinzhu/gorm"
	"oauth2-server/config"
)

type Service struct {
	cfg          *config.Config
	db           *gorm.DB
	allowedRoles []string
}

func NewService(cfg *config.Config, db *gorm.DB) *Service {
	return &Service{
		cfg:          cfg,
		db:           db,
		allowedRoles: []string{},
	}
}
