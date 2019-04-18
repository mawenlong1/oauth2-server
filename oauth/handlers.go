package oauth

import (
	"errors"
	"net/http"
	"oauth2-server/log"
	"oauth2-server/models"
	"oauth2-server/util/response"
)

var (
	// ErrInvalidGrantType ..
	ErrInvalidGrantType = errors.New("Invalid grant type")
	// ErrInvalidClientIDOrSecret ...
	ErrInvalidClientIDOrSecret = errors.New("Ivalid client ID or secret")
	tokenTypes                 = "Bearer"
)

func (s *Service) tokensHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	grantTypes := map[string]func(r *http.Request, client *models.OauthClient) (*AccessTokenResponse, error){
		"authorization_code": s.authorizationCodeGrant,
		"password":           s.passwordGrant,
		"client_credentials": s.clientCredentialsGrant,
		"refresh_token":      s.refreshTokenGrant,
	}
	grantHandler, ok := grantTypes[r.Form.Get("grant_type")]
	if !ok {
		response.Error(w, ErrInvalidGrantType.Error(), http.StatusBadRequest)
		return
	}
	client, err := s.basicAuthClient(r)
	if err != nil {
		response.UnauthorizedError(w, err.Error())
		return
	}
	resp, err := grantHandler(r, client)
	if err != nil {
		response.Error(w, err.Error(), getErrStatusCode(err))
	}
	response.WriteJSON(w, resp, 200)
}

func (s *Service) authorizationCodeGrant(r *http.Request, client *models.OauthClient) (*AccessTokenResponse, error) {
	authorizationCode, err := s.getValidAuthorizationCode(
		r.Form.Get("code"),
		r.Form.Get("redirect_uri"),
		client,
	)
	if err != nil {
		return nil, err
	}
	accessToken, refreshToken, err := s.Login(
		authorizationCode.Client,
		authorizationCode.User,
		authorizationCode.Scope,
	)
	if err != nil {
		return nil, err
	}
	s.db.Unscoped().Delete(&authorizationCode)

	accessTokenResponse, err := NewAccessTokenResponse(
		accessToken,
		refreshToken,
		s.cfg.Oauth.AccessTokenLifeTime,
		tokenTypes,
	)
	if err != nil {
		return nil, err
	}
	return accessTokenResponse, nil
}
func (s *Service) passwordGrant(r *http.Request, client *models.OauthClient) (*AccessTokenResponse, error) {
	scope, err := s.GetScope(r.Form.Get("scope"))
	if err != nil {
		return nil, err
	}
	user, err := s.AuthUser(r.Form.Get("username"), r.Form.Get("password"))
	if err != nil {
		return nil, ErrInvalidUsernameOrPassword
	}
	accessToken, refreshToken, err := s.Login(client, user, scope)
	if err != nil {
		return nil, err
	}
	accessTokenResponse, err := NewAccessTokenResponse(
		accessToken,
		refreshToken,
		s.cfg.Oauth.AccessTokenLifeTime,
		tokenTypes,
	)
	if err != nil {
		return nil, err
	}
	return accessTokenResponse, nil
}

func (s *Service) clientCredentialsGrant(r *http.Request, client *models.OauthClient) (*AccessTokenResponse, error) {
	scope, err := s.GetScope(r.Form.Get("scope"))
	if err != nil {
		return nil, err
	}
	accessToken, err := s.GrandAccessToken(client, nil, s.cfg.Oauth.AccessTokenLifeTime, scope)
	if err != nil {
		return nil, err
	}
	accessTokenResponse, err := NewAccessTokenResponse(
		accessToken,
		nil,
		s.cfg.Oauth.AccessTokenLifeTime,
		tokenTypes,
	)
	if err != nil {
		return nil, err
	}
	return accessTokenResponse, nil
}
func (s *Service) refreshTokenGrant(r *http.Request, client *models.OauthClient) (*AccessTokenResponse, error) {
	theRefreshToken, err := s.GetValidRefreshToken(r.Form.Get("refresh_token"), client)
	if err != nil {
		return nil, err
	}
	scope, err := s.getRefreshTokenScope(theRefreshToken, r.Form.Get("scope"))
	if err != nil {
		return nil, err
	}
	accessToken, refreshToken, err := s.Login(theRefreshToken.Client, theRefreshToken.User, scope)
	if err != nil {
		return nil, err
	}
	accessTokenResponse, err := NewAccessTokenResponse(
		accessToken,
		refreshToken,
		s.cfg.Oauth.AccessTokenLifeTime,
		tokenTypes,
	)
	if err != nil {
		return nil, err
	}
	return accessTokenResponse, nil
}
func (s *Service) introspectHandler(w http.ResponseWriter, r *http.Request) {
	log.INFO.Println("暂未实现")
}

func (s *Service) basicAuthClient(r *http.Request) (*models.OauthClient, error) {
	clientID, secret, ok := r.BasicAuth()
	if !ok {
		return nil, ErrInvalidClientIDOrSecret
	}
	client, err := s.AuthClient(clientID, secret)
	if err != nil {
		return nil, ErrInvalidClientIDOrSecret
	}
	return client, err
}
