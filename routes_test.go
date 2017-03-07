package main

import (
	"github.com/cirm/decker/models"
	"time"
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/cirm/decker/handlers"
	"github.com/unrolled/render"
	"errors"
)

type mockDB struct{}

func (mdb *mockDB) AllPlayers() ([]*models.Player, error) {
	t := time.Date(2006, 1, 1, 12, 0, 0, 0, time.UTC)
	players := make([]*models.Player, 0)
	players = append(players, &models.Player{1, "kasutaja", t, t, t})
	players = append(players, &models.Player{2, "lahutaja", t, t, t})
	return players, nil
}

type mockLog struct{}

func (mlog *mockLog) Debug(*models.DbQuery) () {
}
func (mlog *mockLog) Info(*models.HttpRequest) () {
}
func (mlog *mockLog) Error(*models.HttpRequest) () {
}


func TestPlayersIndex(t *testing.T) {
	env := &handlers.Env{DB: &mockDB{}, Logger: &mockLog{}, Render: render.New()}
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/players/v1/players", nil)

	http.Handler(handlers.AppHandler{env, handlers.GetPlayers}).ServeHTTP(rec, req)

	expected := `[{"id":1,"username":"kasutaja","created":"2006-01-01T12:00:00Z","updated":"2006-01-01T12:00:00Z","visited":"2006-01-01T12:00:00Z"},{"id":2,"username":"lahutaja","created":"2006-01-01T12:00:00Z","updated":"2006-01-01T12:00:00Z","visited":"2006-01-01T12:00:00Z"}]`
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}


type failDB struct{}

func (mdb *failDB) AllPlayers() ([]*models.Player, error) {
	return nil, errors.New("Bad connection")
}

func TestPlayersIndexDBErr(t *testing.T) {
	env := &handlers.Env{DB: &failDB{}, Logger: &mockLog{}, Render: render.New()}
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/players/v1/players", nil)

	http.Handler(handlers.AppHandler{env, handlers.GetPlayers}).ServeHTTP(rec, req)

	statusCode := http.StatusInternalServerError
	if statusCode != rec.Code {
		t.Errorf("\n...expected = %v\n...obtained = %v", statusCode, rec.Code)
	}

	expected := "Bad connection"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}