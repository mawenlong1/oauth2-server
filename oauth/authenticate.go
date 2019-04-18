package oauth

import (
	"errors"
	"github.com/jinzhu/gorm"
	"oauth2-server/models"
	"oauth2-server/session"
	"time"
)

var (
	// ErrAccessTokenNotFound ..
	ErrAccessTokenNotFound = errors.New("Access token not found")
	// ErrAccessTokenExpired ..
	ErrAccessTokenExpired = errors.New("Access token expired")
)

// Authenticate ...
func (s *Service) Authenticate(token string) (*models.OauthAccessToken, error) {
	accessToken := new(models.OauthAccessToken)
	notFound := s.db.Where("token=?", token).First(accessToken).RecordNotFound()
	if notFound {
		return nil, ErrAccessTokenNotFound
	}
	if time.Now().UTC().After(accessToken.ExpiresAt) {
		return nil, ErrRefreshTokenExpired
	}
	// 增加refreshtoken的有效期
	query := s.db.Model(new(models.OauthRefreshToken)).Where("client_id=?", accessToken.ClientID.String)
	if accessToken.UserID.Valid {
		query = query.Where("user_id=?", accessToken.UserID)
	} else {
		query = query.Where("user_id is NULL")
	}
	increasedExpiresAt := gorm.NowFunc().Add(time.Duration(s.cfg.Oauth.RefreshTokenLifeTime) * time.Second)
	if err := query.Update("expires_at", increasedExpiresAt).Error; err != nil {
		return nil, err
	}
	return accessToken, nil
}

// ClearUserTokens ...
func (s *Service) ClearUserTokens(userSession *session.UserSession) {
	refreshToken := new(models.OauthAccessToken)
	found := !models.OauthRefreshTokenPreload(s.db).Where("token=?", userSession.RefreshToken).First(refreshToken).RecordNotFound()
	if found {
		s.db.Unscoped().Where("client_id=? and user_id=?", refreshToken.ClientID, refreshToken.UserID).Delete(models.OauthRefreshToken{})
	}
	accessToken := new(models.OauthAccessToken)
	found = !models.OauthAccessTokenPreload(s.db).Where("token=?", userSession.RefreshToken).First(accessToken).RecordNotFound()
	if found {
		s.db.Unscoped().Where("client_id and user_id=?", accessToken.ClientID, accessToken.UserID).Delete(models.OauthAccessToken{})
	}
}
