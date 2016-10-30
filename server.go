package modsapi

import (
	"encoding/json"
	"net/http"
)

// TODO make a server "class"
func setupServer(modules []Module) {
	// Register all modules
	for _, module := range modules {
		registerModule(module)
	}

	// Register 404 handler
	http.HandleFunc("/", handle404)
}

func respond(response JSONResponse, w http.ResponseWriter) {
	// Set Response Headers
	w.Header().Set("Content-Type", "application/json")

	// Set Status Code
	w.WriteHeader(response.status)

	// Encode json and respond
	json.NewEncoder(w).Encode(response.body)
}

func registerModule(module Module) {
	basePath := "/" + module.ID()
	for _, endpoint := range module.Endpoints() {
		route := basePath + endpoint.Route
		http.HandleFunc(route, makeHandler(endpoint))
	}
}
