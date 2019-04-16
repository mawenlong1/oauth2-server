package web

import (
	"net/http"
	"oauth2-server/config"
	"oauth2-server/oauth"
	"oauth2-server/session"
)

//Service ...
type Service struct {
	cfg            *config.Config
	oauthService   oauth.ServiceInterface
	sessionService session.ServiceInterface
}

//NewService ..
func NewService(cfg *config.Config, oauth oauth.ServiceInterface, session session.ServiceInterface) *Service {
	return &Service{
		cfg:            cfg,
		oauthService:   oauth,
		sessionService: session,
	}
}

//GetConfig ...
func (s *Service) GetConfig() *config.Config {
	return s.cfg
}

//GetOauthService ..
func (s *Service) GetOauthService() oauth.ServiceInterface {
	return s.oauthService
}

//GetSessionService ...
func (s *Service) GetSessionService() session.ServiceInterface {
	return s.sessionService
}

//Close ...
func (s *Service) Close() {

}

func (s *Service) setSessionService(r *http.Request, w http.ResponseWriter) {
	s.sessionService.SetSessionService(r, w)
}
