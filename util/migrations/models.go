package migrations

import (
	"github.com/jinzhu/gorm"
)

// Migration ..
type Migration struct {
	gorm.Model
	Name string `sql:"size:255"`
}

// TableName 表名
func (c *Migration) TableName() string {
	return "migration"
}
