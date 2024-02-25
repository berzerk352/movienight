package models

import (
	"context"
	"log"

	"kikisdeliveryserver.net/movienight_rest/database"
)

type Movie struct {
	Id        int64  `json:"id"`
	MovieName string `json:"movie_name"`
	Submitter string `json:"submitter"`
}

var movies = RetrieveMovieRecords()

func RetrieveMovieRecords() []*Movie {
	var results []*Movie

	rows, err := database.DBConn.Query(context.Background(), "select id, movie_name, submitter from public.movies")
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}

	for rows.Next() {
		values, err := rows.Values()

		if err != nil {
			log.Fatalf("Issue with row: %v\n", err)
		}
		r := &Movie{
			Id:        values[0].(int64),
			MovieName: values[1].(string),
			Submitter: values[2].(string),
		}
		results = append(results, r)
	}

	return results
}

func ListMovies() []*Movie {
	return movies
}
