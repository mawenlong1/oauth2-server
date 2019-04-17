package cmd

import (
	"oauth2-server/services"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	graceful "gopkg.in/tylerb/graceful.v1"
)

// Run 运行服务器
func Run(configFile string) error {
	cfg, db, err := initConfigDB(configFile)
	if err != nil {
		return err
	}
	defer db.Close()
	if err = services.Init(cfg, db); err != nil {
		return err
	}
	defer services.Close()
	// app := negroni.New()
	// app.Use(negroni.NewRecovery())
	// app.Use(negroni.NewLogger())
	// app.Use(gzip.Gzip(gzip.DefaultCompression))
	// app.Use(negroni.NewStatic(http.Dir("public")))
	app := negroni.Classic()
	// app.Use(gzip.Gzip(gzip.DefaultCompression))
	router := mux.NewRouter()
	services.HealthService.RegisterRouters(router, "/v1")
	services.WebService.RegisterRoutes(router, "/web")
	app.UseHandler(router)
	graceful.Run(":"+strconv.Itoa(cfg.ServerPort), 10*time.Second, app)
	return nil
}
