package user

import "oauth2-server/oauth"

//Service ...
type Service struct {
	oauthService oauth.ServiceInterface
}

//NewService ..
func NewService(oauthService oauth.ServiceInterface) *Service {
	return &Service{
		oauthService: oauthService,
	}
}

//GetOauthService ...
func (s *Service) GetOauthService() oauth.ServiceInterface {
	return s.oauthService
}

//Close ...
func (s *Service) Close() {

}
