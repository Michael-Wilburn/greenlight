package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Create a fixed-format JSON response from a string. Notice how we're using a raw
	// string literal(enclose with backticks) so that we can include double-quote
	// characters in the JSON withput needing to escape them? We also use the %q verb to
	// wrap the interpolated values in double-quotes.
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	// Set the Content-Type header to application/json so that the client knows to
	// expect JSON data. Go will deffault to sending a
	//"Content-type: text/plain; charset=utf8" header instead
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON as the HTTP response body.
	w.Write([]byte(js))
}
