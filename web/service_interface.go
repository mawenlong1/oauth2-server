package web

import (
	"github.com/gorilla/mux"
	"net/http"
	"oauth2-server/config"
	"oauth2-server/oauth"
	"oauth2-server/session"
	"oauth2-server/util/routes"
)

//ServiceInterface ...
type ServiceInterface interface {
	GetConfig() *config.Config
	GetOauthService() oauth.ServiceInterface
	GetSessionService() session.ServiceInterface
	GetRoutes() []routes.Route
	RegisterRoutes(router *mux.Router, profix string)
	Close()

	setSessionService(r *http.Request, w http.ResponseWriter)
}
