package main

import (
	"github.com/urfave/negroni"
	"github.com/go-zoo/bone"
	"github.com/unrolled/render"
	"net/http"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/cirm/decker/xrequestid"
	"log"
	"github.com/cirm/decker/models"
	"github.com/cirm/decker/handlers"
)

func initRender(c *handlers.Env) {
	c.Render = render.New()
}

func main() {
	db, err := models.NewDB("dbname=arco user=spark password=salasala host=postgres1.cydec port=5432 sslmode=disable");
	if err != nil {
		log.Panic(err)
	}
	env := &handlers.Env{}
	env.DB = db
	env.XRequestKey = "X-Decker-Request-Id"
	mux := bone.New()
	env.Router = mux
	initRender(env)
	n := negroni.New(negroni.NewRecovery(), negroni.NewStatic(http.Dir("public")))
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.Use(xrequestid.New(4))
	logger, _ := models.NewLogger()
	if err != nil {
		log.Panic(err)
	}
	env.Logger = logger
	//n.Use(zap.NewProduction)
	initializeRoutes(env)
	n.UseHandler(env.Router)
	n.Run(":3030")
}
