package main

import (
	"fmt"
	"github.com/cirm/decker/env"
	"net/http"
)

func initializeRoutes(c *env.AppContext) {
	c.Router.Handle("/api/players/v1/players", env.AppHandler{c, getPlayers})
	// c.Router.Handle("/api/players/v1/players", appHandler{c, createPlayer}).Methods("POST")
	c.Router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

}