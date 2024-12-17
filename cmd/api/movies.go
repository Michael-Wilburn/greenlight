package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Michael-Wilburn/greenlight/internal/data"
)

// Add a creaateMovieHandler for the "POST /v1/movies" endpoint. For now we simply
// return a plain-text placeholder response.
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// Add a shwoMovieHandler for the "GET /v1/movies/:id" endpoint. For now, we retrieve
// the interpolated "id" paramter from the current URL and include it in a placeholder
// response.
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encounterd a problem and could not process your request", http.StatusInternalServerError)
	}

	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Otherwise, interpolate the movie ID in a placeholder response.
	fmt.Fprintf(w, "Showing movie with ID: %d\n", id)
}
