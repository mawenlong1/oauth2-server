package health

import (
	"oauth2-server/util/routes"

	"github.com/gorilla/mux"
)

// ServiceInterface 对外接口
type ServiceInterface interface {
	GetRoutes() []routes.Route
	RegisterRouters(router *mux.Router, prefix string)
	Close()
}
