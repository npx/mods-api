package modsapi

import "net/http"

// DeciderModule is the implementation of a Module that makes decisions
type DeciderModule struct{}

// ID of the decider module, which will be used in the URL
func (module DeciderModule) ID() string {
	return "decider"
}

// Endpoints lists all endpoints of the module
func (module DeciderModule) Endpoints() []Endpoint {
	return []Endpoint{
		Endpoint{"", module.postDecide},
	}
}

// define what a request to the decider should look like
type deciderRequest []string

// response of the decider
type deciderResponse struct {
	Choice string `json:"choice"`
}

// post handler
func (module DeciderModule) postDecide(r *JSONRequest) JSONResponse {
	// Only allow POST requests
	// TODO define allowed methods in endpoints
	if r.RawRequest.Method != "POST" {
		return JSONResponse{
			http.StatusMethodNotAllowed,
			JSONError{"Method Not Allowed", "Only POST allowed"},
		}
	}

	// Decode JSON body
	// TODO wrap http.Request and get parsed json body from there
	var request deciderRequest
	err := r.ParsedBody(request)
	if err != nil {
		return JSONResponse{
			http.StatusBadRequest,
			JSONError{
				"Invalid Parameters",
				"Cannot parse body",
			},
		}
	}

	// Make sure there are at least 2 choices
	choices := request
	if len(choices) < 2 {
		return JSONResponse{
			http.StatusBadRequest,
			JSONError{
				"Invalid Parameters",
				"Give at least two choices",
			},
		}
	}

	// Respond with first item
	return JSONResponse{
		http.StatusOK,
		deciderResponse{module.decide(choices)},
	}
}

// Implement decide algorithm
func (module *DeciderModule) decide(choices []string) string {
	if len(choices) < 1 {
		return ""
	}
	return choices[0]
}
