package oauth

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pborman/uuid"
	"oauth2-server/models"
	"oauth2-server/util"
	pass "oauth2-server/util/password"
	"strings"
	"time"
)

var (
	// MinPasswordLength ..
	MinPasswordLength = 6
	// ErrPasswordTooShort ..
	ErrPasswordTooShort = fmt.Errorf(
		"oauth:密码至少%d个字符",
		MinPasswordLength)
	// ErrUserNotFound ...
	ErrUserNotFound = errors.New("oauth:User not found")
	// ErrInvalidUserPassword ...
	ErrInvalidUserPassword = errors.New("oauth:Invaild user password")
	// ErrCannotSetEmptyUsername ...
	ErrCannotSetEmptyUsername = errors.New("oauth:Cannot set empty username")
	// ErrUserPasswordNotSet ...
	ErrUserPasswordNotSet = errors.New("oauth:User password not set")
	// ErrUsernameTaken ...
	ErrUsernameTaken = errors.New("oauth:Username taken")
)

// UserExists ...
func (s *Service) UserExists(username string) bool {
	_, err := s.FindUserByUserName(username)
	return err == nil
}

// FindUserByUserName ...
func (s *Service) FindUserByUserName(username string) (*models.OauthUser, error) {
	user := new(models.OauthUser)
	notFound := s.db.Where("username=LOWER(?)", username).First(user).RecordNotFound()
	if notFound {
		return nil, ErrUserNotFound
	}
	return user, nil
}

// CreateUser ...
func (s *Service) CreateUser(roleID, username, password string) (*models.OauthUser, error) {
	return s.createUserCommon(s.db, roleID, username, password)
}

// CreateUserTx ..
func (s *Service) CreateUserTx(tx *gorm.DB, roleID, username, password string) (*models.OauthUser, error) {
	return s.createUserCommon(tx, roleID, username, password)
}
func (s *Service) createUserCommon(db *gorm.DB, roleID, username, password string) (*models.OauthUser, error) {
	user := &models.OauthUser{
		MyGormModel: models.MyGormModel{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
		},
		RoleID:   util.StringOrNull(roleID),
		Username: strings.ToLower(username),
		Password: util.StringOrNull(""),
	}
	if password != "" {
		if len(password) < MinPasswordLength {
			return nil, ErrPasswordTooShort
		}
		passwordHash, err := pass.HashPassword(password)
		if err != nil {
			return nil, err
		}
		user.Password = util.StringOrNull(string(passwordHash))
	}
	if s.UserExists(user.Username) {
		return nil, ErrUsernameTaken
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// SetPassword ...
func (s *Service) SetPassword(user *models.OauthUser, password string) error {
	return s.setPasswordCommon(s.db, user, password)
}

// SetPasswordTx ...
func (s *Service) SetPasswordTx(tx *gorm.DB, user *models.OauthUser, password string) error {
	return s.setPasswordCommon(tx, user, password)
}
func (s *Service) setPasswordCommon(db *gorm.DB, user *models.OauthUser, password string) error {
	if len(password) < MinPasswordLength {
		return ErrPasswordTooShort
	}

	passwordHash, err := pass.HashPassword(password)
	if err != nil {
		return err
	}

	return db.Model(user).UpdateColumn(models.OauthUser{
		Password:    util.StringOrNull(string(passwordHash)),
		MyGormModel: models.MyGormModel{UpdatedAt: time.Now().UTC()},
	}).Error
}

// UpdateUsername ...
func (s *Service) UpdateUsername(user *models.OauthUser, username string) error {
	return s.updateUsernameCommon(s.db, user, username)
}

// UpdateUsernameTx ...
func (s *Service) UpdateUsernameTx(tx *gorm.DB, user *models.OauthUser, username string) error {
	return s.updateUsernameCommon(tx, user, username)
}
func (s *Service) updateUsernameCommon(db *gorm.DB, user *models.OauthUser, username string) error {
	if username == "" {
		return ErrCannotSetEmptyUsername
	}
	return db.Model(user).UpdateColumn("username", strings.ToLower(username)).Error
}

// AuthUser ...
func (s *Service) AuthUser(username, password string) (*models.OauthUser, error) {
	user, err := s.FindUserByUserName(username)
	if err != nil {
		return nil, err
	}
	if !user.Password.Valid {
		return nil, ErrUserPasswordNotSet
	}
	if pass.VerifyPassword(user.Password.String, password) != nil {
		return nil, ErrInvalidUserPassword
	}
	return user, nil
}
