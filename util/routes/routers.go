package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Middlewares []negroni.Handler
}

// AddRouters ...
func AddRouters(routers []Route, router *mux.Router) {
	var (
		handler http.Handler
		n       *negroni.Negroni
	)
	for _, route := range routers {
		// 添加中间件
		if len(route.Middlewares) > 0 {
			n = negroni.New()
			for _, middleware := range route.Middlewares {
				n.Use(middleware)
			}
			n.Use(negroni.Wrap(route.HandlerFunc))
			handler = n
		} else {
			handler = route.HandlerFunc
		}
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler)
	}
}
