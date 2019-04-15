package health

import (
	"oauth2-server/util/routes"

	"github.com/gorilla/mux"
)

// RegisterRouters ..
func (s *Service) RegisterRouters(router *mux.Router, prefix string) {
	subRouter := router.PathPrefix(prefix).Subrouter()
	routes.AddRouters(s.GetRoutes(), subRouter)
}

// GetRoutes ...
func (s *Service) GetRoutes() []routes.Route {
	return []routes.Route{
		{
			Name:        "health_check",
			Method:      "GET",
			Pattern:     "/health",
			HandlerFunc: s.healthCheck,
		},
	}
}
