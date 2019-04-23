package oauth

import (
	"errors"
	"net/http"
	"oauth2-server/models"
)

const (
	// AccessTokenHint ..
	AccessTokenHint = "access_token"
	// RefreshTokenHint ..
	RefreshTokenHint = "refresh_token"
)

var (
	// ErrTokenMissing ...
	ErrTokenMissing = errors.New("introspect:Token missing")
	// ErrTokenHintInvalid ...
	ErrTokenHintInvalid = errors.New("introspect:Invalid token hint")
)

// NewIntrospectResponseFromAccessToken ..
func (s *Service) NewIntrospectResponseFromAccessToken(accessToken *models.OauthAccessToken) (*IntrospectResponse, error) {
	var introspectResponse = &IntrospectResponse{
		Active:    true,
		Scope:     accessToken.Scope,
		TokenType: TokenTypes,
		ExpiresAt: int(accessToken.ExpiresAt.Unix()),
	}
	if accessToken.ClientID.Valid {
		client := new(models.OauthClient)
		notFound := s.db.Select("client_key").First(client, accessToken.ClientID.String).RecordNotFound()
		if notFound {
			return nil, ErrClientNotFound
		}
		introspectResponse.ClientID = client.ClientKey
	}

	if accessToken.UserID.Valid {
		user := new(models.OauthUser)
		notFound := s.db.Select("username").Where("id=?", accessToken.UserID.String).First(user).RecordNotFound()
		if notFound {
			return nil, ErrUserNotFound
		}
		introspectResponse.Username = user.Username
	}
	return introspectResponse, nil
}

// NewIntrospectResponseFromRefreshToken ..
func (s *Service) NewIntrospectResponseFromRefreshToken(refreshToken *models.OauthRefreshToken) (*IntrospectResponse, error) {
	var introspectResponse = &IntrospectResponse{
		Active:    true,
		Scope:     refreshToken.Scope,
		TokenType: TokenTypes,
		ExpiresAt: int(refreshToken.ExpiresAt.Unix()),
	}
	if refreshToken.ClientID.Valid {
		client := new(models.OauthClient)
		notFound := s.db.Select("client_key").First(client, refreshToken.ClientID.String).RecordNotFound()
		if notFound {
			return nil, ErrClientNotFound
		}
		introspectResponse.ClientID = client.ClientKey
	}
	if refreshToken.UserID.Valid {
		user := new(models.OauthUser)
		notFound := s.db.Select("username").Where("id=?", refreshToken.UserID.String).First(user).RecordNotFound()
		if notFound {
			return nil, ErrUserNotFound
		}
		introspectResponse.Username = user.Username
	}
	return introspectResponse, nil
}

func (s *Service) introspectToken(r *http.Request, client *models.OauthClient) (*IntrospectResponse, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	token := r.Form.Get("token")
	if token == "" {
		return nil, ErrTokenMissing
	}
	tokenTypeHint := r.Form.Get("token_type_hint")
	if tokenTypeHint == "" {
		tokenTypeHint = AccessTokenHint
	}
	switch tokenTypeHint {
	case AccessTokenHint:
		accessToken, err := s.Authenticate(token)
		if err != nil {
			return nil, err
		}
		return s.NewIntrospectResponseFromAccessToken(accessToken)
	case RefreshTokenHint:
		refreshToken, err := s.GetValidRefreshToken(token, client)
		if err != nil {
			return nil, err
		}
		return s.NewIntrospectResponseFromRefreshToken(refreshToken)
	default:
		return nil, ErrTokenHintInvalid
	}
}
