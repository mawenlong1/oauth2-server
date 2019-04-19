package session

import (
	"encoding/gob"
	"errors"
	"github.com/gorilla/sessions"
	"net/http"
	"oauth2-server/config"
)

// Service ...
type Service struct {
	sessionStore   sessions.Store
	sessionOptions *sessions.Options
	session        *sessions.Session
	r              *http.Request
	w              http.ResponseWriter
}

// UserSession ..
type UserSession struct {
	ClientID     string
	Username     string
	AccessToken  string
	RefreshToken string
}

var (
	// StorageSessionName ...
	StorageSessionName = "oauth2_server_session"
	// UserSessionKey ...
	UserSessionKey = "oauth2_server_user"
	// ErrSessionNotStarted ...
	ErrSessionNotStarted = errors.New("session:Session not started")
)

func init() {
	gob.Register(new(UserSession))
}

// NewService ...
func NewService(config *config.Config, store sessions.Store) *Service {
	return &Service{
		sessionStore: store,
		sessionOptions: &sessions.Options{
			Path:     config.Session.Path,
			MaxAge:   config.Session.MaxAge,
			HttpOnly: config.Session.HTTPOnly,
		},
	}
}

// SetSessionService ...
func (s *Service) SetSessionService(r *http.Request, w http.ResponseWriter) {
	s.r = r
	s.w = w
}

// StartSession ...
func (s *Service) StartSession() error {
	session, err := s.sessionStore.Get(s.r, StorageSessionName)
	if err != nil {
		return err
	}
	s.session = session
	return nil
}

// GetUserSession ...
func (s *Service) GetUserSession() (*UserSession, error) {
	if s.session == nil {
		return nil, ErrSessionNotStarted
	}
	userSession, ok := s.session.Values[UserSessionKey].(*UserSession)
	if !ok {
		return nil, errors.New("session:User session type assertion error")
	}
	return userSession, nil
}

// SetUserSession ...
func (s *Service) SetUserSession(userSession *UserSession) error {
	if s.session == nil {
		return ErrSessionNotStarted
	}
	s.session.Values[UserSessionKey] = userSession
	return s.session.Save(s.r, s.w)
}

// ClearUserSession ...
func (s *Service) ClearUserSession() error {
	if s.session == nil {
		return ErrSessionNotStarted
	}
	delete(s.session.Values, UserSessionKey)
	return s.session.Save(s.r, s.w)
}

// SetFlashMessage ..
func (s *Service) SetFlashMessage(msg string) error {
	if s.session == nil {
		return ErrSessionNotStarted
	}
	s.session.AddFlash(msg)
	return s.session.Save(s.r, s.w)
}

// GetFlashMessage ...
func (s *Service) GetFlashMessage() (interface{}, error) {
	if s.session == nil {
		return nil, ErrSessionNotStarted
	}
	if flashes := s.session.Flashes(); len(flashes) > 0 {
		_ = s.session.Save(s.r, s.w)
		return flashes[0], nil
	}
	return nil, nil
}

// Close ...
func (s *Service) Close() {

}
