package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Michael-Wilburn/greenlight/internal/data"
)

// Add a createMovieHandler for the "POST" /v1/movies endpoint. For now we simply
// return a plain-text placeholder response.

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

/* Add a showMovieHandler for the "GET" /v1/movies/:id endpoint. For now, we retrive
the interpolated "id" parameter from the current URL and include it in a placeholder
response. */

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	/*If the parameter couldn't be converted,
	or is less than 1, we know the ID is invalid so we use the http.NotFound()
	function to return a 404 Not Found Response.
	*/
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	/*
		Create a new instance of the Movie struct, containing the ID we extracted from
		the URL and some dummy data. Also notice that we deliberately haven't set a
		value for the Year field.
	*/
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Year:      1942,
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	// Encode the struct to JSON and sent it as the HTTP response.
	// Create an envelope{"movie": movie} instance and pass it to writeJSON(), instead
	// of passing the plain movie struct.
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
