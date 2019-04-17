package web

import (
	"net/http"
	"oauth2-server/log"
)

func (s *Service) registerForm(w http.ResponseWriter, r *http.Request) {
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	errMsg, _ := sessionService.GetFlashMessage()
	_ = renderTemplate(w, "register.html", map[string]interface{}{
		"error":       errMsg,
		"queryString": getQueryString(r.URL.Query()),
	})
}
func (s *Service) register(w http.ResponseWriter, r *http.Request) {
	sessionService, err := getSessionService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//	TODO 校验是否注册过
	log.INFO.Println(sessionService)
	redirectWithQueryString("/web/login", r.URL.Query(), w, r)
}
