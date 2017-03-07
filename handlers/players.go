package handlers

import (
	"net/http"
	"github.com/cirm/decker/models"
	"fmt"
	//"time"
	"github.com/cirm/decker/errors"
)

func GetPlayers(c *Env, w http.ResponseWriter, r *http.Request) (int, error) {
	qs := "SELECT id, username, created, updated, visited FROM decker.players;"
	fmt.Printf(r.Header.Get(c.XRequestKey))
	c.Logger.Debug(&models.DbQuery{Message: "dbquery",
		QueryString:                    qs,
		XRQK:                           c.XRequestKey,
		XRQV:                           r.Header.Get(c.XRequestKey),
	})
	//c.Logger.Debug("dbQuery",
	//	zap.String("query", qs),
	//	zap.String(c.XRequestKey, r.Header.Get(c.XRequestKey)))
	//start := time.Now()
	players, err := c.DB.AllPlayers()
	//latency := time.Since(start)
	if err != nil {
		c.Logger.Debug(&models.DbQuery{Message: err.Error(),
			QueryString:                    qs,
			XRQK:                           c.XRequestKey,
			XRQV:                           r.Header.Get(c.XRequestKey),
		})
		return 0, errors.StatusError{500, err}
	}
	//c.Logger.Debug("dbQuery",
	//	zap.Duration("latency", latency),
	//	zap.String("status", "OK"),
	//	zap.String(c.XRequestKey, r.Header.Get(c.XRequestKey)))
	//defer rows.Close()
	//for rows.Next() {
	//	err := rows.Scan(
	//		&player.ID,
	//		&player.Username,
	//		&player.Created,
	//		&player.Updated,
	//		&player.Visited)
	//	players = append(players, player)
	//	if err != nil {
	//		return 0, env.StatusError{500, err}
	//	}
	//}
	c.Render.JSON(w, http.StatusOK, players)
	return 200, nil
}

//
//func createPlayer(c *env.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
//	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
//	if err != nil {
//		return 0, env.StatusError{500, err}
//	}
//	var raw RawPlayer
//	if err := json.Unmarshal(body, &raw); err != nil {
//		return 0, env.StatusError{500, err}
//
//	}
//	player, err := SavePlayer(raw, c)
//	if err != nil {
//		return 0, env.StatusError{500, err}
//	}
//	c.Render.JSON(w, http.StatusOK, player)
//	fmt.Println(player)
//	return 200, nil
//
//}
