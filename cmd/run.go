package cmd

import (
	"mwl/oauth2-server/log"
	"mwl/oauth2-server/services"
	"mwl/oauth2-server/util/response"
	"net/http"
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
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.INFO.Println("hello")
		w.Write([]byte("hello"))
	})
	app.UseHandler(router)
	graceful.Run(":3001", 10*time.Second, app)
	return nil
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.INFO.Println("test")
	response.WriteJSON(w, map[string]interface{}{
		"healthy": true,
	}, 200)
	w.Write([]byte("test"))
}
