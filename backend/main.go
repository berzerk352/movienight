package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	database "kikisdeliveryserver.net/movienight_rest/database"
	moviehandler "kikisdeliveryserver.net/movienight_rest/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(cors.Handler((cors.Options{
		AllowedOrigins:   []string{"http://localhost*", "https://localhost*"},
		AllowedHeaders:   []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})))

	defer database.DBConn.Close()

	r.Mount("/movies", MovieRoutes())
	http.ListenAndServe(":3000", r)
}

func MovieRoutes() chi.Router {
	r := chi.NewRouter()
	movieHandler := moviehandler.MovieHandler{}

	r.Get("/", movieHandler.ListMovies)

	return r
}
