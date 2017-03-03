package main

import (
	"time"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"strconv"
	"github.com/cirm/decker/env"
)

type RawPlayer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

type Player struct {
	Id       int  `json:"id"`
	Username string `json:"username"`
	Created  time.Time        `json:"created"`
	Updated  time.Time        `json:"updated"`
	Visited  time.Time        `json:"visited"`
}

func SavePlayer(r RawPlayer, c *env.AppContext) (Player, error) {
	var player Player
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return player, err
	}
	fmt.Println("Password ", string(hashedPassword))
	fmt.Println("username ", r.Username)
	stmt := fmt.Sprintf(`INSERT INTO decker.players (username, hpassword) VALUES (%s, %s);`, strconv.Quote(r.Username), string(hashedPassword))
	fmt.Println(stmt)
	rows := c.Db.QueryRow(`INSERT INTO decker.players (username, hpassword) VALUES ($1, $2);`, r.Username, string(hashedPassword))
	// /rows, err := c.db.Query(`INSERT INTO decker.players (username, hpassword) VALUES (username, hpassword);`, string(r.Username), string(hashedPassword))
	if err != nil {
		fmt.Println(err)
		return player, err
	}
	fmt.Println(rows.Scan(&player))
	fmt.Println("rows")
	return player, err;
}