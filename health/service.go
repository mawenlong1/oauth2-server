package health

import "github.com/jinzhu/gorm"

// Service ..
type Service struct {
	db *gorm.DB
}

// NewService ..
func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

// Close ...
func (s *Service) Close() {

}
