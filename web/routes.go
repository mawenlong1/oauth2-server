package web

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
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
			Name:        "register_form",
			Method:      "GET",
			Pattern:     "/register",
			HandlerFunc: s.registerForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "register",
			Method:      "POST",
			Pattern:     "/register",
			HandlerFunc: s.register,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "login_form",
			Method:      "GET",
			Pattern:     "/login",
			HandlerFunc: s.loginForm,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
		{
			Name:        "login",
			Method:      "POST",
			Pattern:     "/login",
			HandlerFunc: s.login,
			Middlewares: []negroni.Handler{
				new(parseFormMiddleware),
				newGuestMiddleware(s),
				newClientMiddleware(s),
			},
		},
	}
}
