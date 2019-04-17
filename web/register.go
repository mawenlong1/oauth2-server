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
	log.INFO.Println(errMsg)
}
