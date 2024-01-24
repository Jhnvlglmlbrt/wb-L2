// app.go
package app

import (
	"log"
	"net/http"
	"time"

	"github.com/Jhnvlglmlbrt/develop/dev11/internal/api"
	"github.com/Jhnvlglmlbrt/develop/dev11/internal/cache"
)

const (
	serverPort = ":8080"
)

type App struct {
	storage      *cache.Cache
	cacheHandler *api.CacheHandler
}

func NewApp() *App {
	cache := cache.NewCache()
	cacheHandler := api.NewHandlers(cache)

	return &App{
		storage:      cache,
		cacheHandler: cacheHandler,
	}
}

func (a *App) setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", a.mwLogger(a.cacheHandler.CreateEventHandler))
	mux.HandleFunc("/update_event", a.mwLogger(a.cacheHandler.UpdateEventHandler))
	mux.HandleFunc("/delete_event", a.mwLogger(a.cacheHandler.DeleteEventHandler))
	mux.HandleFunc("/events_for_day", a.mwLogger(a.cacheHandler.GetDayEventHandler))
	mux.HandleFunc("/events_for_week", a.mwLogger(a.cacheHandler.GetWeekEventHandler))
	mux.HandleFunc("/events_for_month", a.mwLogger(a.cacheHandler.GetMonthEventHandler))

	return mux
}

func (a *App) Run() {
	mux := a.setupRoutes()

	log.Printf("Server started on localhost%s...", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, mux))
}

func (a *App) mwLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("Executing %s %s %s", r.Method, r.RequestURI, time.Since(start))
	}
}
