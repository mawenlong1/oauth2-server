package oauth

import (
	"net/http"
	"oauth2-server/log"
)

func (s *Service) tokensHandler(w http.ResponseWriter, r *http.Request) {
	log.INFO.Println("暂未实现")
}

func (s *Service) introspectHandler(w http.ResponseWriter, r *http.Request) {
	log.INFO.Println("暂未实现")
}
