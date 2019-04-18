package oauth

import (
	"errors"
	"oauth2-server/models"
	"time"
)

var (
	// ErrAuthorizationCodeNotFound ...
	ErrAuthorizationCodeNotFound = errors.New("Authorization code not found")
	// ErrAuthorizationCodeExpired ...
	ErrAuthorizationCodeExpired = errors.New("Authorization code expired")
	// ErrInvalidRedirectURI ..
	ErrInvalidRedirectURI = errors.New("Invalid redirect uri")
)

// GrantAuthorizationCode ...
func (s *Service) GrantAuthorizationCode(client *models.OauthClient, user *models.OauthUser, expiresIn int, redirectURI, scope string) (*models.OauthAuthorizationCode, error) {
	authorizationCode := models.NewOauthAuthorizationCode(client, user, expiresIn, redirectURI, scope)

	if err := s.db.Create(authorizationCode).Error; err != nil {
		return nil, err
	}
	authorizationCode.Client = client
	authorizationCode.User = user
	return authorizationCode, nil
}

func (s *Service) getValidAuthorizationCode(code, redirectURI string, client *models.OauthClient) (*models.OauthAuthorizationCode, error) {
	authorizationCode := new(models.OauthAuthorizationCode)
	notFound := models.OauthAuthorizationCodePreload(s.db).Where("client_id=?", client.ID).Where("code=?", code).First(authorizationCode).RecordNotFound()

	if notFound {
		return nil, ErrAuthorizationCodeNotFound
	}
	if redirectURI != authorizationCode.RedirectURI.String {
		return nil, ErrInvalidRedirectURI
	}
	if time.Now().After(authorizationCode.ExpiresAt) {
		return nil, ErrAuthorizationCodeExpired
	}
	return authorizationCode, nil
}
