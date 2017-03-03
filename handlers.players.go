package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"encoding/json"
	"github.com/cirm/decker/env"
	"time"
	"go.uber.org/zap"
)

func getPlayers(c *env.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	var (
		player Player
		players []Player
	)
	qs := "SELECT id, username, created, updated, visited FROM decker.players;"
	c.Logger.Debug("dbQuery",
		zap.String("query", qs),
		zap.String("X-Decker-Request-Id", r.Header.Get("X-Decker-Request-Id")))
	start := time.Now()
	rows, err := c.Db.Query(qs);
	latency := time.Since(start)
	if err != nil {
		c.Logger.Debug("dbQuery",
			zap.Duration("latency", latency),
			zap.String("status", err.Error()),
			zap.String("X-Decker-Request-Id", r.Header.Get("X-Decker-Request-Id")))
		return 0, env.StatusError{500, err}
	} else {
		c.Logger.Debug("dbQuery",
			zap.Duration("latency", latency),
			zap.String("status", "OK"),
			zap.String("X-Decker-Request-Id", r.Header.Get("X-Decker-Request-Id")))
		defer rows.Close();
		for rows.Next() {
			err := rows.Scan(
				&player.Id,
				&player.Username,
				&player.Created,
				&player.Updated,
				&player.Visited)
			players = append(players, player)
			if err != nil {
				return 0, env.StatusError{500, err}
			}
		}
		c.Render.JSON(w, http.StatusOK, players)

	}
	return 200, nil
}

func createPlayer(c *env.AppContext, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		c.Render.JSON(w, http.StatusInternalServerError, err)
	}
	var raw RawPlayer
	if err := json.Unmarshal(body, &raw); err != nil {
		c.Render.JSON(w, http.StatusInternalServerError, err)
	}
	player, err := SavePlayer(raw, c)
	c.Render.JSON(w, http.StatusOK, player)
	fmt.Println(player)

}