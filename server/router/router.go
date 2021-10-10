package router

import (
	"bookshelf/server/app"
	"bookshelf/server/requestlog"

	"github.com/go-chi/chi"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()
	// Routes for healthz
	r.Get("/healthz/liveness", app.HandleLive)
	r.Method("GET", "/healthz/readiness", requestlog.NewHandler(a.HandleReady, l))
	r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))

	return r
}
