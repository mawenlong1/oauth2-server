package oauth

import (
	"errors"
	"oauth2-server/models"
)

var (
	// ErrInvalidUsernameOrPassword ...
	ErrInvalidUsernameOrPassword = errors.New("Invalid username or password")
)

// Login ..
func (s *Service) Login(client *models.OauthClient, user *models.OauthUser, scope string) (*models.OauthAccessToken, *models.OauthRefreshToken, error) {
	if s.IsRoleAllowed(user.RoleID.String) {
		return nil, nil, ErrInvalidUsernameOrPassword
	}
	accessToken, err := s.GrandAccessToken(client, user, s.cfg.Oauth.AccessTokenLifeTime, scope)
	if err != nil {
		return nil, nil, err
	}
	refreshToken, err := s.GetOrCreateRefreshToken(client, user, s.cfg.Oauth.RefreshTokenLifeTime, scope)
	if err != nil {
		return nil, nil, err
	}
	return accessToken, refreshToken, err
}
