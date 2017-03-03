package env

import (
	"database/sql"
	"net/http"
	"github.com/go-zoo/bone"
	"github.com/unrolled/render"
	"go.uber.org/zap"
	"time"
)

type AppHandler struct {
	*AppContext
	H func(c *AppContext, w http.ResponseWriter, r *http.Request) (int, error)
}

type AppContext struct {
	Db     *sql.DB
	Router *bone.Mux
	Render *render.Render
	Logger *zap.Logger
}

type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	status, err := h.H(h.AppContext, w, r)
	latency := time.Since(start)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			h.Logger.Error(e.Error(),
				zap.Int("status", e.Status()),
				zap.String("X-Decker-Request-Id", r.Header.Get("X-Decker-Request-Id")),
				zap.String("method", r.Method),
				zap.String("url", r.URL.Path),
				zap.Duration("latency", latency))
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			h.Logger.Error(http.StatusText(http.StatusInternalServerError),
				zap.Int("status", 500),
				zap.String("method", r.Method),
				zap.String("X-Decker-Request-Id", r.Header.Get("X-Decker-Request-Id")),
				zap.String("url", r.URL.Path),
				zap.Duration("latency", latency))
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	} else {
		h.Logger.Info("request",
			zap.Int("status", status),
			zap.String("X-Decker-Request-Id", r.Header.Get("X-Decker-Request-Id")),
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
			zap.Duration("latency", latency))
	}
}
