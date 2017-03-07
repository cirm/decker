package models

import (
	"time"
	//"golang.org/x/crypto/bcrypt"
	"github.com/cirm/decker/errors"
)

// RawPlayer ...
type RawPlayer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

// Player ...
type Player struct {
	ID       int  `json:"id"`
	Username string `json:"username"`
	Created  time.Time        `json:"created"`
	Updated  time.Time        `json:"updated"`
	Visited  time.Time        `json:"visited"`
}

func (db *DB) AllPlayers() ([]*Player, error) {
	var (
		players []*Player
	)
	qs := "SELECT id, username, created, updated, visited FROM decker.players;"
	//c.Logger.Debug("dbQuery",
	//	zap.String("query", qs),
	//	zap.String(c.XRequestKey, r.Header.Get(c.XRequestKey)))
	//start := time.Now()
	rows, err := db.QueryD(qs)
	//latency := time.Since(start)
	if err != nil {
		//	c.Logger.Debug("dbQuery",
		//		zap.Duration("latency", latency),
		//		zap.String("status", err.Error()),
		//		zap.String(c.XRequestKey, r.Header.Get(c.XRequestKey)))
		return nil, errors.StatusError{500, err}
	}
	//c.Logger.Debug("dbQuery",
	//	zap.Duration("latency", latency),
	//	zap.String("status", "OK"),
	//	zap.String(c.XRequestKey, r.Header.Get(c.XRequestKey)))
	defer rows.Close()
	for rows.Next() {
		player := new(Player)
		err := rows.Scan(
			&player.ID,
			&player.Username,
			&player.Created,
			&player.Updated,
			&player.Visited)
		players = append(players, player)
		if err != nil {
			return nil, errors.StatusError{500, err}
		}
	}
	return players, nil
}

// SavePlayer ...
//func SavePlayer(r RawPlayer, c *env.Env) (Player, error) {
//	var player Player
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
//	if err != nil {
//		return player, err
//	}
//	row := c.Db.QueryRow(
//		`INSERT INTO decker.players (username, hpassword) VALUES ($1, $2) RETURNING id, username, created, updated, visited;`,
//		r.Username, string(hashedPassword))
//	if err := row.Scan(
//		&player.ID,
//		&player.Username,
//		&player.Created,
//		&player.Updated,
//		&player.Visited); err != nil {
//		return player, err
//	}
//	return player, nil
//}
