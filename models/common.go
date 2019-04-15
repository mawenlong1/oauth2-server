package models

import "time"

// MyGormModel 数据通用属性
type MyGormModel struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// TimestampModel 时间，不带ID
type TimestampModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// EmailTokenModel email
type EmailTokenModel struct {
	MyGormModel
	Reference   string `sql:"type:varchar(40);unique;not null"`
	EmailSent   bool   `sql:"index;not null"`
	EmailSentAt *time.Time
	ExpiresAt   time.Time `sql:"index;not null"`
}
