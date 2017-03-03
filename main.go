package main

import (
	"github.com/cirm/decker/db"
	"github.com/cirm/decker/env"
	"github.com/urfave/negroni"
	"github.com/go-zoo/bone"
	"github.com/unrolled/render"
	"net/http"
	"go.uber.org/zap"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/cirm/decker/xrequestid"
)

func initRender(c *env.AppContext) {
	c.Render = render.New()
}

func main() {
	ctx := &env.AppContext{}
	mux := bone.New()
	ctx.Router = mux;
	db.InitPg(ctx)
	initRender(ctx)
	n := negroni.New(negroni.NewRecovery(), negroni.NewStatic(http.Dir("public")))
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.Use(xrequestid.New(4))
	logger, _ := zap.NewDevelopment()
	ctx.Logger = logger
	//n.Use(zap.NewProduction)
	initializeRoutes(ctx)
	n.UseHandler(ctx.Router)
	n.Run(":3030")
}
