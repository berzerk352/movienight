package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	database "kikisdeliveryserver.net/movienight_rest/database"
	moviehandler "kikisdeliveryserver.net/movienight_rest/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

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
