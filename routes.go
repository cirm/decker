package main

import (
	"fmt"
	"net/http"
	"github.com/cirm/decker/handlers"
)

func initializeRoutes(c *handlers.Env) {
	c.Router.Get("/api/players/v1/players", handlers.AppHandler{c, handlers.GetPlayers})
	//c.Router.Post("/api/players/v1/players", env.AppHandler{c, createPlayer})
	c.Router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

}