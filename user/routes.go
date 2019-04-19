package user

import (
	"github.com/gorilla/mux"
	"oauth2-server/util/routes"
)

// RegisterRoutes ..
func (s *Service) RegisterRoutes(router *mux.Router, prefix string) {
	subRouter := router.PathPrefix(prefix).Subrouter()
	routes.AddRouters(s.GetRoutes(), subRouter)
}

// GetRoutes ..
func (s *Service) GetRoutes() []routes.Route {
	return []routes.Route{
		{
			Name:        "create_user",
			Method:      "POST",
			Pattern:     "/create",
			HandlerFunc: s.createUser,
		},
	}
}
