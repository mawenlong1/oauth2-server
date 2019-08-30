package web

import (
	"net/http"
	"oauth2-server/user"
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
	if s.oauthService.UserExists(r.Form.Get("email")) {
		_ = sessionService.SetFlashMessage("邮件已经存在！")
		http.Redirect(w, r, r.RequestURI, http.StatusFound)
		return
	}
	_, err = s.oauthService.CreateUser(user.User, r.Form.Get("email"), r.Form.Get("password"))
	if err != nil {
		_ = sessionService.SetFlashMessage(err.Error())
		http.Redirect(w, r, r.RequestURI, http.StatusFound)
		return
	}
	redirectWithQueryString("/web/login", r.URL.Query(), w, r)
}
