package web

import (
	"github.com/gorilla/context"
	"net/http"
	"oauth2-server/log"
	"oauth2-server/session"
)

type parseFormMiddleware struct {
}

// ServeHTTP ....
func (m *parseFormMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// 表单解析
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	next(w, r)
}

type guestMiddleware struct {
	service ServiceInterface
}

func newGuestMiddleware(service ServiceInterface) *guestMiddleware {
	return &guestMiddleware{service: service}
}

func (m *guestMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	m.service.setSessionService(r, w)
	sessionService := m.service.GetSessionService()

	if err := sessionService.StartSession(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	context.Set(r, sessionServiceKey, sessionService)
	next(w, r)
}

// 登录校验
type loggedInMiddleware struct {
	service ServiceInterface
}

func newLoggedInMiddleware(service ServiceInterface) *loggedInMiddleware {
	return &loggedInMiddleware{service: service}
}
func (m *loggedInMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	m.service.setSessionService(r, w)
	sessionService := m.service.GetSessionService()

	if err := sessionService.StartSession(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	context.Set(r, sessionServiceKey, sessionService)
	userSession, err := sessionService.GetUserSession()
	if err != nil {
		query := r.URL.Query()
		query.Set("login_redirect_uri", r.URL.Path)
		redirectWithQueryString("/web/login", query, w, r)
		return
	}
	if err := m.authenticate(userSession); err != nil {
		query := r.URL.Query()
		query.Set("login_redirect_uri", r.URL.Path)
		redirectWithQueryString("/web/login", query, w, r)
		return
	}
	_ = sessionService.SetUserSession(userSession)
	next(w, r)

}

func (m *loggedInMiddleware) authenticate(userSession *session.UserSession) error {
	log.INFO.Println("检查token是否有效，如果无效请重新生成。")
	_, err := m.service.GetOauthService().Authenticate(userSession.AccessToken)
	if err == nil {
		return nil
	}
	client, err := m.service.GetOauthService().FindClientByClientID(userSession.ClientID)
	if err != nil {
		return err
	}
	theRefreshToken, err := m.service.GetOauthService().GetValidRefreshToken(userSession.RefreshToken, client)
	if err != nil {
		return err
	}
	accessToken, refreshToken, err := m.service.GetOauthService().Login(
		theRefreshToken.Client,
		theRefreshToken.User,
		theRefreshToken.Scope,
	)
	if err != nil {
		return err
	}
	userSession.AccessToken = accessToken.Token
	userSession.RefreshToken = refreshToken.Token
	return nil
}

type clientMiddleware struct {
	service ServiceInterface
}

func newClientMiddleware(service ServiceInterface) *clientMiddleware {
	return &clientMiddleware{service: service}
}

func (m *clientMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	client, err := m.service.GetOauthService().FindClientByClientID(r.Form.Get("client_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	context.Set(r, clientKey, client)
	next(w, r)
}
