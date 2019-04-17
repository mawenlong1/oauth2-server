package web

import (
	"github.com/gorilla/context"
	"net/http"
	"oauth2-server/log"
	"oauth2-server/session"
)

type parseFormMiddleware struct {
}

//ServeHTTP ....
func (m *parseFormMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//表单解析
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

//登录校验
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

func (m *loggedInMiddleware) authenticate(session *session.UserSession) error {
	log.INFO.Println("clientMiddleware，需要调用oauth，暂时未实现")
	return nil
}

type clientMiddleware struct {
	service ServiceInterface
}

func newClientMiddleware(service ServiceInterface) *clientMiddleware {
	return &clientMiddleware{service: service}
}

func (m *clientMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.INFO.Println("clientMiddleware，需要调用oauth，暂时未实现")
	next(w, r)
}
