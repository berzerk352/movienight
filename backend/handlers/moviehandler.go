package handlers

import (
	"encoding/json"
	"net/http"

	models "kikisdeliveryserver.net/movienight_rest/models"
)

type MovieHandler struct {
}

func (m MovieHandler) ListMovies(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(models.ListMovies())
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
