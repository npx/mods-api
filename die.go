package modsapi

import (
	"math/rand"
	"net/http"
)

// DiceModule is the implementation of a Dice
type DiceModule struct{}

// ID of the Dice Modul, which will be used in the URL
func (module DiceModule) ID() string {
	return "dice/roll"
}

// Endpoints lists all endpoints of the module
func (module DiceModule) Endpoints() []Endpoint {
	return []Endpoint{
		Endpoint{"", module.rollDice},
	}
}

// Response of the Dice Roll
type diceResponse struct {
	Result int `json:"result"`
}

// The RollDice function
func (module DiceModule) rollDice(r *http.Request) JSONResponse {
	value := rand.Intn(5) + 1

	return JSONResponse{
		http.StatusOK,
		diceResponse{value},
	}
}
