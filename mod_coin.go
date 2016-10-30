package modsapi

import (
	"math/rand"
	"net/http"
)

// CoinModule is the implementation of a Coin
type CoinModule struct{}

// ID of the CoinModule, which will be used in the URL
func (module CoinModule) ID() string {
	return "coin"
}

// Endpoints lists all endpoints of the module
func (module CoinModule) Endpoints() []Endpoint {
	return []Endpoint{
		Endpoint{"/flip", module.flipCoin},
	}
}

// Response of the Coin flip
type coinResponse struct {
	Result string `json:"result"`
}

// The coinFlip function
func (module CoinModule) flipCoin(r *JSONRequest) JSONResponse {
	value := rand.Intn(2) + 1

	var side string

	if value == 1 {
		side = "heads"
	} else {
		side = "tails"
	}

	return JSONResponse{
		http.StatusOK,
		coinResponse{side},
	}
}
