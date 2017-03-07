package handlers

import (
	"net/http"
	"time"
	"github.com/cirm/decker/errors"
	"github.com/cirm/decker/models"
	"github.com/go-zoo/bone"
	"github.com/unrolled/render"
)

type Env struct {
	DB          models.Datastore
	Router      *bone.Mux
	Render      *render.Render
	Logger      models.Logger
	XRequestKey string
}

type AppHandler struct {
	*Env
	H func(env *Env, w http.ResponseWriter, r *http.Request) (int, error)
}

func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	status, err := h.H(h.Env, w, r)
	latency := time.Since(start)
	if err != nil {
		switch e := err.(type) {
		case errors.Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			h.Logger.Error(&models.HttpRequest{
				Message: http.StatusText(http.StatusInternalServerError),
				Status:  500,
				XRQV:    r.Header.Get(h.Env.XRequestKey),
				XRQK:    h.Env.XRequestKey,
				Method:  r.Method,
				Path:    r.URL.Path,
				Latency: latency,
			})
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			h.Logger.Error(&models.HttpRequest{
				Message: http.StatusText(http.StatusInternalServerError),
				Status:  500,
				XRQV:    r.Header.Get(h.Env.XRequestKey),
				XRQK:    h.Env.XRequestKey,
				Method:  r.Method,
				Path:    r.URL.Path,
				Latency: latency,
			})
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	} else {
		h.Logger.Info(&models.HttpRequest{
			Message: "HTTPrequest",
			Status:  status,
			XRQK:    h.Env.XRequestKey,
			XRQV:    r.Header.Get(h.Env.XRequestKey),
			Method:  r.Method,
			Path:    r.URL.Path,
			Latency: latency,
		})
	}
}
