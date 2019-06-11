package models

import (
	"database/sql"
	"oauth2-server/util"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/pborman/uuid"
)

// OauthClient 客户端模型
type OauthClient struct {
	MyGormModel
	ClientKey    string         `sql:"type:varchar(254);unique;not null"`
	ClientSecret string         `sql:"type:varchar(60);not null"`
	RedirectURI  sql.NullString `sql:"type:varchar(200)"`
}

// TableName 表名
func (c *OauthClient) TableName() string {
	return "oauth_client"
}

// OauthScope 认证的权限范围
type OauthScope struct {
	MyGormModel
	Scope       string `sql:"type:varchar(200);unique;not null"`
	Description sql.NullString
	IsDefault   bool `sql:"default:false"`
}

// TableName 表名
func (s *OauthScope) TableName() string {
	return "oauth_scope"
}

// OauthRole 角色
type OauthRole struct {
	TimestampModel
	ID   string `gorm:"primary_key" sql:"type:varchar(20)"`
	Name string `sql:"type:varchar(50);unique;not null"`
}

// TableName 表名
func (r *OauthRole) TableName() string {
	return "oauth_role"
}

// OauthUser 认证用户
type OauthUser struct {
	MyGormModel
	RoleID   sql.NullString `sql:"type:varchar(20);index;not null"`
	Role     *OauthRole
	Username string         `sql:"tyep:varchar(254);unique;not null"`
	Password sql.NullString `sql:"type:varchar(60)"`
}

// TableName 表名
func (u *OauthUser) TableName() string {
	return "oauth_user"
}

// OauthRefreshToken 刷新token
type OauthRefreshToken struct {
	MyGormModel
	ClientID  sql.NullString `sql:"index;not null"`
	UserID    sql.NullString `sql:"index"`
	Client    *OauthClient
	User      *OauthUser
	Token     string    `sql:"type:varchar(40);unique;not nulll"`
	ExpiresAt time.Time `sql:"not null;DEFAULT:current_timestamp"`
	Scope     string    `sql:"type:varchar(200);not null"`
}

// TableName 表名
func (rt *OauthRefreshToken) TableName() string {
	return "oauth_refresh_token"
}

// OauthAccessToken 认证token
type OauthAccessToken struct {
	MyGormModel
	ClientID  sql.NullString `sql:"index;not null"`
	UserID    sql.NullString `sql:"index"`
	Client    *OauthClient
	User      *OauthUser
	Token     string    `sql:"type:varchar(40);unique;not nulll"`
	ExpiresAt time.Time `sql:"not null;DEFAULT:current_timestamp"`
	Scope     string    `sql:"type:varchar(200);not null"`
}

// TableName 表名
func (rt *OauthAccessToken) TableName() string {
	return "oauth_access_token"
}

// OauthAuthorizationCode 认证码
type OauthAuthorizationCode struct {
	MyGormModel
	ClientID    sql.NullString `sql:"index;not null"`
	UserID      sql.NullString `sql:"index"`
	Client      *OauthClient
	User        *OauthUser
	Code        string         `sql:"type:varchar(40);unique;not null"`
	RedirectURI sql.NullString `sql:"type:varchar(200)"`
	ExpiresAt   time.Time      `sql:"not null;DEFAULT:current_timestamp"`
	Scope       string         `sql:"type:varchar(200);not null"`
}

// TableName 表名
func (rt *OauthAuthorizationCode) TableName() string {
	return "oauth_authorization_code"
}

// NewOauthRefreshToken 创建一个新的刷新token
func NewOauthRefreshToken(client *OauthClient, user *OauthUser, expiresIn int, scope string) *OauthRefreshToken {
	refreshToken := &OauthRefreshToken{
		MyGormModel: MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		ClientID:  util.StringOrNull(string(client.ID)),
		Token:     uuid.New(),
		ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		Scope:     scope,
	}
	if user != nil {
		refreshToken.UserID = util.StringOrNull(string(user.ID))
	}
	return refreshToken
}

// NewOauthAccessToken 创建一个新的token
func NewOauthAccessToken(client *OauthClient, user *OauthUser, expiresIn int, scope string) *OauthAccessToken {
	accessToken := &OauthAccessToken{
		MyGormModel: MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		ClientID:  util.StringOrNull(string(client.ID)),
		Token:     uuid.New(),
		ExpiresAt: time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		Scope:     scope,
	}
	if user != nil {
		accessToken.UserID = util.StringOrNull(string(user.ID))
	}
	return accessToken
}

// NewOauthAuthorizationCode 创建新的认证码
func NewOauthAuthorizationCode(client *OauthClient, user *OauthUser, expiresIn int, redirectURI, scope string) *OauthAuthorizationCode {
	return &OauthAuthorizationCode{
		MyGormModel: MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		ClientID:    util.StringOrNull(string(client.ID)),
		UserID:      util.StringOrNull(string(user.ID)),
		Code:        uuid.New(),
		ExpiresAt:   time.Now().UTC().Add(time.Duration(expiresIn) * time.Second),
		RedirectURI: util.StringOrNull(redirectURI),
		Scope:       scope,
	}
}

// OauthAuthorizationCodePreload 预加载
func OauthAuthorizationCodePreload(db *gorm.DB) *gorm.DB {
	return PreloadWithPrefix(db, "")
}

// OauthAccessTokenPreload 预加载
func OauthAccessTokenPreload(db *gorm.DB) *gorm.DB {
	return PreloadWithPrefix(db, "")
}

// OauthRefreshTokenPreload 预加载
func OauthRefreshTokenPreload(db *gorm.DB) *gorm.DB {
	return PreloadWithPrefix(db, "")
}

// PreloadWithPrefix 预加载
func PreloadWithPrefix(db *gorm.DB, prefix string) *gorm.DB {
	return db.Preload(prefix + "Client").Preload(prefix + "User")
}
