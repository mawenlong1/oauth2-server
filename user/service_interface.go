package user

import (
	"github.com/gorilla/mux"
	"oauth2-server/util/routes"
)

//ServiceInterface ...
type ServiceInterface interface {
	GetRoutes() []routes.Route
	RegisterRoutes(router *mux.Router, prefix string)
	Close()
}
