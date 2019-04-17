package oauth

import (
	"github.com/jinzhu/gorm"
	"oauth2-server/config"
)

//Service ...
type Service struct {
	cfg          *config.Config
	db           *gorm.DB
	allowedRoles []string
}

//NewService ...
func NewService(cfg *config.Config, db *gorm.DB) *Service {
	return &Service{
		cfg:          cfg,
		db:           db,
		allowedRoles: []string{},
	}
}

//GetConfig ...
func (s *Service) GetConfig() *config.Config {
	return s.cfg
}

//RestrictToRoles ..
func (s *Service) RestrictToRoles(allowedRoles ...string) {
	s.allowedRoles = allowedRoles
}

//IsRoleAllowed ...
func (s *Service) IsRoleAllowed(role string) bool {
	for _, allowedRole := range s.allowedRoles {
		if role == allowedRole {
			return true
		}
	}
	return false
}

//Close ...
func (s *Service) Close() {

}
