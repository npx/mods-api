package modsapi

import (
	"net/http"
)

func handle404(w http.ResponseWriter, r *http.Request) {
	response := JSONResponse{
		http.StatusNotFound,
		JSONError{
			"Not Found",
			"The Endpoint you requested does not exist",
		},
	}
	respond(response, w)
}

func makeHandler(endpoint Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Run module's handler function
		response := endpoint.Handler(r)
		// Respond
		respond(response, w)
	}
}
