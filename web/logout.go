package web

import "net/http"

func (s *Service) logout(w http.ResponseWriter, r *http.Request) {
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userSession, err := sessionService.GetUserSession()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.oauthService.ClearUserTokens(userSession)
	_ = sessionService.ClearUserSession()
	redirectWithQueryString("/web/login", r.URL.Query(), w, r)
}
