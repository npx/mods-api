package modsapi

import (
	"encoding/json"
	"net/http"
)

// Module Interface every module has to implement
type Module interface {
	ID() string
	Endpoints() []Endpoint
}

// Endpoint is a mapping of a route within a module and its Handler
type Endpoint struct {
	Route   string
	Handler func(r *JSONRequest) JSONResponse
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

// JSONRequest wraps the requests
type JSONRequest struct {
	RawRequest *http.Request
}

// ParsedBody parses the request and returns error
func (r *JSONRequest) ParsedBody(i interface{}) error {

	decoder := json.NewDecoder(r.RawRequest.Body)
	err := decoder.Decode(&i)
	defer r.RawRequest.Body.Close()

	return err
}