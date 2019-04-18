package oauth

import (
	"errors"
	"oauth2-server/models"
	"time"
)

var (
	// ErrRefreshTokenNotFound ...
	ErrRefreshTokenNotFound = errors.New("Referesh token not found")
	// ErrRefreshTokenExpired ..
	ErrRefreshTokenExpired = errors.New("Refresh token expired")
	// ErrRequestedScopeCannotBeGreater ...
	ErrRequestedScopeCannotBeGreater = errors.New("Requested scope cannot be greater")
)

// GetOrCreateRefreshToken ..
func (s *Service) GetOrCreateRefreshToken(client *models.OauthClient, user *models.OauthUser,
	expiresIn int, scope string) (*models.OauthRefreshToken, error) {
	refreshToken := new(models.OauthRefreshToken)
	query := models.OauthRefreshTokenPreload(s.db).Where("client_id=?", client.ID)
	if user != nil && len([]rune(user.ID)) > 0 {
		query = query.Where("user_id=?", client.ID)
	} else {
		query = query.Where("user_id is NULL")
	}
	found := !query.First(refreshToken).RecordNotFound()
	var expired bool
	if found {
		expired = time.Now().UTC().After(refreshToken.ExpiresAt)
	}
	if expired {
		s.db.Unscoped().Delete(refreshToken)
	}
	if expired || !found {
		refreshToken = models.NewOauthRefreshToken(client, user, expiresIn, scope)
		if err := s.db.Create(refreshToken).Error; err != nil {
			return nil, err
		}
		refreshToken.Client = client
		refreshToken.User = user
	}
	return refreshToken, nil
}

// GetValidRefreshToken ...
func (s *Service) GetValidRefreshToken(token string, client *models.OauthClient) (*models.OauthRefreshToken, error) {
	refreshToken := new(models.OauthRefreshToken)
	notFound := models.OauthRefreshTokenPreload(s.db).Where("client_id=?", client.ID).
		Where("token=?", token).First(refreshToken).RecordNotFound()
	if notFound {
		return nil, ErrRefreshTokenNotFound
	}
	if time.Now().UTC().After(refreshToken.ExpiresAt) {
		return nil, ErrRefreshTokenExpired
	}
	return refreshToken, nil
}
