package server

import (
	"net/http"

	"github.com/npx/mods-api/module"
)

func handle404(w http.ResponseWriter, r *http.Request) {
	response := module.JSONResponse{
		Status: http.StatusNotFound,
		Body: module.JSONError{
			Error:  "Not Found",
			Detail: "The Endpoint you requested does not exist",
		},
	}
	respond(response, w)
}

func makeHandler(endpoint module.Endpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// creating an object of JSONRequest
		rq := module.JSONRequest{RawRequest: r}

		// Run module's handler function
		response := endpoint.Handler(&rq)
		// Respond
		respond(response, w)
	}
}
