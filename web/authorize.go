package web

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"oauth2-server/models"
	"oauth2-server/session"
	"strconv"
)

var (
	// ErrIncorrectResponseType ..
	ErrIncorrectResponseType = errors.New("web:Response type not one of token or code")
)

func (s *Service) authorizeForm(w http.ResponseWriter, r *http.Request) {
	sessionService, client, _, responseType, _, err := s.authorizeCommon(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	errMsg, _ := sessionService.GetFlashMessage()
	query := r.URL.Query()
	query.Set("login_redirect_uri", r.URL.Path)
	_ = renderTemplate(w, "authorize.html", map[string]interface{}{
		"error":       errMsg,
		"ClientID":    client.ClientKey,
		"queryString": getQueryString(query),
		"token":       responseType == "token",
	})
}
func (s *Service) authorizeCommon(r *http.Request) (session.ServiceInterface, *models.OauthClient, *models.OauthUser, string, *url.URL, error) {
	sessionService, err := getSessionService(r)
	if err != nil {
		return nil, nil, nil, "", nil, err
	}
	client, err := getClient(r)
	if err != nil {
		return nil, nil, nil, "", nil, err
	}
	userSession, err := sessionService.GetUserSession()
	if err != nil {
		return nil, nil, nil, "", nil, err
	}
	user, err := s.oauthService.FindUserByUserName(userSession.Username)
	if err != nil {
		return nil, nil, nil, "", nil, err
	}
	responseType := r.Form.Get("response_type")
	if responseType != "code" && responseType != "token" {
		return nil, nil, nil, "", nil, err
	}
	redirectURI := r.Form.Get("redirect_uri")
	if redirectURI == "" {
		redirectURI = client.RedirectURI.String
	}
	parsedRedirectURI, err := url.ParseRequestURI(redirectURI)
	if err != nil {
		return nil, nil, nil, "", nil, err
	}

	return sessionService, client, user, responseType, parsedRedirectURI, nil

}
func (s *Service) authorize(w http.ResponseWriter, r *http.Request) {
	_, client, user, responseType, redirectURI, err := s.authorizeCommon(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	state := r.Form.Get("state")
	authorized := len(r.Form.Get("allow")) > 0
	if !authorized {
		errorRedirect(w, r, redirectURI, "access_denied", state, responseType)
		return
	}
	scope, err := s.oauthService.GetScope(r.Form.Get("scope"))
	if err != nil {
		errorRedirect(w, r, redirectURI, "invalid_scope", state, responseType)
		return
	}
	query := redirectURI.Query()
	if responseType == "code" {
		authorizationCode, err := s.oauthService.GrantAuthorizationCode(
			client,
			user,
			s.cfg.Oauth.AuthCodeLifeTime,
			redirectURI.String(),
			scope,
		)
		if err != nil {
			errorRedirect(w, r, redirectURI, "server_error", state, responseType)
			return
		}
		query.Set("code", authorizationCode.Code)
		if state != "" {
			query.Set("state", state)
		}
		redirectWithQueryString(redirectURI.String(), query, w, r)
		return
	}
	if responseType == "token" {
		lifetime, err := strconv.Atoi(r.Form.Get("lifetime"))
		if err != nil {
			errorRedirect(w, r, redirectURI, "server_error", state, responseType)
			return
		}
		accessToken, err := s.oauthService.GrandAccessToken(
			client,
			user,
			lifetime,
			scope,
		)
		if err != nil {
			errorRedirect(w, r, redirectURI, "server_error", state, responseType)
			return
		}
		query.Set("access_token", accessToken.Token)
		query.Set("expires_in", fmt.Sprintf("%d", lifetime))
		query.Set("token_type", "Bearer")
		query.Set("scope", scope)
		if state != "" {
			query.Set("state", state)
		}
		redirectWithFragment(redirectURI.String(), query, w, r)
	}
}
