package modsapi

import "net/http"

// Module Interface every module has to implement
type Module interface {
	ID() string
	Endpoints() []Endpoint
}

// Endpoint is a mapping of a route within a module and its Handler
type Endpoint struct {
	Route   string
	Handler func(r *http.Request) JSONResponse
}

// JSONResponse holds the status code and the body of the response
type JSONResponse struct {
	status int
	body   interface{}
}

// JSONError defines what an error should look like across the API
type JSONError struct {
	Error  string `json:"error"`
	Detail string `json:"detail"`
}
