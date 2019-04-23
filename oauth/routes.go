package oauth

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
			Name:        "oauth_tokens",
			Method:      "POST",
			Pattern:     "/tokens",
			HandlerFunc: s.tokensHandler,
		},
		{
			Name:        "oauth_introspect",
			Method:      "POST",
			Pattern:     "/introspect",
			HandlerFunc: s.introspectHandler,
		},
	}
}
