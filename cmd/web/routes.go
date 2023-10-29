package main

import (
	"github.com/Shubham_bhatt/bookings/pkg/config"
	"github.com/Shubham_bhatt/bookings/pkg/handlers"
	
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
