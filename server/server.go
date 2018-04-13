package server

import (
	"encoding/json"
	"net/http"

	"github.com/npx/mods-api/module"
)

// Setup the server
func Setup(modules []module.Module) {
	// Register all modules
	for _, module := range modules {
		registerModule(module)
	}

	// Register 404 handler
	http.HandleFunc("/", handle404)
}

// Start the server
func Start() {
	http.ListenAndServe("0.0.0.0:8000", nil)
}

func respond(response module.JSONResponse, w http.ResponseWriter) {
	// Set Response Headers
	w.Header().Set("Content-Type", "application/json")

	// Set Status Code
	w.WriteHeader(response.Status)

	// Encode json and respond
	json.NewEncoder(w).Encode(response.Body)
}

func registerModule(module module.Module) {
	basePath := "/" + module.ID()
	for _, endpoint := range module.Endpoints() {
		route := basePath + endpoint.Route
		http.HandleFunc(route, makeHandler(endpoint))
	}
}
